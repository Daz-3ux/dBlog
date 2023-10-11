// Copyright 2023 daz-3ux(杨鹏达) <daz-3ux@proton.me>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Daz-3ux/dBlog.

package token

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/Daz-3ux/dBlog/internal/pkg/log"
)

// Config is the configuration for the token package
type Config struct {
	// key is the secret key used to sign the token
	key string
	// identityKey is the Username
	identityKey string
}

// ErrMissingHeader represent the error when the `Authorization` header is empty
var ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")

var (
	config = Config{"szO8T9zWx+AyZB1le9MaEG7MCToMVcELHZYiakv1rRE", "identityKey"}
	once   sync.Once
)

// Init sets the package-level config, which will used for token issuing and parsing in this package
func Init(key string, identityKey string) {
	once.Do(func() {
		if key != "" {
			config.key = key
		}
		if identityKey != "" {
			config.identityKey = identityKey
		}
	})
}

// Parse uses the specified key to parse the token. If successful, it returns the token context; otherwise, it returns an error
// tokenString is the token to be parsed
// key is the secret key used to sign the token
func Parse(tokenString string, key string) (string, error) {
	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // Check the signing method
			return nil, jwt.ErrSignatureInvalid
		}

		// return the key used to sign the token
		return []byte(key), nil
	})

	// parse failed
	if err != nil {
		return "", err
	}

	var identityKey string
	// if parse success, get the identityKey from the token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		identityKey, ok = claims[config.identityKey].(string)
		if !ok {
			log.Errorw("Invalid identity key in token claims.")
		}
	}

	return identityKey, nil
}

// ParseRequest is used to parse the token from the Authorization header and passes it to the Parse function for Token parsing
func ParseRequest(c *gin.Context) (string, error) {
	header := c.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return "", ErrMissingHeader
	}

	var t string
	_, err := fmt.Sscanf(header, "Bearer %s", &t)
	if err != nil {
		return "", err
	}

	return Parse(t, config.key)
}

// Sign issues a token using jwtSecret
// the token's claims will contain the provided subject
func Sign(identityKey string) (string, error) {
	// create a token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identityKey,                               // JWT token identification
		"iss":              "daz",                                     // token issuer
		"nbf":              time.Now().Unix(),                         // token not valid before
		"iat":              time.Now().Unix(),                         // token issued at
		"exp":              time.Now().Add(time.Hour * 24 * 7).Unix(), // token expiration time
	})

	// sign the token
	// tokenString = Header.Payload.Signature
	// jwt.SigningMethodHS256 is the Header
	// jwt.MapClaims is the Payload
	// []byte(config.key) is the Signature
	tokenString, err := token.SignedString([]byte(config.key))

	return tokenString, err
}
