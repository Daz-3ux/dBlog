// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package id

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenShortID(t *testing.T) {
	shortID := GenShortID()
	assert.NotEqual(t, "", shortID)
	assert.Equal(t, 6, len(shortID))
}

func BenchmarkGenShortID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}

func BenchmarkGenShortIDTimeConsuming(b *testing.B) {
	b.StopTimer()

	shortid := GenShortID()
	if shortid == "" {
		b.Error("Failed to generate short id")
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}
