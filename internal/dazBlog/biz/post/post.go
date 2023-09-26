// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package post

import (
	"context"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
	"github.com/Daz-3ux/dBlog/internal/pkg/errno"
	"github.com/Daz-3ux/dBlog/internal/pkg/model"
	v1 "github.com/Daz-3ux/dBlog/pkg/api/dazBlog/v1"
	"github.com/jinzhu/copier"
	"regexp"
)

// PostBiz defines the methods implemented in the post module at the biz layer
type PostBiz interface {
	Create(ctx context.Context, r *v1.CreatePostRequest) error
}

// postBiz implements the PostBiz interface
type postBiz struct {
	ds store.IStore
}

// ensure that postBiz implements the PostBiz interface
var _ PostBiz = (*postBiz)(nil)

func NewPostBiz(ds store.IStore) PostBiz {
	return &postBiz{ds: ds}
}

func (p *postBiz) Create(ctx context.Context, r *v1.CreatePostRequest) error {
	var postM model.PostM
	_ = copier.Copy(&postM, r)

	if err := p.ds.Posts().Create(ctx, &postM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'title'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}

	return nil
}
