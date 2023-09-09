// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package log

import (
	"go.uber.org/zap/zapcore"
)

// Options the options for the logger
type Options struct {
	// if enabled, the log will display the file and line num of the calling log
	DisableCaller bool
	// if enabled, will print stack info for panic and higher log level
	DisableStacktrace bool
	// log level: debug, info , warn, error, dpanic, panic, fatal
	Level string
	// log format: console, json
	Format string
	// specify log output path
	OutputPaths []string
}

// NewOptions creates an Options object with default parameters
func NewOptions() *Options {
	return &Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             zapcore.InfoLevel.String(),
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	}
}
