// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package post

//go:generate mockgen -destination mock_post.go -package post github.com/Daz-3ux/dBlog/internal/dazBlog/biz/post PostBiz

import (
	"context"
	"errors"
	"regexp"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
	"github.com/Daz-3ux/dBlog/internal/pkg/errno"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	"github.com/Daz-3ux/dBlog/internal/pkg/model"
	v1 "github.com/Daz-3ux/dBlog/pkg/api/dazBlog/v1"
)

// PostBiz defines the methods implemented in the post-module at the biz layer
type PostBiz interface {
	Create(ctx context.Context, username string, r *v1.CreatePostRequest) (*v1.CreatePostResponse, error)
	Update(ctx context.Context, username, postID string, r *v1.UpdatePostRequest) error
	Delete(ctx context.Context, username, postID string) error
	DeleteCollection(ctx context.Context, username string, postIDs []string) error
	Get(ctx context.Context, username, postID string) (*v1.GetPostResponse, error)
	List(ctx context.Context, username string, offset, limit int) (*v1.ListPostsResponse, error)
}

// postBiz implements the PostBiz interface
type postBiz struct {
	ds store.IStore
}

// ensure that postBiz implements the PostBiz interface
var _ PostBiz = (*postBiz)(nil)

func NewPostBiz(ds store.IStore) PostBiz {
	return &postBiz{ds: ds}
}

// Create is the implementation of the Create method defined in the PostBiz interface
func (p *postBiz) Create(ctx context.Context, username string, r *v1.CreatePostRequest) (*v1.CreatePostResponse,
	error) {
	var postM model.PostM
	_ = copier.Copy(&postM, r)
	postM.Username = username

	if err := p.ds.Posts().Create(ctx, &postM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'title'", err.Error()); match {
			return nil, errno.ErrTitleAlreadyExist
		}
		return nil, err
	}

	return &v1.CreatePostResponse{PostID: postM.PostID}, nil
}

// Update is the implementation of the Update method defined in the PostBiz interface
func (p *postBiz) Update(ctx context.Context, username, postID string, r *v1.UpdatePostRequest) error {
	postM, err := p.ds.Posts().Get(ctx, username, postID)
	if err != nil {
		return err
	}

	if r.Title != nil {
		postM.Title = *r.Title
	}
	if r.Content != nil {
		postM.Content = *r.Content
	}

	if err := p.ds.Posts().Update(ctx, postM); err != nil {
		return err
	}

	return nil
}

// Delete is the implementation of the Delete method defined in the PostBiz interface
func (p *postBiz) Delete(ctx context.Context, username, postID string) error {
	if err := p.ds.Posts().Delete(ctx, username, []string{postID}); err != nil {
		return err
	}

	return nil
}

// DeleteCollection is the implementation of the DeleteCollection method defined in the PostBiz interface
func (p *postBiz) DeleteCollection(ctx context.Context, username string, postIDs []string) error {
	if err := p.ds.Posts().Delete(ctx, username, postIDs); err != nil {
		return err
	}

	return nil
}

// Get is the implementation of the Get method defined in the PostBiz interface
func (p *postBiz) Get(ctx context.Context, username, postID string) (*v1.GetPostResponse, error) {
	postM, err := p.ds.Posts().Get(ctx, username, postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrPostNotFound
		}

		return nil, err
	}

	var resp v1.GetPostResponse
	_ = copier.Copy(&resp, postM)

	resp.CreatedAt = postM.CreatedAt.Format("2006-01-02 15:04:05")
	resp.UpdatedAt = postM.UpdatedAt.Format("2006-01-02 15:04:05")

	return &resp, nil
}

// List is the implementation of the List method defined in the PostBiz interface
func (p *postBiz) List(ctx context.Context, username string, offset, limit int) (*v1.ListPostsResponse, error) {
	count, list, err := p.ds.Posts().List(ctx, username, offset, limit)
	if err != nil {
		log.C(ctx).Errorw("Failed to list posts form storage", "err", err)
		return nil, err
	}

	posts := make([]*v1.PostInfo, 0, len(list))
	for _, post := range list {
		posts = append(posts, &v1.PostInfo{
			Username:  post.Username,
			PostID:    post.PostID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.ListPostsResponse{TotalCount: count, Posts: posts}, nil
}
