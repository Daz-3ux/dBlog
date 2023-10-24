// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package ai

//go:generate mockgen -destination mock_ai.go -package ai github.com/Daz-3ux/dBlog/internal/dazBlog/biz/ai AIBiz

import (
	"context"
	"errors"
	"os"
	"regexp"

	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"

	"github.com/Daz-3ux/dBlog/internal/pkg/errno"

	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
	"github.com/Daz-3ux/dBlog/internal/pkg/log"
	"github.com/Daz-3ux/dBlog/internal/pkg/model"
	v1 "github.com/Daz-3ux/dBlog/pkg/api/dazBlog/v1"
)

type AIBiz interface {
	Create(ctx context.Context, r *v1.CreateAIRequest) error
	Get(ctx context.Context, username, postID string) (*v1.GetAIResponse, error)
	Update(ctx context.Context, r *v1.UpdateAIRequest) error
	List(ctx context.Context, username string, offset, limit int) (*v1.ListAIResponse, error)
	Delete(ctx context.Context, username, postID string) error
}

type theAIBiz struct {
	ds store.IStore
}

var _ AIBiz = (*theAIBiz)(nil)

func NewAIBiz(ds store.IStore) AIBiz {
	return &theAIBiz{ds: ds}
}

func createAIComment(postContent string) (string, error) {
	llm, err := openai.NewChat(openai.WithModel(os.Getenv("OPENAI_MODEL")), openai.WithToken(os.Getenv("OPENAI_API_KEY")))
	if err != nil {
		return "", err
	}

	chatMsg := []schema.ChatMessage{
		schema.HumanChatMessage{
			Content: "总结以下内容,生成一份简短的提纲以及内容摘要,不要有多余输出" + postContent,
		},
	}

	aiMsg, err := llm.Call(context.Background(), chatMsg)
	if err != nil {
		log.C(context.Background()).Errorw("failed to call AI", "error", err)
		return "", err
	}

	return aiMsg.GetContent(), nil
}

func (t *theAIBiz) Create(ctx context.Context, r *v1.CreateAIRequest) error {
	var aiM model.AIM

	postContent, err := t.ds.Posts().Get(ctx, r.Username, r.PostID)
	if err != nil {
		log.C(ctx).Errorw("failed to get post", "error", err)
		return err
	}

	r.Content, err = createAIComment(postContent.Content)
	if err != nil {
		log.C(ctx).Errorw("failed to create AI comment", "error", err)
		return err
	}

	_ = copier.Copy(&aiM, r)

	if err := t.ds.AIs().Create(ctx, &aiM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'postID'", err.Error()); match {
			return errno.ErrAICommentAlreadyExist
		}

		return err
	}

	return nil
}

func (t *theAIBiz) Get(ctx context.Context, username, postID string) (*v1.GetAIResponse, error) {
	aiM, err := t.ds.AIs().Get(ctx, username, postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrPostNotFound
		}

		return nil, err
	}

	var resp v1.GetAIResponse
	_ = copier.Copy(&resp, aiM)

	return &resp, nil
}

func (t *theAIBiz) Update(ctx context.Context, r *v1.UpdateAIRequest) error {
	aiM, err := t.ds.AIs().Get(ctx, r.Username, r.PostID)
	if err != nil {
		log.C(ctx).Errorw("failed to get AI comment", "error", err)
		return err
	}

	r.Content, err = createAIComment(aiM.Content)
	if err != nil {
		log.C(ctx).Errorw("failed to create AI comment", "error", err)
		return err
	}

	aiM.Content = r.Content

	if err := t.ds.AIs().Update(ctx, aiM); err != nil {
		return err
	}

	return nil
}

func (t *theAIBiz) List(ctx context.Context, username string, offset, limit int) (*v1.ListAIResponse, error) {
	count, list, err := t.ds.AIs().List(ctx, username, offset, limit)
	if err != nil {
		log.C(ctx).Errorw("failed to list AI comments from storage", "err", err)
		return nil, err
	}

	AIs := make([]*v1.AIInfo, 0, len(list))
	for _, ai := range list {
		AIs = append(AIs, &v1.AIInfo{
			Username: ai.Username,
			PostID:   ai.PostID,
			Content:  ai.Content,
		})
	}

	return &v1.ListAIResponse{TotalCount: count, AIs: AIs}, nil
}

func (t *theAIBiz) Delete(ctx context.Context, username, postID string) error {
	if err := t.ds.AIs().Delete(ctx, username, postID); err != nil {
		return err
	}

	return nil
}
