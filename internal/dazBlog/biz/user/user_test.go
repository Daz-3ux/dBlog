// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package user

import (
	"context"
	"fmt"
	"github.com/AlekSi/pointer"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
	"github.com/Daz-3ux/dBlog/internal/pkg/errno"
	"github.com/Daz-3ux/dBlog/internal/pkg/model"
	v1 "github.com/Daz-3ux/dBlog/pkg/api/dazBlog/v1"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func fakeUser(id int64) *model.UserM {
	return &model.UserM{
		ID:        id,
		PostCount: 10,
		Username:  fmt.Sprintf("daz%d", id),
		Password:  fmt.Sprintf("daz%d", id),
		Nickname:  fmt.Sprintf("daz%d", id),
		Email:     "daz@qq.com",
		Phone:     "18139238989",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Test_NewUserBiz tests the NewUserBiz function
func Test_NewUserBiz(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock an interface -- IStore, which is required by NewUserBiz
	mockStore := store.NewMockIStore(ctrl)

	type args struct {
		ds store.IStore
	}
	tests := []struct {
		name string
		args args
		want *userBiz
	}{
		{name: "default", args: args{mockStore}, want: &userBiz{mockStore}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUserBiz(tt.args.ds)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestUserBiz_Create tests the Create method of the UserBiz interface
func TestUserBiz_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// create a mockUserStore and specify the expected behavior
	// When the Create method is called, it returns nil
	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	// create a mockStore and specify the expected behavior
	// When the Users method is called, it returns the mockUserStore
	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().Return(mockUserStore).AnyTimes()

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx context.Context
		r   *v1.CreateUserRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "default", fields: fields{mockStore}, args: args{context.Background(), nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			/*
			   当调用 b.Create 方法时, 它会间接地调用 mockStore.Users().Create 方法
			   因为我们已经设置了 mockStore.Users() 的预期行为并返回 mockUserStore
			   所以实际上调用的是 mockUserStore.Create 方法,返回 nil (预期行为)
			*/
			assert.Nil(t, b.Create(tt.args.ctx, tt.args.r))
		})
	}
}

// TestUserBiz_Get tests the Get method of the UserBiz interface
func TestUserBiz_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeUser := fakeUser(1)
	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(fakeUser, nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().Return(mockUserStore).AnyTimes()

	var want v1.GetUserResponse
	_ = copier.Copy(&want, fakeUser)
	want.CreatedAt = fakeUser.CreatedAt.Format("2006-01-02 15:04:05")
	want.UpdatedAt = fakeUser.UpdatedAt.Format("2006-01-02 15:04:05")

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.GetUserResponse
	}{
		{name: "default", fields: fields{mockStore}, args: args{context.Background(), "daz1"}, want: &want},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			got, err := b.Get(tt.args.ctx, tt.args.username)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestUserBiz_List tests the List method of the UserBiz interface
func TestUserBiz_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeUsers := []*model.UserM{fakeUser(1), fakeUser(2), fakeUser(3)}
	wantUsers := make([]*v1.UserInfo, 0, len(fakeUsers))
	for _, u := range fakeUsers {
		wantUsers = append(wantUsers, &v1.UserInfo{
			Username:  u.Username,
			Nickname:  u.Nickname,
			Email:     u.Email,
			Phone:     u.Phone,
			PostCount: 10,
			CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(5), fakeUsers, nil).Times(1)

	mockPostStore := store.NewMockPostStore(ctrl)
	mockPostStore.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(10), nil, nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().Return(mockUserStore).Times(1)
	mockStore.EXPECT().Posts().Return(mockPostStore).AnyTimes()

	tests := []struct {
		name    string
		want    *v1.ListUserResponse
		wantErr bool
	}{
		{name: "default", want: &v1.ListUserResponse{TotalCount: 5, Users: wantUsers}, wantErr: false},
	}

	ub := NewUserBiz(mockStore)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ub.List(context.Background(), 0, 10)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestUserBiz_Update tests the Update method of the UserBiz interface
func TestUserBiz_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeUser := fakeUser(1)
	r := &v1.UpdateUserRequest{
		Nickname: pointer.ToString("ddaazz"),
		Email:    pointer.ToString("qq@daz.com"),
		Phone:    pointer.ToString("12312312312"),
	}
	wantedUser := *fakeUser
	wantedUser.Nickname = *r.Nickname
	wantedUser.Email = *r.Email
	wantedUser.Phone = *r.Phone

	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(fakeUser, nil).AnyTimes()
	// if update successfully, return nil
	mockUserStore.EXPECT().Update(gomock.Any(), &wantedUser).Return(nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().Return(mockUserStore).AnyTimes()

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx      context.Context
		username string
		user     *v1.UpdateUserRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "default", fields: fields{mockStore}, args: args{context.Background(), "daz1", r}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			err := b.Update(tt.args.ctx, tt.args.username, tt.args.user)
			assert.Nil(t, err)
		})
	}
}

// TestUserBiz_ChangePassword tests the ChangePassword method of the UserBiz interface
func TestUserBiz_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().Return(mockUserStore).AnyTimes()

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "default", fields: fields{mockStore}, args: args{context.Background(), "daz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			assert.Nil(t, b.Delete(tt.args.ctx, tt.args.username))
		})
	}
}

// TestUserBiz_ChangePassword tests the ChangePassword method of the UserBiz interface
func TestUserBiz_ChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeUser := fakeUser(1)
	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(fakeUser, nil).AnyTimes()
	mockUserStore.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().Return(mockUserStore).AnyTimes()

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx      context.Context
		username string
		r        *v1.ChangePasswordRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "default",
			fields: fields{mockStore},
			args:   args{context.Background(), "daz", &v1.ChangePasswordRequest{"daz1", "daz123"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			err := b.ChangePassword(tt.args.ctx, tt.args.username, tt.args.r)
			assert.Equal(t, errno.ErrPasswordIncorrect, err)
		})
	}
}

// TestUserBiz_Login tests the Login method of the UserBiz interface
func TestUserBiz_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeUser := fakeUser(1)
	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(fakeUser, nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().AnyTimes().Return(mockUserStore)

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx context.Context
		r   *v1.LoginRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *errno.Errno
	}{
		{
			name:   "default",
			fields: fields{mockStore},
			args:   args{context.Background(), &v1.LoginRequest{Username: "daz1", Password: "daz1"}},
			want:   errno.ErrPasswordIncorrect,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			got, err := b.Login(tt.args.ctx, tt.args.r)
			assert.Nil(t, got)
			assert.Equal(t, tt.want, err)
		})
	}
}
