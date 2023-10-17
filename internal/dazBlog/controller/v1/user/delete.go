// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package user

import (
	"github.com/gin-gonic/gin"

	"github.com/Daz-3ux/dBlog/internal/pkg/core"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
)

func (ctrl *UserController) Delete(c *gin.Context) {
	log.C(c).Infow("Delete user function called")

	username := c.Param("name")

	if err := ctrl.b.Users().Delete(c, username); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	if _, err := ctrl.a.RemoveNamedPolicy("p", username, "/v1/users/"+username, defaultMethods); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, map[string]string{"delete user": "OK"})
}
