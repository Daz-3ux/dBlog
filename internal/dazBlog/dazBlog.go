// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package dazBlog

import (
	"context"
	"errors"
	"fmt"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	mw "github.com/Daz-3ux/dBlog/internal/pkg/middleware"
	"github.com/Daz-3ux/dBlog/pkg/version/verflag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var cfgFile string

// NewDazBlogCommand create the *cobra.Command object
// so can use the Execute method to start
func NewDazBlogCommand() *cobra.Command {
	// cmd is the *cobra.Command object and the top-level command
	cmd := &cobra.Command{
		// specify the name of the command
		Use: "dBlog",
		// specify the short description of the command
		Short: "dBlog is a simple blog system",
		// specify the long description of the command
		Long: `dBlog is a simple but not easy blog system,
Find more dBlog information at:
	https://github.com/daz-3ux/dBlog#readme`,

		// when an error occurs, the command will not print usage information
		SilenceUsage: true,
		// specify the run function to execute when cmd.Execute() is called
		// if the function fails, an error message will be returned
		RunE: func(cmd *cobra.Command, args []string) error {
			verflag.PrintAndExitIfRequested()

			log.Init(logOptions())
			defer log.Sync()

			return run()
		},
		// no need to specify command line parameters
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
	// set Gin mode
	gin.SetMode(viper.GetString("runmode"))

	// create Gin engine
	g := gin.New()

	// gin.Recovery() middleware is used to capture any panics and recover from them
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.Secure, mw.RequestID()}

	g.Use(mws...)

	// register 404 handler
	g.LoadHTMLGlob("internal/resource/*.html")
	g.NoRoute(func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{"code": 10003, "message": "404 Not Found"})
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	// register /healthz handler
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	// create HTTP server
	httpsrv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: g,
	}

	// start HTTP server
	log.Infow("Start HTTP listening", "address", httpsrv.Addr)
	// anonymous goroutine, listen and serve on HTTPS
	go func() {
		if err := httpsrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalw(err.Error())
		}
	}()

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

	log.Infow("Server exiting")

	return nil
}
