// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package dazBlog

import (
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const (
	recommendedHomeDir = ".dazblog "
	defaultConfigName  = "dazBlog.yaml"
)

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(defaultConfigName)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("MINIBLOG")

	replacer := strings.NewReplacer(".", "-")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read config file", "err", err)
	}

	log.Infow("Using config file", "file", viper.ConfigFileUsed())
}

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
