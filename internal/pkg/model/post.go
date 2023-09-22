package model

import "time"

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
