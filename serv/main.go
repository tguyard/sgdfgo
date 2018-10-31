package serv

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var Secret []byte

type Server struct {
	srv  *http.Server
	port int
}

func New(port int) *Server {
	return &Server{
		port: port,
		srv:  &http.Server{Addr: ":" + strconv.Itoa(port)},
	}
}

func (s *Server) Start() {
	http.HandleFunc("/rest/connect", connect)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting webserver on %d", s.port)
		if err := s.srv.ListenAndServe(); err != nil {
			log.Printf(err.Error())
		}
	}()

	<-sigs

	log.Printf("Shutdowning webserver")
	s.srv.Shutdown(context.Background())
}
