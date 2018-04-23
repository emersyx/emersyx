package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	parseFlags()
	loadConfig()
	os.Exit(m.Run())
}

func TestParsing(t *testing.T) {
	if len(ec.Peripherals) != 3 {
		t.Log(fmt.Sprintf("expected 3 peripherals, got %d instead", len(ec.Peripherals)))
		t.Fail()
	}
	if len(ec.Routes) != 2 {
		t.Log(fmt.Sprintf("expected 2 routes in the config, got %d instead", len(ec.Routes)))
		t.Fail()
	}
	if t.Failed() {
		return
	}

	peripheral := ec.Peripherals[0]
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

	rt := ec.Routes[0]
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