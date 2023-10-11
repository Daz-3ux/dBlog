// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package post

import (
	"github.com/gin-gonic/gin"

	"github.com/Daz-3ux/dBlog/internal/pkg/core"
	"github.com/Daz-3ux/dBlog/internal/pkg/known"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
)

func (ctrl *PostController) Get(c *gin.Context) {
	log.C(c).Infow("Get post function called")

	postID := c.Param("postID")
	post, err := ctrl.b.Posts().Get(c, c.GetString(known.XUsernameKey), postID)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, post)
}
