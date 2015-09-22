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
