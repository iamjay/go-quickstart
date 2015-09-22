// +build !appengine

package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/mux"

	"github.com/iamjay/go-quickstart/app"
	"github.com/iamjay/go-quickstart/server"
)

func main() {
	r := mux.NewRouter()
	app.SetupRoutes(r)

	s := server.NewServer()
	s.Addr = ":8000"
	s.Handler = r

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	if err := s.Run(); err != nil {
		log.Panic(err)
	}

	quit := false
	for !quit {
		select {
		case <-s.Exited:
			log.Printf("Server exited\n")
			quit = true
			break
		case <-c:
			s.Stop()
			log.Printf("Signal received\n")
		}
	}
}
