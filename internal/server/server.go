// Package server implements server.
package server

import (
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

// Subscriber base struct.
type Subscriber func(msg string) error

// Server base struct.
type Server struct {
	router   *chi.Mux
	upgrader *websocket.Upgrader
	sync.RWMutex
	subscribers map[string]Subscriber
}

// New returns new server.
func New() *Server {
	router := chi.NewRouter()

	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	serv := &Server{
		router:      router,
		upgrader:    upgrader,
		subscribers: map[string]Subscriber{},
	}

	serv.ApplyHandlers()

	return serv
}

// Start starts server.
func (serv *Server) Start() error {
	return http.ListenAndServe(":8080", serv.router)
}
