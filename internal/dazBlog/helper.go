// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package dazBlog

import (
	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	"github.com/Daz-3ux/dBlog/pkg/db"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const (
	// recommendedHomeDir defines the default directory for storing dBlog service configurations
	recommendedHomeDir = ".dblog "
	// defaultConfigName specifies the default configuration file name for dBlog service
	defaultConfigName = "dazBlog.yaml"
)

// initConfig set the config file name to be read, env variable, and read config file content into viper
func initConfig() {
	if cfgFile != "" {
		// read config file from the specified file
		viper.SetConfigFile(cfgFile)
	} else {
		// search user's homedir
		home, err := os.UserHomeDir()
		// if failed, print 'Error: xxx' and do exit(1)
		cobra.CheckErr(err)

		// add `$HOME/<recommendedHomeDir>` into search path
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))
		// add current dir into search path
		viper.AddConfigPath(".")
		// set the config file's type
		viper.SetConfigType("yaml")
		// set the config file's name
		viper.SetConfigName(defaultConfigName)
	}

	// read matching env variable
	viper.AutomaticEnv()
	// set env variable's prefix
	viper.SetEnvPrefix("MINIBLOG")

	// Replace '.' and '-' with '_' in the key string before calling viper.Get(key).
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read config file", "err", err)
	}

	// print the currently used config file
	log.Infow("Using config file", "file", viper.ConfigFileUsed())
}

// logOptions reads log config from viper
func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}

// initStore read the DB config from viper and init the store layer
func initStore() error {
	// init the store layer
	dbOptions := &db.MySQLOptions{
		Host:                  viper.GetString("db.host"),
		Username:              viper.GetString("db.username"),
		Password:              viper.GetString("db.password"),
		Database:              viper.GetString("db.database"),
		MaxIdleConnections:    viper.GetInt("db.max-idle-connections"),
		MaxOpenConnections:    viper.GetInt("db.max-open-connections"),
		MaxConnectionLifeTime: viper.GetDuration("db.max-connection-life-time"),
		LogLevel:              viper.GetInt("db.log-level"),
	}

	ins, err := db.NewMySql(dbOptions)
	if err != nil {
		return err
	}

	_ = store.NewStore(ins)

	return nil
}
