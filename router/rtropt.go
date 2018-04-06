package router

import (
	"emersyx.net/emersyx/api"
	"errors"
	"fmt"
	"io"
)

// Logging sets the io.Writer instance to write logging messages to and the verbosity level.
func Logging(writer io.Writer, level uint) func(Router) error {
	return func(rtr Router) error {
		if writer == nil {
			return errors.New("writer argument cannot be nil")
		}
		rtr.log.SetOutput(writer)
		rtr.log.SetLevel(level)
		return nil
	}
}

// Peripherals sets the emersyx Peripheral instances for the router.
func Peripherals(prls ...api.Peripheral) func(Router) error {
	return func(rtr Router) error {
		rtr.peripherals = make([]api.Peripheral, 0)
		for _, prl := range prls {
			if prl == nil {
				return errors.New("a peripheral cannot be nil")
			}
			rtr.procs = append(rtr.peripherals, prl)
		}
		return nil
	}
}

// Routes sets the emersyx routes required to forward events between components.
func Routes(routes map[string][]string) func(Router) error {
	return func(rtr Router) error {
		rtr.routes = make(map[string][]string)
		for src, dsts := range routes {
			if len(src) == 0 {
				return errors.New("a route cannot have the empty string as a source")
			}
			if dsts == nil || len(dsts) == 0 {
				return fmt.Errorf("route with source \"%s\" has an invalid set of destinations", src)
			}
			rtr.routes[src] = make([]string, 0)
			for _, dst := range dsts {
				if len(dst) == 0 {
					return errors.New("a route cannot have the empty string in the set of destinations")
				}
				rtr.routes[src] = append(rtr.routes[src], dst)
			}
		}
		return nil
	}
}
