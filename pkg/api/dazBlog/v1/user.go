// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package v1

type CreateUserRequest struct {
	Username string `json:"username" valid:"alphanum,required,stringlength(1|255)"`
	Password string `json:"password" valid:"required,stringlength(6|18)"`
	Nickname string `json:"nickname" valid:"required,stringlength(1|255)"`
	Email    string `json:"email" valid:"required,email"`
	Phone    string `json:"phone" valid:"required,numeric,stringlength(11|11)"`
}
