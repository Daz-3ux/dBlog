// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// MySQLOptions defines the options for connecting to MySQL
type MySQLOptions struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
}

func (o *MySQLOptions) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		o.Username,
		o.Password,
		o.Host,
		o.Database,
		true,
		"Local")
}

func NewMySql(opts *MySQLOptions) (*gorm.DB, error) {
	logLevel := logger.Silent
	if opts.LogLevel != 0 {
		logLevel = logger.LogLevel(opts.LogLevel)
	}
	db, err := gorm.Open(mysql.Open(opts.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}

	// sqlDB is the low-level SQL connection object
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	return db, nil
}
