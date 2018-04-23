package api

import (
	"errors"
	"fmt"
	"io"
	"plugin"
)

// Peripheral is a low-level interface (w.r.t. hierarchy of types in the emersyx framework) for all components which
// have to be uniquely identifiable. This includes gateways and processors, regardless of their implementation of the
// Receptor interface.
type Peripheral interface {
	// GetIdentifier must return the identifier of the peripheral.
	GetIdentifier() string
}

// Processor is the interface for all event processors part of the emersyx platform. Each processor component must
// expose a channel via which events are received for processing. An emersyx component may implement the Processor
// interface (next to the Peripheral interface) if the component is meant to process events received via Receptors.
type Processor interface {
	// GetEventsInChannel must return the channel via which the Processor implementation receives Event objects. The
	// channel is write-only and can not be read from.
	GetEventsInChannel() chan<- Event
}

// Receptor is the interface for all event receptors part of the emersyx platform. Each receptor component must expose a
// channel via which events are pushed. An emersyx component may implement the Receptor interface (next to the
// Peripheral interface) if the component can capture events.
type Receptor interface {
	// GetEventsOutChannel must return the channel via which the Receptor implementation pushes Event objects. The
	// channel is read-only and can not be written to.
	GetEventsOutChannel() <-chan Event
}

// Event is the interface for all events supported by the various emersyx components. The emersyx event router uses this
// type to support multiple event types.
type Event interface {
	// GetSourceIdentifier must return the identifier of the emersyx peripheral which generated the event.
	GetSourceIdentifier() string
}

// PeripheralOptions specifies the options common to all Peripherals. Instances of this type are to be used when
// creating new peripherals, using the NewPeripheral function.
type PeripheralOptions struct {
	// Identifier is the unique ID string for the Peripheral instance.
	Identifier string
	// Core is the emersyx core instance which provides services to the Peripheral instance.
	Core Core
	// LogWriter is the io.Writer instance where logging messages are written to.
	LogWriter io.Writer
	// LogLevel is the verbosity level for logging messages.
	LogLevel uint
	// ConfigPath is the path to the configuration file from which the peripheral instance loads additional options.
	ConfigPath string
}

// Validate performs basic validation of the members of the PeripheralOptions instance. An error is returned if either
// member is found to have invalid values.
func (opts PeripheralOptions) Validate() error {
	if len(opts.Identifier) == 0 {
		return errors.New("identifier value cannot have zero length")
	}
	if opts.Core == nil {
		return errors.New("core cannot be nil")
	}
	if opts.LogWriter == nil {
		return errors.New("writer cannot be nil")
	}
	return nil
}

// NewPeripheral is a utility wrapper function. It opens a go plugin from the specified path and looks up the function
// with the same name. On success, it calls the exported function with the the options given as argument, and returns
// the same values as returned by the exported function. On failure, it returns an error.
func NewPeripheral(opts PeripheralOptions, path string) (Peripheral, error) {

	plug, err := plugin.Open(path)
	if err != nil {
		err := fmt.Errorf("could not open plugin file %s", path)
		return nil, err
	}

	f, err := plug.Lookup("NewPeripheral")
	if err != nil {
		err := errors.New("plugin does not export NewPeripheral")
		return nil, err
	}

	fc, ok := f.(func(opts PeripheralOptions) (Peripheral, error))
	if ok == false {
		err := errors.New("function NewPeripheral has incorrect signature")
		return nil, err
	}

	return fc(opts)
}
