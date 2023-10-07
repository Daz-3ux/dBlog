// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package store

import (
	"context"
	"errors"
	"github.com/Daz-3ux/dBlog/internal/pkg/model"
	"gorm.io/gorm"
)

// PostStore defines the methods that need to be implemented by the post model in the store layer
type PostStore interface {
	Create(ctx context.Context, user *model.PostM) error
	Get(ctx context.Context, username string, postID string) (*model.PostM, error)
	Update(ctx context.Context, user *model.PostM) error
	List(ctx context.Context, username string, offset, limit int) (int64, []*model.PostM, error)
	Delete(ctx context.Context, username string, postIDs []string) error
}

// posts is the implementation of PostStore interface
type posts struct {
	db *gorm.DB
}

// ensure that posts implement the PostStore interface
var _ PostStore = (*posts)(nil)

func newPosts(db *gorm.DB) *posts {
	return &posts{db}
}

// Create a new post
func (p *posts) Create(ctx context.Context, post *model.PostM) error {
	return p.db.Create(&post).Error
}

// Get a post by username and postID
func (p *posts) Get(ctx context.Context, username string, postID string) (*model.PostM, error) {
	var post model.PostM
	if err := p.db.Where("username = ? AND postID = ?", username, postID).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// Update a post
func (p *posts) Update(ctx context.Context, post *model.PostM) error {
	return p.db.Save(post).Error
}

func (p *posts) List(ctx context.Context, username string, offset, limit int) (count int64, ret []*model.PostM, err error) {
	err = p.db.Where("username = ?", username).
		// set the offset and limit
		Offset(offset).
		Limit(defaultLimit(limit)).
		// descending order the results by id
		Order("id desc").
		// store the result ti ret
		Find(&ret).
		// reset the offset and limit
		Offset(-1).
		Limit(-1).
		// calculate the total number of results and store to count
		Count(&count).
		Error

	return
}

func (p *posts) Delete(ctx context.Context, username string, postIDs []string) error {
	err := p.db.Where("username = ? AND postID in (?)", username, postIDs).Delete(&model.PostM{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}
