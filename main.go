// The entry point of this app
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"

	"github.com/Yoga-Saputra/go-boilerplate/app"
	"github.com/Yoga-Saputra/go-boilerplate/cli"
	"github.com/Yoga-Saputra/go-boilerplate/config"
)

func main() {
	// Execute CLI app if have argumen that given
	cli.Execute()

	// Run the app on the overseer
	overseer.Run(overseer.Config{
		Program: gracefulSystemStart,
		Address: fmt.Sprintf(":%v", config.Of.App.Port),
		Fetcher: &fetcher.File{
			Path:     config.Of.App.ProgramFile,
			Interval: time.Duration(config.Of.App.AUWI) * time.Second,
		},
		Debug: config.Of.App.Debug(),
	})
}

// Starting system gracefully
func gracefulSystemStart(s overseer.State) {
	app.Up(&app.AppArgs{
		NL: s.Listener,
	})

	// Listen OS Signal of interuption
	sigs := make(chan os.Signal, 1)
	signal.Notify(
		sigs,
		syscall.SIGTSTP,
		os.Interrupt,
		overseer.SIGTERM,
		overseer.SIGUSR1,
		overseer.SIGUSR2,
		syscall.SIGINT,
	)

	// Stop the app gracefully
	// As soon as posible after signal from OS that already registered listened
	<-sigs
	gracefulSystemStop()
}

// Stoping system gracefully
func gracefulSystemStop() {
	// Wait until current proccess finish their tasks
	app.FinishTask()

	// Shutdown the process
	app.Down()
}
