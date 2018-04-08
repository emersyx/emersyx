package router

import (
	"emersyx.net/emersyx/api"
	"emersyx.net/emersyx/log"
	"errors"
	"fmt"
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
func NewRouter(options ...func(*Router) error) (*Router, error) {
	var err error

	rtr := new(Router)

	// generate a logger, to be updated via options
	rtr.log, err = log.NewEmersyxLogger(nil, "emrtr", log.ELNone)
	if err != nil {
		return nil, errors.New("could not create a bare logger")
	}

	// create a sink channel where events from all peripherals are sent
	rtr.sink = make(chan api.Event, 10)

	// apply the configuration options received as arguments
	for _, option := range options {
		err := option(rtr)
		if err != nil {
			return nil, err
		}
	}

	// check if the mandatory options have been set
	if rtr.core == nil {
		return nil, errors.New("the Core option has not been set")
	}
	if rtr.routes == nil {
		return nil, errors.New("the Routes option has not been set")
	}

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
					prl.GetEventsInChannel() <- ev
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
