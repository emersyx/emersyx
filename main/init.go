package main

import (
	"flag"
	"io"
	"os"
)

// flConfFile holds the value of the command line flag which specifies the emersyx configuration file.
var flConfFile *string

// parseFlags parses the command line arguments given to the emersyx binary.
func parseFlags() {
	// set the expected flags
	flConfFile = flag.String("conffile", "", "file to read configuration parameters from")

	// parse the flags
	flag.Parse()
}

// initLogging configures the logger (i.e. the el global variable). The parseFlags function needs to be called before
// this one.
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

// loadRoutes formats the route information from the global emersyxConfig instance (initialized via loadConfig) such
// that it can be passed as argument to the api.Options.Routes method.
func loadRoutes() map[string][]string {
	var m = make(map[string][]string)

	for _, cfg := range ec.Routes {
		val, ok := m[cfg.Source]
		if ok {
			val := append(val, cfg.Destinations...)
			m[cfg.Source] = val
		} else {
			narr := make([]string, len(cfg.Destinations))
			copy(narr, cfg.Destinations)
			m[cfg.Source] = narr
		}
	}

	return m
}

// newRouter creates an api.Router object as specified in the emersyx configuration file. Under the hood, the
// router.NewRouter function is used.
func newRouter(c *core, routes map[string][]string) (*Router, error) {
	rtr, err := NewRouter(
		RouterOptions{
			Core:      c,
			LogWriter: ec.LogWriter,
			LogLevel:  ec.LogLevel,
			Routes:    routes,
		},
	)
	if err != nil {
		return nil, err
	}
	return rtr, nil
}
