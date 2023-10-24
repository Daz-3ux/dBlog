// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package ai

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/Daz-3ux/dBlog/internal/pkg/core"
	"github.com/Daz-3ux/dBlog/internal/pkg/errno"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	v1 "github.com/Daz-3ux/dBlog/pkg/api/dazBlog/v1"
)

func (ctrl *AIController) Create(c *gin.Context) {
	log.C(c).Infow("Create AI comment function called")

	var r v1.CreateAIRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, map[string]string{"json bind": "failed"})
		log.C(c).Errorw("Create AI comment function called failed", "error", err.Error())

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.AIs().Create(c, &r); err != nil {
		core.WriteResponse(c, err, map[string]string{"create AI comment": "failed"})

		return
	}
}
