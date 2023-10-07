// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package user

import (
	"context"
	"errors"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
	"github.com/Daz-3ux/dBlog/internal/pkg/errno"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	"github.com/Daz-3ux/dBlog/internal/pkg/model"
	v1 "github.com/Daz-3ux/dBlog/pkg/api/dazBlog/v1"
	"github.com/Daz-3ux/dBlog/pkg/auth"
	"github.com/Daz-3ux/dBlog/pkg/token"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"regexp"
)

// UserBiz defines the methods implemented in the user module at the biz layer
// implement the specific implementations of the REST resources for the user
type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
	Get(ctx context.Context, id string) (*v1.GetUserResponse, error)
	List(ctx context.Context, offset, limit int) (*v1.ListUserResponse, error)
	Update(ctx context.Context, username string, r *v1.UpdateUserRequest) error
	Delete(ctx context.Context, username string) error
	ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error
	Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error)
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

// Get is the implementation of the `Get` method in the UserBiz interface
func (b *userBiz) Get(ctx context.Context, id string) (*v1.GetUserResponse, error) {
	user, err := b.ds.Users().Get(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}

		return nil, err
	}

	var resp v1.GetUserResponse
	_ = copier.Copy(&resp, user)

	resp.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
	resp.UpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")

	return &resp, nil
}

// List is the implementation of the `List` method of the UserBiz interface
func (b *userBiz) List(ctx context.Context, offset, limit int) (*v1.ListUserResponse, error) {
	count, list, err := b.ds.Users().List(ctx, offset, limit)
	if err != nil {
		log.C(ctx).Errorw("failed to list users from storage", "err", err)
		return nil, err
	}

	users := make([]*v1.UserInfo, 0, len(list))
	for _, user := range list {
		users = append(users, &v1.UserInfo{
			Username:  user.Username,
			Nickname:  user.Nickname,
			Email:     user.Email,
			Phone:     user.Phone,
			PostCount: user.PostCount,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	log.C(ctx).Debugw("Get users from storage", "count", count)

	return &v1.ListUserResponse{TotalCount: count, Users: users}, nil
}

// Update is the implementation of the `Update` method of the UserBiz interface
func (b *userBiz) Update(ctx context.Context, username string, user *v1.UpdateUserRequest) error {
	userM, err := b.ds.Users().Get(ctx, username)
	if err != nil {
		return err
	}

	if user.Email != "" {
		userM.Email = user.Email
	}
	if user.Nickname != "" {
		userM.Nickname = user.Nickname
	}
	if user.Phone != "" {
		userM.Phone = user.Phone
	}

	if err := b.ds.Users().Update(ctx, userM); err != nil {
		return err
	}

	return nil
}

func (b *userBiz) Delete(ctx context.Context, username string) error {
	if err := b.ds.Users().Delete(ctx, username); err != nil {
		return err
	}

	return nil
}

// ChangePassword is the implementation of the `ChangePassword` method of the UserBiz interface
func (b *userBiz) ChangePassword(ctx context.Context, username string, r *v1.ChangePasswordRequest) error {
	userM, err := b.ds.Users().Get(ctx, username)
	if err != nil {
		return err
	}

	if err := auth.Compare(userM.Password, r.OldPassword); err != nil {
		return errno.ErrPasswordIncorrect
	}

	userM.Password, _ = auth.Encrypt(r.NewPassword)
	if err := b.ds.Users().Update(ctx, userM); err != nil {
		return err
	}

	return nil
}

// Login is the implementation of the `Login` method of the UserBiz interface
func (b *userBiz) Login(ctx context.Context, r *v1.LoginRequest) (*v1.LoginResponse, error) {
	// get all infos of the logged-in user
	user, err := b.ds.Users().Get(ctx, r.Username)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}

	// compare the password
	if err := auth.Compare(user.Password, r.Password); err != nil {
		return nil, errno.ErrPasswordIncorrect
	}

	// if the password is correct, generate a JWT token
	t, err := token.Sign(r.Username)
	if err != nil {
		return nil, errno.ErrSignToken
	}

	return &v1.LoginResponse{Token: t}, nil
}
