package main

import (
	"emersyx.net/emersyx/api"
)

func main() {
	parseFlags()
	initLogging()
	loadConfig()

	rtr := newRouter()
	gws := initGateways()
	procs := initProcessors(rtr)
	routes := initRoutes()

	initRouter(rtr, gws, procs, routes)

	ce := emcomapi.NewCoreEvent(api.CoreUpdate, api.ComponentsLoaded)
	for _, gw := range gws {
		if ch := gw.GetEventsInChannel(); ch != nil {
			ch <- ce
		}
	}
	for _, proc := range procs {
		proc.GetEventsInChannel() <- ce
	}

	rtr.Run()
}
