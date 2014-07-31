package omega

import "code.google.com/p/gcfg"

var (
	configFileLocation = "examples/omegad.conf"
)

// Configuration Struct
type Config struct {
	// Engine configuration
	Engine engineConfg
	// HTTP section config
	Http httpConf
	// Hostname config
	Hosts map[string]*hostConfig
	// Log section config
	Log logConf
}

// Core HTTP request engine config
type engineConf struct {
	MaxClients int
	MaxProc    int
}

// HTTP Request/Response behavior config
type httpConf struct {
	KeepAlive    bool
	ReadTimeout  int
	WriteTimeout int
}

// (V)Host config
type hostConfig struct {
	Hostname string
	Port     int
	RootDir  string
}

// Logging config
type logConf struct {
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
}
