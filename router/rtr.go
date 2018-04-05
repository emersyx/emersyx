package router

import (
	"emersyx.net/emersyx_apis/emcomapi"
	"emersyx.net/emersyx_log/emlog"
	"errors"
	"sync"
)

// router is the struct type which implements the emcomapi.Router interface. Objects of this type are used to route
// events from receptor gateways to processors.
type router struct {
	gws       []emcomapi.Identifiable
	procs     []emcomapi.Processor
	routes    map[string][]string
	isRunning bool
	log       *emlog.EmersyxLogger
	mutex     sync.Mutex
}

// NewRouter creates a new router instance, applies the options given as argument, checks for error conditions and if
// none are met, returns the object.
func NewRouter() (emcomapi.Router, error) {
	var err error

	rtr := new(router)

	// generate a logger, to be updated via options
	rtr.log, err = emlog.NewEmersyxLogger(nil, "emrtr", emlog.ELNone)
	if err != nil {
		return nil, errors.New("could not create a bare logger")
	}

	// the router is initially not running
	// this member is set to true once the router.Run method is called
	rtr.isRunning = false

	// create member arrays with default sizes
	rtr.gws = make([]emcomapi.Identifiable, 0)
	rtr.procs = make([]emcomapi.Processor, 0)
	rtr.routes = make(map[string][]string)

	return rtr, nil
}
