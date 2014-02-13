package omega

import ()

type Router struct {

	// The formatted host:port address the router is connected to
	addr string

	// Host
	host string

	// Port
	port int

	// DocumentRoot
	root string

	// Number of connections currently in use
	connections uint

	// Hosts under the same port using the same router
	hosts map[string]*Host

	// Rules parsed out of the pseudo-htaccess files
	rules map[string]interface{}
}

// Creates a new router to handle requests
func NewRouter(host string, port int, root string) *Router {
	r := &Router{
		addr:        getAddr(host, port),
		host:        host,
		port:        port,
		root:        root,
		connections: 0,
	}

}

// Satisfies the Handler interface
func (r *Router) ServeHTTP() {

}

// Parse out the vhosts from the global config that are for the given router (identified by port)
func parseVHosts(r *Router) map[string]*Host {
	m := make(map[string]*Host)
	for host := range cfg.Hosts {
		h := cfg.Hosts[host]
		if r.port == h.Port {
			m[host] = &h
		}
	}

	return m
}
