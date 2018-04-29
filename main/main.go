package main

import (
	"emersyx.net/emersyx/api"
	"flag"
	"fmt"
	"io"
	"os"
)

// flConfFile holds the value of the command line flag which specifies the emersyx configuration file.
var flConfFile *string

// initLogging configures the io.Writer instance to be used by the emersyx logger.
func initLogging() error {
	var sinks []io.Writer

	if ec.LogStdout == true {
		sinks = append(sinks, os.Stdout)
	}

	if len(ec.LogFile) > 0 {
		f, err := os.OpenFile(ec.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		sinks = append(sinks, f)
	}

	ec.LogWriter = io.MultiWriter(sinks...)
	return nil
}

func main() {
	// parse command line arguments
	flConfFile = flag.String("conffile", "", "file to read configuration parameters from")
	flag.Parse()

	// load the toml configuration file
	loadConfig()

	// initialize the logger
	err := initLogging()

	core, err := newCore()
	if err != nil {
		fmt.Println(err)
		fmt.Println("could not initialize the emersyx core")
		os.Exit(1)
	}

	rtr, err := newRouter(
		api.PeripheralOptions{
			Core:      core,
			LogWriter: ec.LogWriter,
			LogLevel:  ec.LogLevel,
		},
	)
	if err != nil {
		fmt.Println(err)
		fmt.Println("could not initialize the router")
		os.Exit(1)
	}

	rtr.run()
}
