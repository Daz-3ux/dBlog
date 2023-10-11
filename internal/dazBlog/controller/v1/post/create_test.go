// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package post

import (
	"bytes"
	"encoding/json"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/biz"
	"github.com/Daz-3ux/dBlog/internal/dazBlog/biz/post"
	v1 "github.com/Daz-3ux/dBlog/pkg/api/dazBlog/v1"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func TestPostController_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	want := &v1.CreatePostResponse{PostID: "post-daz1234"}

	mockPostBiz := post.NewMockPostBiz(ctrl)
	mockPostBiz.EXPECT().Create(gomock.Any(), gomock.Any(), gomock.Any()).Return(want, nil).Times(1)

	mockBiz := biz.NewMockIBiz(ctrl)
	mockBiz.EXPECT().Posts().Return(mockPostBiz).AnyTimes()

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	body := bytes.NewBufferString(`{"title":"dBlog guide","content":"content"}`)
	c.Request, _ = http.NewRequest("POST", "/v1/posts", body)
	c.Request.Header.Set("Content-Type", "application/json")

	blw := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
	}
	c.Writer = blw

	type fields struct {
		b biz.IBiz
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.CreatePostResponse
	}{
		{
			name:   "default",
			fields: fields{b: mockBiz},
			args: args{
				c: c,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := &PostController{
				b: tt.fields.b,
			}
			ctrl.Create(tt.args.c)
			var resp v1.CreatePostResponse
			err := json.Unmarshal(blw.body.Bytes(), &resp)
			assert.Nil(t, err)
			assert.Equal(t, resp.PostID, want.PostID)
		})
	}
}
