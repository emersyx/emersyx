package main

import (
	"fmt"
	"os"
)

func main() {
	parseFlags()
	loadConfig()
	err := initLogging()

	core, err := newCore()
	if err != nil {
		fmt.Println(err)
		fmt.Println("could not initialize the emersyx core")
		os.Exit(1)
	}

	routes := loadRoutes()
	rtr, err := newRouter(core, routes)
	if err != nil {
		fmt.Println(err)
		fmt.Println("could not initialize the router")
		os.Exit(1)
	}

	rtr.Run()
}
