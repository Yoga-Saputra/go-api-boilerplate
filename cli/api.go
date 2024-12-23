package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Yoga-Saputra/go-boilerplate/config"
	"github.com/pterm/pterm"
)

var (
	// Main variable argument
	apiCLI bool

	routeList bool
)

var apiCLICommands = cli{
	argVar:   &apiCLI,
	argName:  "api",
	argUsage: "--api To run API CLI",
	run:      apiCLIRun,
	boolOptions: []optionBool{
		{
			optionVar:          &routeList,
			optionName:         "route:list",
			optionUsage:        "--route:list To show list of API routes",
			optiondefaultValue: false,
		},
	},
}

func apiCLIRun() {
	switch true {
	case routeList:
		printRouteList()

	default:
		pterm.Info.Print("Please give sub arguments")
		fmt.Println("")
	}
}

// Show and print table of API route list
func printRouteList() {
	var datas [][]string

	// Create pterm table
	tbl := pterm.DefaultTable.WithHasHeader().WithBoxed()
	header := []string{
		"Method",
		"Path",
		"Handler",
	}
	datas = append(datas, header)

	// Read from file
	f := config.Of.App.ResolveFilePathInWorkDir("route-list.json")
	r, err := ioutil.ReadFile(f)
	if err != nil {
		pterm.Error.Print(err.Error())
		return
	}

	// Unmarshal to json
	var routes []map[string]interface{}
	if err := json.Unmarshal(r, &routes); err != nil {
		pterm.Error.Print(err.Error())
		return
	}

	// Append route to table datas
	for _, v := range routes {
		datas = append(datas, []string{
			fmt.Sprintf("%v", v["method"]),
			fmt.Sprintf("%v", v["path"]),
			fmt.Sprintf("%v", v["name"]),
		})
	}

	tbl.Data = datas
	tbl.Render()
}
