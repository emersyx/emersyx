package main

import (
	"emersyx.net/emersyx/api"
	"errors"
	"fmt"
)

// router provides the functionality to route emersyx events between the different components (i.e. the core plus
// peripherals) loaded by emersyx.
type router struct {
	core   api.Core
	routes map[string][]string
	log    *api.EmersyxLogger
	sink   chan api.Event
}

// newRouter creates a new router instance, applies the options given as argument, checks for error conditions and if
// none are met, returns the object.
func newRouter(opts api.PeripheralOptions) (*router, error) {
	var err error

	// validate options
	if opts.Core == nil {
		return nil, errors.New("core option cannot be nil")
	}

	// create a new router instance
	rtr := new(router)

	rtr.core = opts.Core

	// create a logger, to be updated via options
	rtr.log, err = api.NewEmersyxLogger(opts.LogWriter, "router", opts.LogLevel)
	if err != nil {
		return nil, errors.New("could not create a bare logger")
	}

	// set up the routing
	rtr.routes = loadRoutes()

	// create a sink channel where events from all peripherals are sent
	rtr.sink = make(chan api.Event, 10)

	return rtr, nil
}

// loadRoutes formats the route information from the global emersyxConfig instance (initialized via loadConfig).
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

// Run starts receiving events from peripherals. The events are forwarded to other peripherals based on the configured
// routes. The forwardEvent method is used for this purpose.
func (rtr *router) run() error {
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
func (rtr *router) funnelEvents(source <-chan api.Event) {
	if source != nil {
		go func() {
			for ev := range source {
				rtr.sink <- ev
			}
		}()
	}
}

// forwardEvent forwards the event given as argument to peripherals based on the configured routes.
func (rtr *router) forwardEvent(ev api.Event) error {
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
