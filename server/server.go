package server

import (
	"net"
	"net/http"
)

type Server struct {
	*http.Server

	Exited chan int

	listener net.Listener
}

func NewServer() *Server {
	return &Server{
		Server: &http.Server{},
		Exited: make(chan int),
	}
}

func (s *Server) Run() error {
	var err error
	if s.listener, err = net.Listen("tcp", s.Addr); err != nil {
		return err
	}

	go func() {
		s.Serve(s.listener)
		s.Exited <- 0
	}()

	return nil
}

func (s *Server) Stop() {
	s.listener.Close()
}
