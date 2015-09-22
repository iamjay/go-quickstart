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
