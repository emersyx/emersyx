package router

import (
	"emersyx.net/emersyx/apis"
	"emersyx.net/emersyx_log/emlog"
	"errors"
	"sync"
)

// Router provides the functionality to route emersyx events between the different components (i.e. the core plus
// peripherals) loaded by emersyx.
type Router struct {
	peripherals []api.Peripheral
	routes      map[string][]string
	log         *emlog.EmersyxLogger
	sink        chan api.Event
}

// NewRouter creates a new router instance, applies the options given as argument, checks for error conditions and if
// none are met, returns the object.
func NewRouter() (Router, error) {
	var err error

	rtr := new(router)

	// generate a logger, to be updated via options
	rtr.log, err = emlog.NewEmersyxLogger(nil, "emrtr", emlog.ELNone)
	if err != nil {
		return nil, errors.New("could not create a bare logger")
	}

	// create a sink channel where events from all peripherals are sent
	rtr.sink = make(chan api.Event, 10)

	// apply the configuration options received as arguments
	for _, option := range options {
		err := option(rtr)
		if err != nil {
			return err
		}
	}

	// check that the peripherals and routes have been set
	if rtr.peripherals == nil {
		return errors.New("the Peripherals option has not been set")
	}
	if rtr.routes == nil {
		return errors.New("the Routes option has not been set")
	}

	return rtr, nil
}

// Run starts receiving events from peripherals. The events are forwarded to other peripherals based on the configured
// routes. The forwardEvent method is used for this purpose.
func (rtr *Router) Run() error {
	// iterate through all peripherals
	r.log.Debugln("funelling all gateways to the sink channel")
	for _, prl := range r.peripherals {
		// check if they are also receptors
		if rec, ok := prl.(api.Receptor); ok {
			funnelEvents(rec.GetEventsOutChannel())
		}
	}

	// start an infinite loop where events are received from the sink channel and forwarded to peripherals based on the
	// configured routes
	r.log.Debugln("start forwarding events")
	for ev := range sink {
		if err := r.forwardEvent(ev); err != nil {
			return err
		}
	}

	r.log.Debugln("exiting the router.Run method")
	return nil
}

// funnelEvents starts a goroutine which receives events from a source channel and pushes them down the router's sink
// channel.
func (rtr *Router) funnelEvents(sink chan api.Event, source <-chan api.Event) {
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
	r.log.Debugf("forwarding event from source \"%s\"", evsrc)

	if dsts, ok := r.routes[evsrc]; ok {
		r.log.Debugf("forwarding to %d destinations\n", len(dsts))
		for _, dst := range dsts {
			for _, prl := range r.peripherals {
				if prl.GetIdentifier() == dst {
					prl.GetEventsInChannel() <- ev
					r.log.Debugf("event forwarded to destination \"%s\"", prl.GetIdentifier())
				}
			}
		}
	} else {
		return fmt.Errorf("event received with invalid source identifier \"%s\"", evsrc)
	}

	return nil
}
