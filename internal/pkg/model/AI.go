// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package model

type AIM struct {
	ID       int64  `gorm:"column:id;primary_key"`     // unique id of the ai content, server as the primary key
	Username string `gorm:"column:username"`           // author of the ai content
	PostID   string `gorm:"column:postID;primary_key"` // post of the ai content
	Content  string `gorm:"column:content"`            // content of the ai content
}

func (u *AIM) TableName() string {
	return "ai"
}
