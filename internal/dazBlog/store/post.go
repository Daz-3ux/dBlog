// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package store

import (
	"context"
	"github.com/Daz-3ux/dBlog/internal/pkg/model"
	"gorm.io/gorm"
)

// PostStore defines the methods that need to be implemented by the post model in the store layer
type PostStore interface {
	Create(ctx context.Context, user *model.PostM) error
}

// posts is the implementation of PostStore interface
type posts struct {
	db *gorm.DB
}

// ensure that posts implement the PostStore interface
var _ PostStore = (*posts)(nil)

func newPosts(db *gorm.DB) *posts {
	return &posts{db}
}

func (p *posts) Create(ctx context.Context, post *model.PostM) error {
	return p.db.Create(&post).Error
}
