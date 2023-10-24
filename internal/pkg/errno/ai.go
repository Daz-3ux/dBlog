// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package errno

// ErrAICommentAlreadyExist is the error when the post-comment already exists
var ErrAICommentAlreadyExist = &Errno{HTTP: 400, Code: "ResourceAlreadyExist.CommentExisted", Message: "comment already existed"}
