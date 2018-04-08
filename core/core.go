package main

import (
	"emersyx.net/emersyx/api"
	"emersyx.net/emersyx/log"
	"errors"
	"fmt"
	"plugin"
)

// core is the type implementing the emersyx core component. It loads all peripherals, and provides services to them
// (e.g. finding other peripherals by ID). This type implements the api.Core interface.
type core struct {
	peripherals map[string]api.Peripheral
	log         *log.EmersyxLogger
}

// newCore generates a new *core instance.
func newCore() (*core, error) {
	var err error
	c := new(core)
	c.peripherals = make(map[string]api.Peripheral)

	c.log, err = log.NewEmersyxLogger(ec.LogWriter, "emcore", ec.LogLevel)
	if err != nil {
		// do not use the logger here since it might have not been initialized
		fmt.Println(err.Error())
		return nil, errors.New("could not initialize the logger for the emersyx core")
	}

	// load the peripherals from the configuration file
	err = c.loadPeripherals()
	if err != nil {
		c.log.Errorln(err.Error())
		c.log.Errorln("could not load all peripherals")
		return nil, err
	}

	return c, nil
}

// initPeripherals creates and initializez api.Peripheral objects for all peripherals specified in the emersyx
// configuration file.
func (c *core) loadPeripherals() error {
	for _, pcfg := range ec.Peripherals {
		plug, err := plugin.Open(pcfg.PluginPath)
		if err != nil {
			return err
		}

		opts, err := api.NewPeripheralOptions(plug)
		if err != nil {
			return err
		}

		peripheral, err := api.NewPeripheral(plug,
			opts.Logging(ec.LogWriter, ec.LogLevel),
			opts.Identifier(pcfg.Identifier),
			opts.ConfigPath(pcfg.ConfigPath),
			opts.Core(c),
		)
		if err != nil {
			return err
		}

		c.peripherals[pcfg.Identifier] = peripheral
	}

	// after loading all peripherals, send the core update that all components have been loaded
	ce := api.NewCoreEvent(api.CoreUpdate, api.PeripheralsLoaded)
	for _, peripheral := range c.peripherals {
		if ch := peripheral.GetEventsInChannel(); ch != nil {
			ch <- ce
		}
	}

	return nil
}

// GetPeripheral searches for the api.Peripheral object with the specified identifier. The boolean return value
// specifies if the instance with the desired ID has been found or not.
func (c *core) GetPeripheral(id string) (api.Peripheral, bool) {
	peripheral, ok := c.peripherals[id]
	return peripheral, ok
}

// ForEachPeripheral applies the function received as argument to all api.Peripheral objects loaded by the emersyx core.
func (c *core) ForEachPeripheral(f func(api.Peripheral)) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()

	for _, peripheral := range c.peripherals {
		f(peripheral)
	}

	return
}
