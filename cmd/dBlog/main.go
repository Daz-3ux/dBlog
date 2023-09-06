// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package main

import (
	"os"

	_ "go.uber.org/automaxprocs/maxprocs"

	"github.com/Daz-3ux/dBlog/internal/dazBlog"
)

// program entry function
func main() {
	command := dazBlog.NewDazBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
