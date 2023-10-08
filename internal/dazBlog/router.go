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
	"github.com/Daz-3ux/dBlog/pkg/auth"
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

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	uc := user.NewUserController(store.S, authz)
	pc := post.NewPostController(store.S)

	g.POST("/login", uc.Login)
	// create v1 route group
	v1 := g.Group("/v1")
	{
		// create user's route group
		userv1 := v1.Group("/users")
		{
			// create user
			userv1.POST("", uc.Create)
			// change user password
			userv1.PUT(":name/change-password", uc.ChangePassword)
			// middleware
			userv1.Use(mw.Authn(), mw.Authz(authz))
			// get user info
			userv1.GET(":name", uc.Get)
			// update user info
			userv1.PUT(":name", uc.Update)
			// list all users, only root user can access
			userv1.GET("", uc.List)
			// delete user
			userv1.DELETE(":name", uc.Delete)
		}

		postv1 := v1.Group("/posts")
		{
			// create post
			postv1.POST("", pc.Create)
			// get post
			postv1.GET(":postID", pc.Get)
			// update post
			postv1.PUT(":postID", pc.Update)
			// delete post
			postv1.DELETE(":postID", pc.Delete)
			// batch delete posts
			postv1.DELETE("", pc.DeleteCollection)
			// list posts
			postv1.GET("", pc.List)
		}
	}

	return nil
}
