// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package user

import (
	"github.com/Daz-3ux/dBlog/internal/dazBlog/biz"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
)

type UserController struct {
	b biz.IBiz
}

func NewUserController(ds store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}
