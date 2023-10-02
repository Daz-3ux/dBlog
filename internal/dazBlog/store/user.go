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

// UserStore defines the methods that need to be implemented by the user model in the store layer
type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
	Get(ctx context.Context, username string) (*model.UserM, error)
	Update(ctx context.Context, user *model.UserM) error
}

// users is the implementation of UserStore interface
type users struct {
	db *gorm.DB
}

// ensure that users implement the UserStore interface
var _ UserStore = (*users)(nil)

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

// Create insects a new user record into the database
func (u *users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}

// Get returns the user record with the specified username
func (u *users) Get(ctx context.Context, username string) (*model.UserM, error) {
	var user model.UserM
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates the user record in the database
func (u *users) Update(ctx context.Context, user *model.UserM) error {
	return u.db.Save(&user).Error
}
