package omega

import (
	"net/http"
)

type HostRouter struct {
	// The formatted host:port address the router is connected to
	addr string
	// Host
	host string
	// Port
	port int
	// DocumentRoot
	rootDir string
	// Number of connections currently in use
	connections uint
	// Hosts under the same port using the same router
	hosts map[string]*Host
	// Rules parsed out of the pseudo-htaccess files
	rules map[string]interface{}
	// HTTP behavior config
	httpConf httpConf
}

// Creates a new router to handle requests
func NewHostRouter(host string, port int, root string) *HostRouter {
	r := &Router{
		addr:        getAddr(host, port),
		host:        host,
		port:        port,
		rootDir:     root,
		connections: 0,
	}
	r.hosts = parseVHosts(r)
	return r
}

// Run a custom http.Server instantce
func (r *HostRouter) run(s Server) {
	httpServer := &http.Server{
		Addr:         r.addr,
		Handler:      r,
		ReadTimeout:  r.httpConf.ReadTimeout,
		WriteTimeout: r.httpConf.WriteTimeout,
	}
}

// satisfy the http.Handler interface
func (r *HostRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {

}

// Parse out the vhosts from the global config that are for the given router (identified by port)
func parseVHosts(r Router) map[string]*Host {
	m := make(map[string]*Host)
	for host := range cfg.Hosts {
		h := cfg.Hosts[host]
		if r.port == h.Port {
			m[host] = &h
		}
	}

	return m
}
