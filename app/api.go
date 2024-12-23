package app

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Yoga-Saputra/go-boilerplate/api"
)

// API singleton of API instance
var API *api.Instance

// Start REST API server
func apiUp(args *AppArgs) {
	if HardMaintenance == "false" {
		// Create new instance
		api := api.New()
		API = api

		// Print info
		printOutUp(fmt.Sprintf("API services running on: %v  PID: %v", args.NL.Addr().String(), os.Getpid()))
		printOutUp(fmt.Sprintf("⇨ Echo http server (v%v) started\n", API.Version))

		// Start REST API Server using goroutine
		go func() {
			if err := api.Start(args.NL, args.Address); err != nil && err != http.ErrServerClosed {
				if !strings.Contains(err.Error(), "use of closed network connection") {
					panic(err)
				}
			}
		}()
	}
}

// Saving registered route list
func apiSaveRouteList(args *AppArgs) {
	if HardMaintenance == "false" {
		if err := API.SaveRouteList(); err != nil {
			panic(err)
		}
	}
}

// Stop REST API server
func apiDown() {
	if API != nil {
		printOutDown("Shutting down API services...")

		API.Stale()
	}
}
