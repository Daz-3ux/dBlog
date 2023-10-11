// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/Daz-3ux/dBlog/internal/pkg/known"
)

// RequestID is a Gin middleware used to inject the `X-Request-ID` key-value pair into the context and response of each HTTP request
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the request id from the Gin context
		requestID := c.Request.Header.Get(known.XRequestIDKey)

		// If the request id is empty, set a new one
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Set the request id to the Gin context
		c.Set(known.XRequestIDKey, requestID)

		// Set the request id to the response header
		c.Writer.Header().Set(known.XRequestIDKey, requestID)

		// Continue
		c.Next()
	}
}
