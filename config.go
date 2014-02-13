package omega

import (
	"code.google.com/p/gcfg"
)

var (
	configFile = "examples/omega.conf"
	cfg        Config
)

// Config Struct
type Config struct {
	// HTTP section config
	Http struct {
		BaseHttpConfig
		MaxClients int
		MaxProc    int
		KeepAlive  bool
		Timeout    int
	}

	// Hostname config
	Hosts map[string]*Host

	// Log section config
	Log struct {
		LogLevel  string
		Syslog    bool
		LogFormat string
	}
}

type Host struct {
	BaseHttpConfig
}

// Basic HTTP configurations for HTTP & Hostname sections
type BaseHttpConfig struct {
	Hostname string
	Port     int
	Root     string
}

// Parse command line args
func init() {
	// Just testing config structs
	err := gcfg.ReadFileInto(&cfg, configFile)
	if err != nil {
		panic(err)
	}
}
