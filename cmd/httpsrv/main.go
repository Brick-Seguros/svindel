package httpsrv

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	port    string
	handler http.Handler
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{
		port:    port,
		handler: handler,
	}
}

func (s *Server) Start() {
	addr := fmt.Sprintf(":%s", s.port)
	log.Printf("HTTP server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, s.handler))
}
