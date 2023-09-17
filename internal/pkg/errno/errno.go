// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package errno

import (
	"errors"
	"fmt"
)

// Errno define dBlog's error code
type Errno struct {
	HTTP    int
	Code    string
	Message string
}

// Error implements the 'Error' method of the error interface
func (err *Errno) Error() string {
	return err.Message
}

// SetMessage sets the message of the Errno
func (err *Errno) SetMessage(format string, args ...interface{}) *Errno {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

// Decode attempts to extract the business error code and error message from error
func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}

	var typed *Errno
	switch {
	case errors.As(err, &typed):
		// if err is successfully converted to *Errno, return the error code and error message
		return typed.HTTP, typed.Code, typed.Message
	default:
	}

	// default returns an unknown error code and error message, which represents a server-side error
	return InternalServerError.HTTP, InternalServerError.Code, InternalServerError.Error()
}
