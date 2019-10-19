package main

import (
	"emersyx.net/emersyx/api"
	"errors"
	"fmt"
)

// emersyxRouter provides the functionality to route emersyx events between the different components (i.e. the core plus
// peripherals) loaded by emersyx.
type emersyxRouter struct {
	core   *emersyxCore
	routes map[string][]string
	sink   chan api.Event
}

// newRouter creates a new router instance, applies the options given as argument, checks for error conditions and if
// none are met, returns the object.
func newRouter(config *emersyxConfig, core *emersyxCore) (*emersyxRouter, error) {
	// validate the core argument
	if core == nil {
		return nil, errors.New("core option cannot be nil")
	}

	// create a new router instance
	rtr := new(emersyxRouter)

	// configure the core and routes
	rtr.core = core
	rtr.loadRoutes(config)

	// create a sink channel where events from all peripherals are sent
	rtr.sink = make(chan api.Event, 10)

	return rtr, nil
}

// loadRoutes formats the route information from the emersyxConfig member of the core.
func (rtr *emersyxRouter) loadRoutes(config *emersyxConfig) {
	rtr.routes = make(map[string][]string)

	for _, r := range config.Routes {
		val, ok := rtr.routes[r.Source]
		if ok {
			val := append(val, r.Destinations...)
			rtr.routes[r.Source] = val
		} else {
			narr := make([]string, len(r.Destinations))
			copy(narr, r.Destinations)
			rtr.routes[r.Source] = narr
		}
	}
}

// Run starts receiving events from peripherals. The events are forwarded to other peripherals based on the configured
// routes. The forwardEvent method is used for this purpose.
func (rtr *emersyxRouter) run() error {
	rtr.core.log.Debugln("funelling all gateways to the sink channel")
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
	rtr.core.log.Debugln("start forwarding events")
	for ev := range rtr.sink {
		if err := rtr.forwardEvent(ev); err != nil {
			return err
		}
	}

	rtr.core.log.Debugln("exiting the emersyxRouter.Run method")
	return nil
}

// funnelEvents starts a goroutine which receives events from a source channel and pushes them down the router's sink
// channel.
func (rtr *emersyxRouter) funnelEvents(source <-chan api.Event) {
	if source != nil {
		go func() {
			for ev := range source {
				rtr.sink <- ev
			}
		}()
	}
}

// forwardEvent forwards the event given as argument to peripherals based on the configured routes.
func (rtr *emersyxRouter) forwardEvent(ev api.Event) error {
	evsrc := ev.GetSourceIdentifier()
	rtr.core.log.Debugf("forwarding event from source \"%s\"", evsrc)

	if dsts, ok := rtr.routes[evsrc]; ok {
		rtr.core.log.Debugf("forwarding to %d destinations\n", len(dsts))
		for _, dst := range dsts {
			fn := func(prl api.Peripheral) {
				if prl.GetIdentifier() == dst {
					proc, ok := prl.(api.Processor)
					if ok == false {
						return
					}
					proc.GetEventsInChannel() <- ev
					rtr.core.log.Debugf("event forwarded to destination \"%s\"", prl.GetIdentifier())
				}
			}
			rtr.core.ForEachPeripheral(fn)
		}
	} else {
		return fmt.Errorf("event received with invalid source identifier \"%s\"", evsrc)
	}

	return nil
}
