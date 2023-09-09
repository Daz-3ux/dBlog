// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package dazBlog

import (
	"encoding/json"
	"fmt"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// NewDazBlogCommand create the *cobra.Command object
// so can use the Execute method to start
func NewDazBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		// specify the name of the command
		Use: "dBlog",
		// specify the short description of the command
		Short: "dBlog is a simple blog system",
		// specify the long description of the command
		Long: `dBlog is a simple but not easy blog system,
Find more dBlog information at:
	https://github.com/daz-3ux/dBlog#readme`,

		// when error occurs, the command will not print usage information
		SilenceUsage: true,
		// specify the run function to execute when cmd.Execute() is called
		// if the function fails, an error message will be returned
		RunE: func(cmd *cobra.Command, args []string) error {
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

	// cobra support persistent flag
	// for providing options that work across all sub-commands of a command
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the dBlog configuration file. Empty string for no configuration file.")

	// cobra supports local flag
	// which can only be used within the command to which they are bound
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	return cmd
}

func run() error {
	fmt.Println("Hello, dBlog!")
	settings, _ := json.Marshal(viper.AllSettings())
	log.Infow(string(settings))
	log.Infow(viper.GetString("db.username"))
	return nil
}
