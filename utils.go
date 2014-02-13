package omega

import (
	"github.com/op/go-logging"
	"strconv"
)

// Return the formatted host:port string of an address
func getAddr(hostname string, port int) string {
	return hostname + ":" + strconv.Itoa(port)
}

func debug(msg string, args ...string) {}

func warn(msg string, args ...string) {}

func err(msg string, args ...string) {}
