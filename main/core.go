package main

import (
	"emersyx.net/emersyx/api"
	"errors"
	"fmt"
)

// core is the type implementing the emersyx core component. It loads all peripherals, and provides services to them
// (e.g. finding other peripherals by ID). This type implements the api.Core interface.
type core struct {
	peripherals map[string]api.Peripheral
	log         *api.EmersyxLogger
}

// newCore generates a new *core instance.
func newCore() (*core, error) {
	var err error
	c := new(core)
	c.peripherals = make(map[string]api.Peripheral)

	c.log, err = api.NewEmersyxLogger(ec.LogWriter, "core", ec.LogLevel)
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
// TODO make this function multi-threaded and load all peripherals at the same time instead of sequentially.
func (c *core) loadPeripherals() error {
	for _, pcfg := range ec.Peripherals {
		c.log.Debugf("creating peripheral %s\n", pcfg.Identifier)
		prl, err := api.NewPeripheral(
			api.PeripheralOptions{
				Identifier: pcfg.Identifier,
				Core:       c,
				LogWriter:  ec.LogWriter,
				LogLevel:   ec.LogLevel,
				ConfigPath: pcfg.ConfigPath,
			},
			pcfg.PluginPath,
		)
		if err != nil {
			c.log.Errorf("could occured while calling \"NewPeripheral\" from plugin file \"%s\"\n", pcfg)
			return err
		}
		c.peripherals[pcfg.Identifier] = prl
	}

	// after loading all peripherals, send the core update that all components have been loaded
	ce := api.NewCoreEvent(api.CoreUpdate, api.PeripheralsLoaded)
	for _, prl := range c.peripherals {
		proc, ok := prl.(api.Processor)
		if ok {
			proc.GetEventsInChannel() <- ce
		}
	}

	return nil
}

// GetPeripheral searches for the api.Peripheral object with the specified identifier. The boolean return value
// specifies if the instance with the desired ID has been found or not.
func (c *core) GetPeripheral(id string) (api.Peripheral, bool) {
	prl, ok := c.peripherals[id]
	return prl, ok
}

// ForEachPeripheral applies the function received as argument to all api.Peripheral objects loaded by the emersyx core.
func (c *core) ForEachPeripheral(f func(api.Peripheral)) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()

	for _, prl := range c.peripherals {
		f(prl)
	}

	return
}
