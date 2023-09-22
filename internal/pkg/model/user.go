package model

import "time"

type UserM struct {
	ID        int64     `gorm:"column:id;primary_key"` // unique id for the user, server as the primary key
	Username  string    `gorm:"column:username"`       // username of the user
	Password  string    `gorm:"column:password"`       // password of the user
	Nickname  string    `gorm:"column:nickname"`       // nickname of the user
	Email     string    `gorm:"column:email"`          // email of the user
	Phone     string    `gorm:"column:phone"`          // phone number of the user
	CreatedAt time.Time `gorm:"column:createdAt"`      // time when the user was created
	UpdatedAt time.Time `gorm:"column:updatedAt"`      // time when the user was updated
}

// TableName sets the insert table name for this struct type
func (u *UserM) TableName() string {
	return "users"
}
