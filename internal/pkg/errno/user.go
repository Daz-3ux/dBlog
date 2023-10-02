// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package errno

var (
	// ErrUserAlreadyExist represents user already exists
	ErrUserAlreadyExist  = &Errno{HTTP: 400, Code: "FailedOperation.UserAlreadyExist", Message: "User already exists."}
	ErrUserNotFound      = &Errno{HTTP: 404, Code: "ResourceNotFound.UserNotFound", Message: "User not found."}
	ErrPasswordIncorrect = &Errno{HTTP: 401, Code: "InvalidParameter.PasswordIncorrect", Message: "Password incorrect."}
)
