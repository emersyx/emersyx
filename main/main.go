package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// initLogging configures the io.Writer instance to be used by the emersyx logger.
func initLogging(config *emersyxConfig) error {
	var sinks []io.Writer

	if config.LogStdout == true {
		sinks = append(sinks, os.Stdout)
	}

	if len(config.LogFile) > 0 {
		f, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		sinks = append(sinks, f)
	}

	config.LogWriter = io.MultiWriter(sinks...)
	return nil
}

// fail prints the error and message received as arguments. Afterwards, it calls os.Exit with the received exit code
// argument. This function is called if early initialization of emersyx fails.
func fail(err error, msg string, code int) {
	fmt.Println(err)
	fmt.Println("could not load the configuration file")
	os.Exit(code)
}

func main() {
	// parse command line arguments
	var confFile *string
	confFile = flag.String("conffile", "", "file to read configuration parameters from")
	flag.Parse()

	// load the toml configuration file
	config, err := loadConfig(confFile)
	if err != nil {
		fail(err, "could not load the configuration file", 1)
	}

	// initialize the logger
	err = initLogging(config)
	if err != nil {
		fail(err, "could not initialize logging", 2)
	}

	// create the core
	core, err := newCore(config)
	if err != nil {
		fail(err, "could not initialize the core", 3)
	}

	// create the router
	rtr, err := newRouter(config, core)
	if err != nil {
		fail(err, "could not initialize the router", 4)
	}

	rtr.run()
}
