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
