// +build appengine

package quickstart

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/iamjay/go-quickstart/app"
)

var router = mux.NewRouter()

func init() {
	http.Handle("/", router)
	app.SetupRoutes(router)
}
