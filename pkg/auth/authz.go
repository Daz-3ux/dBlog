// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package auth

import (
	casbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"time"
)

const (
	// casbin access control model
	aclModel = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)`
)

// Authz defines an adapter for casbin, providing authorization function
type Authz struct {
	*casbin.SyncedEnforcer
}

// NewAuthz creates an authorizer using casbin for authorization
func NewAuthz(db *gorm.DB) (*Authz, error) {
	// Init a Gorm adapterByDB and use it in a Casbin enforcer
	adapterByDB, err := adapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}

	m, _ := model.NewModelFromString(aclModel)

	// Init the enforcer
	enforcer, err := casbin.NewSyncedEnforcer(m, adapterByDB)
	if err != nil {
		return nil, err
	}

	// Load the policy from DB
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	enforcer.StartAutoLoadPolicy(5 * time.Second)

	a := &Authz{enforcer}

	return a, nil
}

// Authorize is used for authorization
func (a *Authz) Authorize(sub, obj, act string) (bool, error) {
	return a.Enforce(sub, obj, act)
}
