package cli

import (
	"fmt"
	"time"

	"github.com/Yoga-Saputra/go-boilerplate/app"
	"github.com/Yoga-Saputra/go-boilerplate/config"
	"github.com/pterm/pterm"
)

var runApp bool

var appCommands = cli{
	argVar:   &runApp,
	argName:  "run",
	argUsage: "--run To run the App services",
	run:      printInfo,
	cb:       printUsage,
}

const (
	// Year and copyright
	yc     = "(c) 2021-%v go-boilerplate"
	banner = `
  ______                 ______        _ _                   _                 
 / _____)               (____  \      (_) |                 | |      _         
| /  ___  ___     ___    ____)  ) ___  _| | ____  ____ ____ | | ____| |_  ____ 
| | (___)/ _ \   (___)  |  __  ( / _ \| | |/ _  )/ ___)  _ \| |/ _  |  _)/ _  )
| \____/| |_| |         | |__)  ) |_| | | ( (/ /| |   | | | | ( ( | | |_( (/ / 
 \_____/ \___/          |______/ \___/|_|_|\____)_|   | ||_/|_|\_||_|\___)____)
                                                      |_|                      
  %s %s`
)

func printInfo() {
	// App name
	// s, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(config.Of.App.Name)).Srender()
	pyc := fmt.Sprintf(yc, time.Now().Year())
	header := fmt.Sprintf(pterm.LightGreen(banner), pterm.Red(app.Version), pterm.LightGreen(pyc))
	pterm.DefaultCenter.Println(header)

	additional := config.Of.App.Desc

	// App version and last build info
	lastBuild := "N/A"
	if app.LastBuildAt != "" && app.LastBuildAt != " " {
		lastBuild = app.LastBuildAt
	}
	additional += fmt.Sprintf("\nLast Build Binary at: %v", lastBuild)

	// Print additional info
	pterm.DefaultCenter.WithCenterEachLineSeparately().Println(pterm.Cyan(additional))

	// Comand list and usage headers
	fmt.Println(" Usage: --<argument>...")
	fmt.Println("")
	fmt.Println(" Arguments:")
}

func printUsage(commands map[string]cli) {
	var lists []pterm.BulletListItem
	for _, c := range commands {
		text := fmt.Sprintf("%v  [%v]", c.argName, c.argUsage)
		lists = append(lists, pterm.BulletListItem{
			Level: 2,
			Text:  text,
		})

		for _, v := range c.boolOptions {
			lists = append(lists, pterm.BulletListItem{
				Level: 4,
				Text:  fmt.Sprintf("%v  [%v]", v.optionName, v.optionUsage),
			})
		}
		for _, v := range c.float64Options {
			lists = append(lists, pterm.BulletListItem{
				Level: 4,
				Text:  fmt.Sprintf("%v  [%v]", v.optionName, v.optionUsage),
			})
		}
		for _, v := range c.intOptions {
			lists = append(lists, pterm.BulletListItem{
				Level: 4,
				Text:  fmt.Sprintf("%v  [%v]", v.optionName, v.optionUsage),
			})
		}
		for _, v := range c.stringOptions {
			lists = append(lists, pterm.BulletListItem{
				Level: 4,
				Text:  fmt.Sprintf("%v  [%v]", v.optionName, v.optionUsage),
			})
		}
		for _, v := range c.uintOptions {
			lists = append(lists, pterm.BulletListItem{
				Level: 4,
				Text:  fmt.Sprintf("%v  [%v]", v.optionName, v.optionUsage),
			})
		}
	}

	pterm.DefaultBulletList.WithItems(lists).Render()
}
