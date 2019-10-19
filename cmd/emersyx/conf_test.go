package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

// config is a global *emersyxConfig instance to be used by all test cases.
var config *emersyxConfig

// confFile is the path to the configuration file to be loaded during testing.
var confFile *string

func TestMain(m *testing.M) {
	// parse command line arguments
	confFile = flag.String("conffile", "", "file to read configuration parameters from")
	flag.Parse()

	os.Exit(m.Run())
}

func TestParsing(t *testing.T) {
	config, err := loadConfig(confFile)
	if err != nil {
		t.Log(fmt.Sprintf("error occured while loading the configuration file - %s", err.Error()))
		t.Fail()
	}

	if len(config.Peripherals) != 3 {
		t.Log(fmt.Sprintf("expected 3 peripherals, got %d instead", len(config.Peripherals)))
		t.Fail()
	}
	if len(config.Routes) != 2 {
		t.Log(fmt.Sprintf("expected 2 routes in the config, got %d instead", len(config.Routes)))
		t.Fail()
	}
	if t.Failed() {
		return
	}

	peripheral := config.Peripherals[0]
	if peripheral.Identifier != "emirc" {
		t.Log(fmt.Sprintf("incorrect peripheral identifier for emirc, got \"%s\"", peripheral.Identifier))
		t.Fail()
	}
	if peripheral.PluginPath != "path/to/emirc.so" {
		t.Log(fmt.Sprintf("incorrect peripheral plugin path for emirc, got \"%s\"", peripheral.PluginPath))
		t.Fail()
	}
	if peripheral.ConfigPath != "path/to/emirc.toml" {
		t.Log(fmt.Sprintf("incorrect peripheral config file path for emirc, got \"%s\"", peripheral.ConfigPath))
		t.Fail()
	}

	rt := config.Routes[0]
	if rt.Source != "emirc" {
		t.Log(fmt.Sprintf("incorrect values for the source of the first route, got \"%d\"", len(rt.Source)))
		t.Fail()
	}
	if len(rt.Destinations) != 1 {
		t.Log(fmt.Sprintf(
			"incorrect number of destinations for the emirc route, expected 1, got %d",
			len(rt.Destinations)),
		)
		t.Fail()
	}

	if t.Failed() {
		return
	}

	if rt.Destinations[0] != "emi2t" {
		t.Log("incorrect values for destinations of the example_irc_id")
		t.Fail()
	}
}
