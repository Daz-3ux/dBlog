package post

import (
    "github.com/Daz-3ux/dBlog/internal/dazBlog/biz"
    "github.com/Daz-3ux/dBlog/internal/dazBlog/store"
    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"
    "testing"
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
