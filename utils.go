package omega

import "strconv"

// Return the formatted host:port string of an address
func getAddr(hostname string, port int) string {
	return hostname + ":" + strconv.Itoa(port)
}
