package omega

import "code.google.com/p/gcfg"

var (
	configFileLocation = "examples/omegad.conf"
	cfg                Config
)

// Configuration Struct
type Config struct {
	// Engine configuration
	Engine engine
	// HTTP section config
	Http http
	// Hostname config
	Hosts map[string]*host
	// Log section config
	Log log
}

// Core HTTP request engine config
type engine struct {
	MaxClients int
	MaxProc    int
}

// HTTP Request/Response behavior config
type http struct {
	KeepAlive bool
	Timeout   int
}

// (V)Host config
type host struct {
	Hostname string
	Port     int
	RootDir  string
}

// Logging config
type log struct {
	LogLevel  string
	Syslog    bool
	LogFormat string
}

func parseConfigFile(filename string, cfg *Config) error {
	// Just testing config structs
	err := gcfg.ReadFileInto(cfg, configFile)
	if err != nil {
		panic(err)
	}
	yaml
}
