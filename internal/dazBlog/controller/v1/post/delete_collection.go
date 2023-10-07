// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package post

import (
	"github.com/Daz-3ux/dBlog/internal/pkg/core"
	"github.com/Daz-3ux/dBlog/internal/pkg/known"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

func (ctrl *PostController) DeleteCollection(c *gin.Context) {
	log.C(c).Infow("Batch delete post function called")

	postIDs := c.QueryArray("postID")
	if err := ctrl.b.Posts().DeleteCollection(c, c.GetString(known.XUsernameKey), postIDs); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
