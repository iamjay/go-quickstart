/*

Copyright (c) 2015, Pathompong Puengrostham
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

 * Redistributions of source code must retain the above copyright notice,
   this list of conditions and the following disclaimer.
 * Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND ANY
EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH
DAMAGE.

*/

package server

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type JwtAuth struct {
	InvalidTokenHandler func(http.ResponseWriter, *http.Request)
	TokenValidated      func(map[string]interface{}, *http.Request)
	SecretKey           string
}

func NewJwtAuth() *JwtAuth {
	return &JwtAuth{}
}

func (j *JwtAuth) HandlerFunc(h http.HandlerFunc) http.HandlerFunc {
	return j.Handler(h)
}

func (j *JwtAuth) Handler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if err := j.checkToken(w, req); err != nil {
			return
		}
		h.ServeHTTP(w, req)
	}
}

func (j *JwtAuth) GenerateToken(claims map[string]interface{}) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("HS256"))

	for k, v := range claims {
		t.Claims[k] = v
	}

	ts, err := t.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	return ts, nil
}

func (j *JwtAuth) checkToken(w http.ResponseWriter, req *http.Request) error {
	token, err := jwt.ParseFromRequest(req, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		j.InvalidTokenHandler(w, req)
		return err
	}

	if !token.Valid {
		j.InvalidTokenHandler(w, req)
		return fmt.Errorf("Invalid token")
	}

	if j.TokenValidated != nil {
		j.TokenValidated(token.Claims, req)
	}

	return nil
}
