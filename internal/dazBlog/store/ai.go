// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package store

import (
	"context"

	"gorm.io/gorm"

	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	"github.com/Daz-3ux/dBlog/internal/pkg/model"
)

type AIStore interface {
	Create(ctx context.Context, ai *model.AIM) error
	Get(ctx context.Context, username string, postID string) (*model.AIM, error)
	Update(ctx context.Context, ai *model.AIM) error
	List(ctx context.Context, username string, offset, limit int) (int64, []*model.AIM, error)
	Delete(ctx context.Context, username string, postID string) error
}

type ais struct {
	db *gorm.DB
}

var _ AIStore = (*ais)(nil)

func newAIs(db *gorm.DB) *ais {
	return &ais{db}
}

func (a *ais) Create(ctx context.Context, ai *model.AIM) error {
	return a.db.Create(&ai).Error
}

func (a *ais) Get(ctx context.Context, username string, postID string) (*model.AIM, error) {
	var ai model.AIM
	if err := a.db.Where("username = ? AND postID = ?", username, postID).First(&ai).Error; err != nil {
		return nil, err
	}
	return &ai, nil
}

func (a *ais) Update(ctx context.Context, ai *model.AIM) error {
	return a.db.Save(ai).Error
}

func (a *ais) List(ctx context.Context, username string, offset, limit int) (count int64, ret []*model.AIM, err error) {
	err = a.db.Where("username = ?", username).
		// set the offset and limit
		Offset(offset).
		Limit(defaultLimit(limit)).
		// find all the records
		Find(&ret).
		// count the total number of records
		Count(&count).Error
	if err != nil {
		log.C(ctx).Errorw("failed to list AIPosts", "error", err)
	}

	return
}

func (a *ais) Delete(ctx context.Context, username, postID string) error {
	return a.db.Where("username = ? AND postID in (?)", username, postID).Delete(&model.AIM{}).Error
}
