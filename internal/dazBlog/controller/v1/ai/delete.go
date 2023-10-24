// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package ai

import (
	"github.com/gin-gonic/gin"

	"github.com/Daz-3ux/dBlog/internal/pkg/core"
	"github.com/Daz-3ux/dBlog/internal/pkg/known"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
)

func (ctrl *AIController) Delete(c *gin.Context) {
	log.C(c).Infow("Delete post function called")

	postID := c.Param("postID")
	if err := ctrl.b.AIs().Delete(c, c.GetString(known.XUsernameKey), postID); err != nil {
		core.WriteResponse(c, err, map[string]string{"delete post_id": "failed"})

		return
	}

	core.WriteResponse(c, nil, map[string]string{"delete post_id": postID, "status": "ok"})
}
