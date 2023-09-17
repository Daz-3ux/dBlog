// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package errno

var (
	// OK represents success
	OK = &Errno{HTTP: 200, Code: "", Message: ""}

	// InternalServerError represents all unknown server errors
	InternalServerError = &Errno{HTTP: 500, Code: "InternalError", Message: "Internal Server Error"}

	ErrPageNotFound = &Errno{HTTP: 404, Code: "ResourceNotFound.PageNotFound", Message: "Page Not Found"}
)
