package router

import (
	"emersyx.net/emersyx/api"
	"errors"
	"fmt"
	"io"
)

// Options is a utility struct type which combines all options for a router. Each method of this type returns a
// function, which applies a specific configuration to a Router object. This type is called Options (and not
// RouterOptions) because of a suggestion made by the go linter.
type Options struct {
}

// NewOptions generates a new *Options instance which can be used to set options for a new Router instance.
func NewOptions() *Options {
	return new(Options)
}

// Logging sets the io.Writer instance to write logging messages to and the verbosity level.
func (ropts *Options) Logging(writer io.Writer, level uint) func(*Router) error {
	return func(rtr *Router) error {
		if writer == nil {
			return errors.New("writer argument cannot be nil")
		}
		rtr.log.SetOutput(writer)
		rtr.log.SetLevel(level)
		return nil
	}
}

// Core sets the api.Core instance which provides services to the router.
func (ropts *Options) Core(core api.Core) func(*Router) error {
	return func(rtr *Router) error {
		if core == nil {
			return errors.New("core argument cannot be nil")
		}
		rtr.core = core
		return nil
	}
}

// Routes sets the emersyx routes required to forward events between components.
func (ropts *Options) Routes(routes map[string][]string) func(*Router) error {
	return func(rtr *Router) error {
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
