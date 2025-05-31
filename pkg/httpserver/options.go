package httpserver

import "time"

type Option func(*Server)

// ReadTimeout
func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.readTimeout = timeout
	}
}

// WriteTimeout
func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.writeTimeout = timeout
	}
}

// IdleTimeout
func IdleTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.idleTimeout = timeout
	}
}

// ShutdownTimeout
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

// Address
func Address(addr string) Option {
	return func(s *Server) {
		s.address = addr
	}
}
