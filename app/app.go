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

package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"

	"github.com/iamjay/go-quickstart/server"
)

var auth *server.JwtAuth = server.NewJwtAuth()

func userHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, %s", context.Get(req, "user"))
}

func publicHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	// Validate credential
	user := req.FormValue("user")
	pass := req.FormValue("password")

	if pass != "password" {
		w.WriteHeader(403)
		return
	}

	token, err := auth.GenerateToken(map[string]interface{}{"user": user})
	if err != nil {
		w.WriteHeader(500)
		return
	}

	fmt.Fprintf(w, `{ "token": "%s" }`, token)
}

func forbiddenHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(403)
}

func tokenValidated(claims map[string]interface{}, req *http.Request) {
	context.Set(req, "user", claims["user"])
}

func SetupRoutes(r *mux.Router) {
	auth.InvalidTokenHandler = forbiddenHandler
	auth.TokenValidated = tokenValidated
	auth.SecretKey = "secret"

	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/user", auth.HandlerFunc(userHandler))
	r.HandleFunc("/public", publicHandler)
}
