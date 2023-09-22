package store

import (
	"gorm.io/gorm"
	"sync"
)

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
	Users() UserStore
}

// datastore is the `factory` for creating store layer instance
// datastore is connected gorm in the under and implements the IStore methods in the upper
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
		S = &datastore{db}
	})

	return S
}

// Users return an instance of UserStore
func (ds *datastore) Users() UserStore {
	return newUsers(ds.db)
}
