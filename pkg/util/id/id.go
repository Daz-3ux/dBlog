// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package id

import (
	shortid "github.com/jasonsoft/go-short-id"
)

func GenShortID() string {
	opt := shortid.Options{
		Number:        4,
		StartWithYear: true,
		EndWithHost:   false,
	}

	return toLower(shortid.Generate(opt))
}

func toLower(ss string) string {
	var lower string
	for _, s := range ss {
		lower += string(s | ' ')
	}

	return lower
}
