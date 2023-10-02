// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package dazBlog

import (
	"github.com/Daz-3ux/dBlog/internal/dazBlog/controller/v1/post"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/controller/v1/user"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
	"github.com/Daz-3ux/dBlog/internal/pkg/core"
	"github.com/Daz-3ux/dBlog/internal/pkg/errno"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	mw "github.com/Daz-3ux/dBlog/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func installRouters(g *gin.Engine) error {
	// register 404 handler
	g.LoadHTMLGlob("internal/resource/404.html")
	g.NoRoute(func(c *gin.Context) {
		log.C(c).Errorw("Page not found", "url:", c.Request.URL)
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// register /healthz handler
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")
		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	uc := user.NewUserController(store.S)
	pc := post.NewPostController(store.S)

	g.POST("/login", uc.Login)
	// create v1 route group
	v1 := g.Group("/v1")
	{
		// create user's route group
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(mw.Authn())
		}
		postv1 := v1.Group("/posts")
		{
			postv1.POST("", pc.Create)
		}
	}

	return nil
}
