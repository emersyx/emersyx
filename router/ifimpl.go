package router

import (
	"emersyx.net/emersyx_apis/emcomapi"
	"errors"
	"fmt"
)

// SetOptions sets the options received as argument.
func (r *router) SetOptions(options ...func(emcomapi.Router) error) error {
	// apply the configuration options received as arguments
	for _, option := range options {
		err := option(r)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetGateway iterates over all loaded gateways and searches for the one with the specified identifier. An error is
// returned if the identifier argument is empty or if a gateway with the specified identifier is not found.
func (r *router) GetGateway(id string) (emcomapi.Identifiable, error) {
	if id == "" {
		return nil, errors.New("method cannot be called with an empty identifier argument")
	}
	for _, gw := range r.gws {
		if gw.GetIdentifier() == id {
			return gw, nil
		}
	}
	return nil, errors.New("no gateway with the requested identifier is available")
}

// Run starts receiving messages from gateways (which are also receptors) and processors. The events are forwarded to
// processors based on the configured routes. The forwardEvent method is used for this purpose.
func (r *router) Run() error {
	// lock the mutex which protects access to the router object members (e.g. isRunning)
	r.mutex.Lock()

	// mark the router as running
	r.isRunning = true

	// create a sink channel where events from all receptor gateways are sent
	sink := make(chan emcomapi.Event)

	// iterate through all gateways
	r.log.Debugln("funelling all gateways to the sink channel")
	for _, gw := range r.gws {
		// check if they are also receptors
		if rec, ok := gw.(emcomapi.Receptor); ok {
			funnelEvents(sink, rec.GetEventsOutChannel())
		}
	}

	// iterate through all processors and start routing events from them as well if they are receptors
	r.log.Debugln("funelling all processors to the sink channel")
	for _, proc := range r.procs {
		if prec, ok := proc.(emcomapi.Receptor); ok {
			funnelEvents(sink, prec.GetEventsOutChannel())
		}
	}

	// unlock the mutex just before the possibly infinite loop which forwards events
	r.mutex.Unlock()

	// start an infinite loop where events are received from the sink channel and forwarded to the processors based on
	// the configured routes
	r.log.Debugln("start forwarding events")
	for ev := range sink {
		if err := r.forwardEvent(ev); err != nil {
			return err
		}
	}

	r.log.Debugln("exiting the router.Run method")
	return nil
}

// funnelEvents starts a goroutine which receives events from a source channel and pushes them down a sink channel. The
// same sink channel is reused for all calls to this function throughout the codebase of the router. This is why the
// function name contains the word "funnel".
func funnelEvents(sink chan emcomapi.Event, source <-chan emcomapi.Event) {
	if source != nil {
		go func() {
			for ev := range source {
				sink <- ev
			}
		}()
	}
}

// forwardEvent simply forwards the event given as argument to processors based on the configured routes.
func (r *router) forwardEvent(ev emcomapi.Event) error {
	evsrc := ev.GetSourceIdentifier()
	r.log.Debugf("forwarding event from source \"%s\"", evsrc)
	dsts, ok := r.routes[evsrc]
	r.log.Debugf("forwarding to %d destinations\n", len(dsts))
	if ok {
		fwd := ""
		for _, dst := range dsts {
			for _, proc := range r.procs {
				if proc.GetIdentifier() == dst {
					proc.GetEventsInChannel() <- ev
					fwd = dst
				}
			}
		}
		if fwd != "" {
			r.log.Debugf("event forwarded to destination \"%s\"", fwd)
		} else {
			r.log.Debugf("event was not forwarded")
		}
	} else {
		return fmt.Errorf("event received with invalid source identifier \"%s\"", evsrc)
	}
	return nil
}
