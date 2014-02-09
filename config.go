package omega

import (
	"code.google.com/p/gcfg/"
)

var (
	cfg        Config
	configFile = "/etc/omega/omega.conf"
)

// Config Struct
type Config struct {
	// HTTP section config
	Http struct {
		BaseHttpConfig
		MaxClients string
		MaxProc    int
		KeepAlive  bool
		Timeout    int
	}

	// Hostname config
	Hostname map[string]*struct {
		BaseHttpConfig
	}

	// Log section config
	Log struct {
		LogLevel  string
		Syslog    bool
		LogFormat string
	}
}

// Basic HTTP configurations for HTTP & Hostname sections
type BaseHttpConfig struct {
	Hostname string
	Port     int
	WWW      string
}

// Parse command line args
func init() {
	// Just testing config structs
	err := gcfg.ReadFileInto(&cfg, configFile)
}
