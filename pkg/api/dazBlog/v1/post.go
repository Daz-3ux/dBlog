// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package v1

// CreatePostRequest specifies the request parameters for
// `POST /v1/posts`
type CreatePostRequest struct {
	Title   string `json:"title" valid:"required,stringlength(1|255)"`
	Content string `json:"content" valid:"required"`
}

// CreatePostResponse specifies the response parameters for
// `POST /v1/posts`
type CreatePostResponse struct {
	PostID string `json:"post_id"`
}

// GetPostResponse specifies the request parameters for
// `GET /v1/posts/{post_id}`
type GetPostResponse PostInfo

// PostInfo is the post's detail info
type PostInfo struct {
	Username  string `json:"username,omitempty"`
	PostID    string `json:"post_id,omitempty"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// UpdatePostRequest specifies the request parameters for
// `PUT /v1/posts`
type UpdatePostRequest struct {
	Title   *string `json:"title" valid:"stringlength(1|255)"`
	Content *string `json:"content" valid:"stringlength(1|65535)"`
}

// ListPostsRequest specifies the request parameters for
// `GET /v1/posts`
type ListPostsRequest struct {
	Offset int `form:"offset"`
	Limit  int `form:"limit"`
}

// ListPostsResponse specifies the response parameters for
// `GET /v1/posts`
type ListPostsResponse struct {
	TotalCount int64       `json:"totalCount"`
	Posts      []*PostInfo `json:"posts"`
}
