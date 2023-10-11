// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/Daz-3ux/dBlog/pkg/util/id"
)

type PostM struct {
	ID        int64     `gorm:"column:id;primary_key"` // unique id for the post, server as the primary key
	Username  string    `gorm:"column:username"`       //	author of the post
	PostID    string    `gorm:"column:postID"`         //	unique id for the post, used as a user-friendly ID
	Title     string    `gorm:"column:title"`          //	title of the post
	Content   string    `gorm:"column:content"`        // content of the post
	CreatedAt time.Time `gorm:"column:createdAt"`      // time when the post was created
	UpdatedAt time.Time `gorm:"column:updatedAt"`      // time when the post was updated
}

// TableName sets the insert table name for this struct type
func (p *PostM) TableName() string {
	return "posts"
}

func (p *PostM) BeforeCreate(tx *gorm.DB) error {
	p.PostID = "post-" + id.GenShortID()

	return nil
}
