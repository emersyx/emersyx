package api

import (
	"io"
)

// The RouterOptions interface specifies the options which can be set on a Router implementation. Each method returns a
// function, which can apply the appropriate configuration to the Router implementation. Each Router implementation
// needs to also provide a related RouterOptions implementation. The return values of each method of the RouterOptions
// implementation must be directly usable as arguments to the Router.SetOptions implementation. Different RouterOptions
// implementations may not be compatible with the same Router implementation.
type RouterOptions interface {
	// Core must set the emersyx core instance to be used when requiring its services.
	Core(core Core) func(Router) error
	// Logging must set the io.Writer where messages are written to and the logging verbosity level.
	Logging(writer io.Writer, level uint) func(Router) error
	// Gateways must set the emersyx gateway components to route events from.
	Gateways(gws ...Peripheral) func(Router) error
	// Processors must set the emersyx processor components to route events to.
	Processors(procs ...Peripheral) func(Router) error
	// Routes must set the links through which events are sent from sources to destinations. Sources can be either
	// gateways or processors, whereas destinations can only be processors.
	Routes(routes map[string][]string) func(Router) error
}

// Router is the interface which must be implemented by all emersyx routers. The standard router implementation follows
// this interface. The emersyx core also expects implementations to follow it.
type Router interface {
	// SetOptions sets the router options given as arguments. These are options implemented via RouterOptions.
	SetOptions(options ...func(Router) error) error
	// Run must start a loop in which events are forwarded via the routes specified by the RouterOptions.Routes option.
	Run() error
}

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
	// Core must set the emersyx core instance to be used when requiring its services.
	Core(core Core) func(Peripheral) error
	// Config sets the path to the configuration file from which the peripheral instance loads additional configuration
	// options.
	Config(cfg string) func(Peripheral) error
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
