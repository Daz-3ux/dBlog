// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package core

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Daz-3ux/dBlog/internal/pkg/errno"
)

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// WriteResponse writes an error or response to the HTTP response body
// WriteResponse uses the errno.Decode method to extract the business error code and error message from the error based on the error type
func WriteResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		if err == errno.ErrPageNotFound {
			c.HTML(http.StatusNotFound, "404.html", gin.H{})
			return
		} else {
			hcode, code, message := errno.Decode(err)
			c.JSON(hcode, ErrResponse{
				Code:    code,
				Message: message,
			})
		}

		return
	}

	c.JSON(http.StatusOK, data)
}
