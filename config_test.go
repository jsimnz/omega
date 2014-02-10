package omega

import (
	"code.google.com/p/gcfg"
	"fmt"
	"testing"
)

var (
	isSetup = false

	basehttpConfig = BaseHttpConfig{
		Hostname: "",
		Port:     8080,
		Root:     "/var/www",
	}
	testHttpConf = Config{
		Http: struct {
			BaseHttpConfig
			MaxClients int
			MaxProc    int
			KeepAlive  bool
			Timeout    int
		}{
			BaseHttpConfig: basehttpConfig,
			MaxClients:     100,
			MaxProc:        0,
			KeepAlive:      false,
			Timeout:        30,
		},
		Log: struct {
			LogLevel  string
			Syslog    bool
			LogFormat string
		}{
			LogLevel:  "warn",
			Syslog:    false,
			LogFormat: "{Unix-Time}: {METHOD}, {URI} - {Referer}|{IP}|{User-Agent}",
		},
	}
)

func setup() error {
	if !isSetup {
		err := gcfg.ReadFileInto(&cfg, configFile)
		if err != nil {
			isSetup = false
			return err
		}
	}
	isSetup = true
	return nil
}

func loadConfig() error {
	err := gcfg.ReadFileInto(&cfg, configFile)
	if err != nil {
		return err
	}
	return nil
}

func TestHttpConf(t *testing.T) {
	err := setup()
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}

	if cfg.Http != testHttpConf.Http {
		t.Errorf("Invalid HTTP Configuration, Expected: %v, Got: %v", testHttpConf, cfg.Http)
	}
}

func TestLogConf(t *testing.T) {
	err := setup()
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}

	if cfg.Log != testHttpConf.Log {
		t.Errorf("Invalid Log configuration, Exptected: %v, Got: %v", testHttpConf.Log, cfg.Log)
	}
}

func TestHostsConf(t *testing.T) {
	// Leave for now
}
