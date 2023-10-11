// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FuzzDefaultLimit(f *testing.F) {
	testcases := []int{0, 1, 2}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig int) {
		limit := defaultLimit(orig)
		if orig == 0 {
			assert.Equal(t, defaultLimitNumber, limit)
		} else {
			assert.Equal(t, orig, limit)
		}
	})
}
