package api

import (
	"io"
)

// Event is the interface for all events supported by the various emersyx components. The emersyx event router uses this
// type to support multiple event types.
type Event interface {
	// GetSourceIdentifier must return the identifier of the emersyx peripheral which generated the event.
	GetSourceIdentifier() string
}

// The PeripheralOptions interface specifies the options which can be set on a Peripheral implementation. Each method
// returns a function, which can apply the appropriate configuration to the Peripheral implementation. Each Peripheral
// implementation needs to also provide a related PeripheralOptions implementation. The return values of each method of
// the PeripheralOptions implementation must be directly usable as arguments to the NewPeripheral implementation.
// Different PeripheralOptions implementations may not be compatible with the same Peripheral implementation.
type PeripheralOptions interface {
	// Logging must set the io.Writer where messages are written to and the logging verbosity level.
	Logging(writer io.Writer, level uint) func(Peripheral) error
	// Identifier must set the unique ID string for the Peripheral instance.
	Identifier(id string) func(Peripheral) error
	// ConfigPath sets the path to the configuration file from which the peripheral instance loads additional
	// configuration options.
	ConfigPath(cfg string) func(Peripheral) error
	// Core must set the emersyx core instance to be used when requiring its services.
	Core(core Core) func(Peripheral) error
}

// Peripheral is a low-level interface (w.r.t. hierarchy of types in the emersyx framework) for all components which
// have to be uniquely identifiable. This includes gateways and processors (regardless of the implementation of the
// Receptor interface).
type Peripheral interface {
	// GetIdentifier must return the identifier of the peripheral.
	GetIdentifier() string
	// GetEventsInChannel must return the channel via which the Processor implementation receives Event objects. The
	// channel is write-only and can not be read from.
	GetEventsInChannel() chan<- Event
}

// Receptor is the interface for all event receptors part of the emersyx platform. Each receptor component must expose
// a channel via which events are pushed. An emersyx component may implement the Receptor interface (next to the
// Peripheral interface) if the component can capture events.
type Receptor interface {
	// GetEventsOutChannel must return the channel via which the Receptor implementation pushes Event objects. The
	// channel is read-only and can not be written to.
	GetEventsOutChannel() <-chan Event
}
