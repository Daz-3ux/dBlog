// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/Daz-3ux/dBlog/internal/pkg/core"
	"github.com/Daz-3ux/dBlog/internal/pkg/errno"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	v1 "github.com/Daz-3ux/dBlog/pkg/api/dazBlog/v1"
)

func (ctrl *UserController) Update(c *gin.Context) {
	log.C(c).Infow("Update user function called")

	var r v1.UpdateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		log.C(c).Errorw("binding request failed", "error", err.Error())
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.Users().Update(c, c.Param("name"), &r); err != nil {
		core.WriteResponse(c, err, nil)
		log.C(c).Errorw("updating user failed", "error", err.Error())
		return
	}

	core.WriteResponse(c, nil, map[string]string{"update user info": "ok"})
}
