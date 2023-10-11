// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package post

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Daz-3ux/dBlog/internal/dazBlog/biz"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/store"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockIStore(ctrl)

	type args struct {
		ds store.IStore
	}
	tests := []struct {
		name string
		args args
		want *PostController
	}{
		{
			name: "default",
			args: args{mockStore},
			want: &PostController{b: biz.NewBiz(mockStore)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewPostController(tt.args.ds))
		})
	}
}
