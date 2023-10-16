// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package v1

// CreateUserRequest specifies the request parameters for
// `POST /v1/users`
type CreateUserRequest struct {
	Postcount int64  `json:"postcount" valid:"stringlength(1|255)"`
	Username  string `json:"username" valid:"alphanum,required,stringlength(1|30)"`
	Password  string `json:"password" valid:"required,stringlength(6|18)"`
	Nickname  string `json:"nickname" valid:"required,stringlength(1|30)"`
	Email     string `json:"email" valid:"required,email"`
	Gender    string `json:"gender" valid:"required"`
	Phone     string `json:"phone" valid:"required,numeric,stringlength(11|11)"`
	QQ        string `json:"qq" valid:"numeric,stringlength(5|16)"`
}

// GetUserResponse specifies the response parameters for
// `GET /v1/users/{username}`
type GetUserResponse UserInfo

// UserInfo is the user's all information that can be listed
type UserInfo struct {
	PostCount int64  `json:"postcount"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	QQ        string `json:"qq"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// ListUserRequest specifies the request parameters for
// `GET /v1/users`
type ListUserRequest struct {
	Offset int `form:"offset" valid:"numeric"`
	Limit  int `form:"limit" valid:"numeric"`
}

// ListUserResponse specifies the response parameters for
// `GET /v1/users`
type ListUserResponse struct {
	TotalCount int64       `json:"totalCount" valid:"numeric"`
	Users      []*UserInfo `json:"users"`
}

// UpdateUserRequest specifies the request parameters for
// `PUT /v1/users/{username}`
type UpdateUserRequest struct {
	Nickname *string `json:"nickname" valid:"stringlength(1|255)"`
	Email    *string `json:"email" valid:"email"`
	Gender   *string `json:"gender" valid:"stringlength(1|10)"`
	Phone    *string `json:"phone" valid:"numeric,stringlength(11|11)"`
	QQ       *string `json:"qq" valid:"numeric,stringlength(5|16)"`
}

// ChangePasswordRequest specifies the request parameters for
// `POST /v1/users/{username}/change-password`
type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" valid:"required,stringlength(6|18)"`
	NewPassword string `json:"newPassword" valid:"required,stringlength(6|18)"`
}

// LoginRequest specifies the request parameters for
// `POST /login`
type LoginRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
}

// LoginResponse specifies the response parameters for
// `POST /login`
type LoginResponse struct {
	Token string `json:"token"`
}
