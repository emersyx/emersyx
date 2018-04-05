package main

import (
	"emersyx.net/emersyx_apis/emcomapi"
	"errors"
	"fmt"
	"io"
)

// routerOptions implements the emcomapi.RouterOptions interface. Each method returns a function, which applies a
// specific configuration to an IRCGateway object.
type routerOptions struct {
}

// Logging sets the io.Writer instance to write logging messages to and the verbosity level.
func (o routerOptions) Logging(writer io.Writer, level uint) func(emcomapi.Router) error {
	return func(rtr emcomapi.Router) error {
		if writer == nil {
			return errors.New("writer argument cannot be nil")
		}
		crtr, ok := rtr.(*router)
		if ok == false {
			return errors.New("unsupported Router implementation")
		}

		crtr.log.SetOutput(writer)
		crtr.log.SetLevel(level)
		return nil
	}
}

// Gateways sets the emersyx gateway instances for the router.
func (o routerOptions) Gateways(gws ...emcomapi.Gateway) func(emcomapi.Router) error {
	return func(rtr emcomapi.Router) error {
		crtr, ok := rtr.(*router)
		if ok == false {
			return errors.New("unsupported Router implementation")
		}

		crtr.mutex.Lock()
		defer crtr.mutex.Unlock()

		if crtr.isRunning {
			return errors.New("cannot set the Gateways option after calling the Router.Run method")
		}

		for _, gw := range gws {
			crtr.gws = append(crtr.gws, gw)
		}
		return nil
	}
}

// Processors sets the emersyx processor instances for the router.
func (o routerOptions) Processors(procs ...emcomapi.Processor) func(emcomapi.Router) error {
	return func(rtr emcomapi.Router) error {
		crtr, ok := rtr.(*router)
		if ok == false {
			return errors.New("unsupported Router implementation")
		}

		crtr.mutex.Lock()
		defer crtr.mutex.Unlock()

		if crtr.isRunning {
			return errors.New("cannot set the Gateways option after calling the Router.Run method")
		}

		for _, proc := range procs {
			crtr.procs = append(crtr.procs, proc)
		}
		return nil
	}
}

// Routes sets the emersyx routes required to forward events between components.
func (o routerOptions) Routes(routes map[string][]string) func(emcomapi.Router) error {
	return func(rtr emcomapi.Router) error {
		crtr, ok := rtr.(*router)
		if ok == false {
			return errors.New("unsupported Router implementation")
		}

		crtr.mutex.Lock()
		defer crtr.mutex.Unlock()

		if crtr.isRunning {
			return errors.New("cannot set the Gateways option after calling the Router.Run method")
		}

		for src, dsts := range routes {
			if len(src) == 0 {
				return errors.New("provided route with empty source is not valid")
			}
			if dsts == nil || len(dsts) == 0 {
				return fmt.Errorf("route with source \"%s\" has an invalid set of destinations", src)
			}
			crtr.routes[src] = make([]string, 0)
			for _, dst := range dsts {
				if len(dst) == 0 {
					return fmt.Errorf("route with source \"%s\" has an invalid destination", src)
				}
				crtr.routes[src] = append(crtr.routes[src], dst)
			}
		}
		return nil
	}
}

// NewRouterOptions generates a new routerOptions object and returns a pointer to it.
func NewRouterOptions() emcomapi.RouterOptions {
	return new(routerOptions)
}
