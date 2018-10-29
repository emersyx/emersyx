package main

import (
	"github.com/BurntSushi/toml"
	"io"
)

// peripheralConfig is the struct for holding processor configuration values from the emersyx configuration file.
type peripheralConfig struct {
	Identifier string
	ConfigPath string `toml:"config_path"`
	PluginPath string `toml:"plugin_path"`
}

// routeConfig is the struct for holding route configuration values from the emersyx configuration file.
type routeConfig struct {
	Source       string
	Destinations []string
}

// emersyxConfig is the container struct for holding all configuration values from the emersyx configuration file.
type emersyxConfig struct {
	LogStdout   bool   `toml:"log_stdout"`
	LogFile     string `toml:"log_file"`
	LogLevel    uint   `toml:"log_level"`
	LogWriter   io.Writer
	Peripherals []peripheralConfig
	Routes      []routeConfig
}

// loadConfig opens, reads and parses the toml configuration file received as argument.
func loadConfig(confFile *string) (*emersyxConfig, error) {
	config := new(emersyxConfig)

	// read the parameters from the specified configuration file
	_, err := toml.DecodeFile(*confFile, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
