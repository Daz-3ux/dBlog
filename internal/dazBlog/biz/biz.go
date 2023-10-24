// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package biz

//go:generate mockgen -destination mock_biz.go -package biz github.com/Daz-3ux/dBlog/internal/dazBlog/biz IBiz

import (
	"github.com/Daz-3ux/dBlog/internal/dazBlog/biz/ai"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/biz/post"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/biz/user"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
)

// IBiz defines the method need to be implemented by the Biz layer
type IBiz interface {
	Users() user.UserBiz
	Posts() post.PostBiz
	AIs() ai.AIBiz
}

// ensure that biz implements the IBiz interface
var _ IBiz = (*biz)(nil)

// biz implements the IBiz interface
// biz layer is connected to the store layer in the under
type biz struct {
	ds store.IStore
}

// NewBiz create an instance of type IBiz
func NewBiz(ds store.IStore) IBiz {
	return &biz{ds: ds}
}

// Users return an instance of UserBiz
func (b *biz) Users() user.UserBiz {
	return user.NewUserBiz(b.ds)
}

func (b *biz) Posts() post.PostBiz {
	return post.NewPostBiz(b.ds)
}

func (b *biz) AIs() ai.AIBiz {
	return ai.NewAIBiz(b.ds)
}
