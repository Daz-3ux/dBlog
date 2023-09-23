// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package user

import (
	"context"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
	"github.com/Daz-3ux/dBlog/internal/pkg/errno"
	"github.com/Daz-3ux/dBlog/internal/pkg/model"
	v1 "github.com/Daz-3ux/dBlog/pkg/api/dazBlog/v1"
	"github.com/jinzhu/copier"
	"regexp"
)

// UserBiz defines the methods implemented in the user module at the biz layer
// implement the specific implementations of the REST resources for the user
type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
}

// userBiz implements the UserBiz interface
type userBiz struct {
	ds store.IStore
}

// ensure that userBiz implements the UserBiz interface
var _ UserBiz = (*userBiz)(nil)

// NewUserBiz create an instance of type UserBiz
func NewUserBiz(ds store.IStore) UserBiz {
	return &userBiz{ds: ds}
}

// Create is the implementation of the `Create` method of the UserBiz interface
func (b *userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, r)

	if err := b.ds.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}

		return err
	}

	return nil

}
