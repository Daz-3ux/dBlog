// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package v1

type AIInfo struct {
	Username string `gorm:"column:username"` // username of the AI content
	PostID   string `gorm:"column:postId"`   // post of the AI content
	Content  string `gorm:"column:content"`  // content of the AI content
}

type CreateAIRequest AIInfo

type UpdateAIRequest AIInfo

type GetAIResponse AIInfo

type ListAIRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type ListAIResponse struct {
	TotalCount int64     `json:"total_count,omitempty"`
	AIs        []*AIInfo `json:"AIs,omitempty"`
}
