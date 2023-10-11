// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package store

import (
	"sync"

	"gorm.io/gorm"
)

//go:generate mockgen -destination mock_store.go -package store github.com/Daz-3ux/dBlog/internal/dazBlog/store IStore,UserStore,PostStore

/*
	singleton factory pattern
*/

var (
	// in order to avoid the instance being created repeatedly,
	// use sync.Once to ensure that the instance is created only once
	once sync.Once
	// S global variable used for easy access to the initialized instance of store layer by other packages
	S *datastore
)

// IStore defines the methods that need to be implemented by the Store layer
// such as IStore defines Users, Users defines the specific methods
type IStore interface {
	DB() *gorm.DB
	Users() UserStore
	Posts() PostStore
}

// datastore is the `factory` for creating store layer instance
// datastore is connected gorm in the under and implements the IStore methods in the upper
/*
* use datastore instead of Database:
* hide inner struct, expose only the interface
 */
type datastore struct {
	// The core of the store layer revolves around the *gorm.DB object
	db *gorm.DB
}

// ensure that datastore implements the IStore interface
var _ IStore = (*datastore)(nil)

// NewStore create an instance of type IStore
func NewStore(db *gorm.DB) *datastore {
	// ensure that the instance is created only once
	once.Do(func() {
		S = &datastore{db: db}
	})

	return S
}

// DB return the *gorm.DB object in the datastore
func (ds *datastore) DB() *gorm.DB {
	return ds.db
}

// Users return an instance of UserStore
func (ds *datastore) Users() UserStore {
	return newUsers(ds.db)
}

// Posts return an instance of PostStore
func (ds *datastore) Posts() PostStore {
	return newPosts(ds.db)
}
