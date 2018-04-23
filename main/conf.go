package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io"
)

// ec is the emersyxConfig global instance which holds all values from the config file.
var ec emersyxConfig

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

// loadConfig opens, reads and parses the toml configuration file specified as command line argument. This function must
// be called after parseFlags().
func loadConfig() {
	// read the parameters from the specified configuration file
	_, err := toml.DecodeFile(*flConfFile, &ec)
	if err != nil {
		// use fmt.Printf as the logger has not been initialized yet
		fmt.Printf(err.Error())
		fmt.Printf("error occured while loading the configuration file")
	}
}
