// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package dazBlog

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/Daz-3ux/dBlog/internal/dazBlog/controller/v1/user"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
	"github.com/Daz-3ux/dBlog/internal/pkg/known"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	mw "github.com/Daz-3ux/dBlog/internal/pkg/middleware"
	pb "github.com/Daz-3ux/dBlog/pkg/proto/dazBlog/v1"
	"github.com/Daz-3ux/dBlog/pkg/token"
	"github.com/Daz-3ux/dBlog/pkg/version/verflag"
)

var cfgFile string

// NewDazBlogCommand create the *cobra.Command object
// so can use the Execute method to start
func NewDazBlogCommand() *cobra.Command {
	// cmd is the *cobra.Command object and the top-level command
	cmd := &cobra.Command{
		// specify the name of the command
		Use:   "dBlog",                         // specify the short description of the command
		Short: "dBlog is a simple blog system", // specify the long description of the command
		Long: `dBlog is a simple but not easy blog system,
Find more dBlog information at:
	https://github.com/daz-3ux/dBlog#readme`,

		// when an error occurs, the command will not print usage information
		SilenceUsage: true, // specify the run function to execute when cmd.Execute() is called
		// if the function fails, an error message will be returned
		RunE: func(cmd *cobra.Command, args []string) error {
			verflag.PrintAndExitIfRequested()

			log.Init(logOptions())
			defer log.Sync()

			return run()
		}, // no need to specify command line parameters
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		},
	}

	// The following settings make the initConfig function be called
	// every time when the command is executed to read the configuration.
	cobra.OnInitialize(initConfig)

	// define custom flag and config

	// cobra supports a persistent flag
	// for providing options that work across all subs-commands of a command
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the dBlog configuration file. Empty string for no configuration file.")

	// cobra supports local flag
	// which can only be used within the command to which they are bound
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add --version flag
	verflag.AddFlags(cmd.PersistentFlags())

	return cmd
}

func run() error {
	// init the store layer
	if err := initStore(); err != nil {
		return err
	}

	// set the secret key used to sign and parse the token
	token.Init(viper.GetString("jwt-secret"), known.XUsernameKey)

	// set Gin mode
	gin.SetMode(viper.GetString("runmode"))

	// create Gin engine
	g := gin.New()

	// gin.Recovery() middleware is used to capture any panics and recover from them
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.RequestID()}

	g.Use(mws...)

	if err := installRouters(g); err != nil {
		return err
	}

	// create HTTP server
	httpsrv := startInsecureServer(g)
	// create HTTPS server
	httpssrv := startSecureServer(g)
	// create gRPC server
	grpcsrv := startGRPCServer()

	// wait for an interrupt signal to gracefully shut down the server with a 10s timeout
	quit := make(chan os.Signal, 1)
	// kill (no param) default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT(CTRL + C)
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // block until receive the signal
	log.Infow("Shutting down server......")

	// create a context with a timeout of 10 seconds
	// used to signal server goroutines that they have 10s to complete the current request
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// gracefully shut down the server: handle any unprocessed requests
	if err := httpsrv.Shutdown(ctx); err != nil {
		log.Fatalw("Insecure Server forced to shutdown:", "err", err)
		return err
	}
	if err := httpssrv.Shutdown(ctx); err != nil {
		log.Errorw("Secure Server forced to shutdown:", "err", err)
		return err
	}
	grpcsrv.GracefulStop()

	log.Infow("Server exiting")

	return nil
}

// startInsecureServer create and Run HTTP server
func startInsecureServer(g *gin.Engine) *http.Server {
	// create HTTP server instance
	httpsrv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: g,
	}

	// start the server in a goroutine
	log.Infow("Start to listening the incoming requests on http address", "addr", viper.GetString("addr"))

	fn := func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}
	go fn()

	return httpsrv
}

// startSecureServer create and Run HTTPS server
func startSecureServer(g *gin.Engine) *http.Server {
	// create HTTPS Server instance
	httpssrv := &http.Server{
		Addr:    viper.GetString("tls.addr"),
		Handler: g,
	}

	// start the server in a goroutine
	log.Infow("Start to listening the incoming requests on https address", "addr", viper.GetString("tls.addr"))
	cert, key := viper.GetString("tls.cert"), viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			if err := httpssrv.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
				log.Fatalw(err.Error())
			}
		}()
	}

	return httpssrv
}

// startGRPCServer create and Run gRPC server
func startGRPCServer() *grpc.Server {
	lis, err := net.Listen("tcp", viper.GetString("grpc.addr"))
	if err != nil {
		log.Fatalw("failed to listen", "err", err)
	}

	// create gRPC server instance
	grpcsrv := grpc.NewServer()
	pb.RegisterDazBlogServer(grpcsrv, user.NewUserController(store.S, nil))

	// start the server in a goroutine
	log.Infow("Start to listening the incoming requests on gRPC address", "addr", viper.GetString("grpc.addr"))

	fn := func() {
		if err := grpcsrv.Serve(lis); err != nil {
			log.Fatalw("failed to serve", "err", err)
		}
	}
	go fn()

	return grpcsrv
}
