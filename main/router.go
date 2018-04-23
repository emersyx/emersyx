package main

import (
	"emersyx.net/emersyx/api"
	"emersyx.net/emersyx/log"
	"errors"
	"fmt"
	"io"
)

// Router provides the functionality to route emersyx events between the different components (i.e. the core plus
// peripherals) loaded by emersyx.
type Router struct {
	core   api.Core
	routes map[string][]string
	log    *log.EmersyxLogger
	sink   chan api.Event
}

// NewRouter creates a new router instance, applies the options given as argument, checks for error conditions and if
// none are met, returns the object.
func NewRouter(opts RouterOptions) (*Router, error) {
	var err error

	// validate the router options
	if err := opts.validate(); err != nil {
		return nil, err
	}

	// create a new Router instance
	rtr := new(Router)

	// create an empty map for routing information
	rtr.routes = make(map[string][]string)

	// create a logger, to be updated via options
	rtr.log, err = log.NewEmersyxLogger(nil, "router", log.ELNone)
	if err != nil {
		return nil, errors.New("could not create a bare logger")
	}

	// create a sink channel where events from all peripherals are sent
	rtr.sink = make(chan api.Event, 10)

	// apply the configuration options received as argument
	opts.apply(rtr)

	return rtr, nil
}

// Run starts receiving events from peripherals. The events are forwarded to other peripherals based on the configured
// routes. The forwardEvent method is used for this purpose.
func (rtr *Router) Run() error {
	rtr.log.Debugln("funelling all gateways to the sink channel")
	fn := func(prl api.Peripheral) {
		// check if the peripheral is also a receptor
		if rec, ok := prl.(api.Receptor); ok {
			rtr.funnelEvents(rec.GetEventsOutChannel())
		}
	}
	// iterate through all peripherals
	rtr.core.ForEachPeripheral(fn)

	// start an infinite loop where events are received from the sink channel and forwarded to peripherals based on the
	// configured routes
	rtr.log.Debugln("start forwarding events")
	for ev := range rtr.sink {
		if err := rtr.forwardEvent(ev); err != nil {
			return err
		}
	}

	rtr.log.Debugln("exiting the router.Run method")
	return nil
}

// funnelEvents starts a goroutine which receives events from a source channel and pushes them down the router's sink
// channel.
func (rtr *Router) funnelEvents(source <-chan api.Event) {
	if source != nil {
		go func() {
			for ev := range source {
				rtr.sink <- ev
			}
		}()
	}
}

// forwardEvent forwards the event given as argument to peripherals based on the configured routes.
func (rtr *Router) forwardEvent(ev api.Event) error {
	evsrc := ev.GetSourceIdentifier()
	rtr.log.Debugf("forwarding event from source \"%s\"", evsrc)

	if dsts, ok := rtr.routes[evsrc]; ok {
		rtr.log.Debugf("forwarding to %d destinations\n", len(dsts))
		for _, dst := range dsts {
			fn := func(prl api.Peripheral) {
				if prl.GetIdentifier() == dst {
					proc, ok := prl.(api.Processor)
					if ok == false {
						return
					}
					proc.GetEventsInChannel() <- ev
					rtr.log.Debugf("event forwarded to destination \"%s\"", prl.GetIdentifier())
				}
			}
			rtr.core.ForEachPeripheral(fn)
		}
	} else {
		return fmt.Errorf("event received with invalid source identifier \"%s\"", evsrc)
	}

	return nil
}

// RouterOptions specifies the options for configuring the emersyx router.
type RouterOptions struct {
	// Core is the emersyx core instance which provides services to the Peripheral instance.
	Core api.Core
	// LogWriter is the io.Writer instance where logging messages are written to.
	LogWriter io.Writer
	// LogLevel is the verbosity level for logging messages.
	LogLevel uint
	// Routes is a map containing event routing information. Keys are sources and values are arrays of destination
	// identifiers.
	Routes map[string][]string
}

// validate checks that the members of the RouterOptions instance are valid. An error is returned if either member is
// found to have invalid values.
func (opts RouterOptions) validate() error {
	if opts.Core == nil {
		return errors.New("core cannot be nil")
	}
	if opts.LogWriter == nil {
		return errors.New("writer cannot be nil")
	}
	for src, dsts := range opts.Routes {
		if len(src) == 0 {
			return errors.New("a route cannot have an empty string as a source")
		}
		if dsts == nil || len(dsts) == 0 {
			return fmt.Errorf("route with source %s has an invalid set of destinations", src)
		}
		for _, dst := range dsts {
			if len(dst) == 0 {
				return errors.New("a route cannot have an empty string in the set of destinations")
			}
		}
	}
	return nil
}

// apply sets the router options to the Router instance received as argument.
func (opts RouterOptions) apply(rtr *Router) {
	rtr.core = opts.Core
	rtr.log.SetOutput(opts.LogWriter)
	rtr.log.SetLevel(opts.LogLevel)
	for src, dsts := range opts.Routes {
		rtr.routes[src] = make([]string, len(dsts))
		copy(rtr.routes[src], dsts)
	}
}
