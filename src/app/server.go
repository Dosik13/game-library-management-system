package app

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Router *mux.Router
	Port   string
}

func NewServer(port string) *Server {
	router := mux.NewRouter()

	return &Server{
		Router: router,
		Port:   port,
	}
}

func (s *Server) Start() error {
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         ":" + s.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	errChan := make(chan error, 1)

	go func() {
		log.Printf("Server is starting at %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	select {
	case <-c:
		log.Println("Shutting down the server...")
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		return srv.Shutdown(ctx)
	case err := <-errChan:
		return err
	}
}
