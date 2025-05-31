package httpserver

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	_defaultAddr            = ":8080"
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultIdleTimeout     = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	Engine *gin.Engine

	notify          chan error
	listener        *http.Server
	address         string
	readTimeout     time.Duration
	writeTimeout    time.Duration
	idleTimeout     time.Duration
	shutdownTimeout time.Duration
}

func New(opts ...Option) *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	server := &Server{
		Engine:          engine,
		notify:          make(chan error, 1),
		address:         _defaultAddr,
		readTimeout:     _defaultReadTimeout,
		writeTimeout:    _defaultWriteTimeout,
		idleTimeout:     _defaultIdleTimeout,
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(server)
	}

	listener := &http.Server{
		Addr:         server.address,
		Handler:      server.Engine,
		ReadTimeout:  server.readTimeout,
		WriteTimeout: server.writeTimeout,
		IdleTimeout:  server.idleTimeout,
	}

	server.listener = listener

	return server
}

func (s *Server) Start() {
	go func() {
		err := s.listener.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server error: %v", err)
		}

		s.notify <- err
	}()
}

func (s *Server) RunBlocking() error {
	log.Printf("Starting HTTP server on %s", s.address)
	return s.listener.ListenAndServe()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	log.Println("Shutting down HTTP server...")

	if err := s.listener.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down HTTP server: %v", err)
		return err
	}

	log.Println("Shut down successfully")
	return nil
}

func (s *Server) Notify() <-chan error {
	return s.notify
}
