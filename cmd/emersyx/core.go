package main

import (
	"emersyx.net/emersyx/pkg/api"
	"emersyx.net/emersyx/pkg/log"
	"errors"
	"fmt"
)

// emersyxCore is the type implementing the emersyx core component. It loads all peripherals, and provides services to
// them (e.g. finding other peripherals by ID). This type implements the api.Core interface.
type emersyxCore struct {
	peripherals map[string]api.Peripheral
	log         *log.EmersyxLogger
}

// newCore generates a new *emersyxCore instance.
func newCore(config *emersyxConfig) (*emersyxCore, error) {
	var err error
	core := new(emersyxCore)
	core.peripherals = make(map[string]api.Peripheral)

	core.log, err = log.NewEmersyxLogger(config.LogWriter, "emersyx", config.LogLevel)
	if err != nil {
		// do not use the logger here since it might have not been initialized
		fmt.Println(err.Error())
		return nil, errors.New("could not initialize the logger for the emersyx core")
	}

	// load the peripherals from the configuration file
	err = core.loadPeripherals(config)
	if err != nil {
		core.log.Errorln(err.Error())
		core.log.Errorln("could not load all peripherals")
		return nil, err
	}

	return core, nil
}

// initPeripherals creates and initializez api.Peripheral objects for all peripherals specified in the emersyx
// configuration file.
// TODO make this function multi-threaded and load all peripherals at the same time instead of sequentially.
func (core *emersyxCore) loadPeripherals(config *emersyxConfig) error {
	for _, pcfg := range config.Peripherals {
		core.log.Debugf("creating peripheral %s\n", pcfg.Identifier)
		prl, err := api.NewPeripheral(
			api.PeripheralOptions{
				Identifier: pcfg.Identifier,
				Core:       core,
				LogWriter:  config.LogWriter,
				LogLevel:   config.LogLevel,
				ConfigPath: pcfg.ConfigPath,
			},
			pcfg.PluginPath,
		)
		if err != nil {
			core.log.Errorf("could occured while calling \"NewPeripheral\" from plugin file \"%s\"\n", pcfg)
			return err
		}
		core.peripherals[pcfg.Identifier] = prl
	}

	// after loading all peripherals, send them the core event
	core.sendEvent(api.CoreUpdate, api.PeripheralsLoaded)

	return nil
}

// sendEvent sends a new event to all peripherals with the specified type and status.
func (core *emersyxCore) sendEvent(evType string, evStatus string) {
	ev := api.NewCoreEvent(api.CoreUpdate, api.PeripheralsLoaded)
	for _, prl := range core.peripherals {
		proc, ok := prl.(api.Processor)
		if ok {
			proc.GetEventsInChannel() <- ev
		}
	}
}

// GetPeripheral searches for the api.Peripheral object with the specified identifier. The boolean return value
// specifies if the instance with the desired ID has been found or not.
func (core *emersyxCore) GetPeripheral(id string) (api.Peripheral, bool) {
	prl, ok := core.peripherals[id]
	return prl, ok
}

// ForEachPeripheral applies the function received as argument to all api.Peripheral objects loaded by the emersyx core.
func (core *emersyxCore) ForEachPeripheral(f func(api.Peripheral)) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()

	for _, prl := range core.peripherals {
		f(prl)
	}

	return
}
