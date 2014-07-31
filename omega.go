package omega

import (
	"github.com/op/go-logging"
)

type Server struct {
	// Collection of all the routers
	hostRouters []HostRouter // map
	// Number of current connections
	connections int
	// Number of processors to use
	numProcs int
	// All loggers
	log logger
}

// All the required loggers
type logger struct {
	router logging.Logger
	engine logging.Logger
	//modules logging.Logger
}

// Creates a new Omega Server
func NewServer() *Server {}

// Main entry point into the server
// Blocks execution while server is running
func (s *Server) Run() {}
