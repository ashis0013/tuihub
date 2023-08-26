package main

import (
	"flag"

	"github.com/rivo/tview"
)

func mai() {
	status := flag.Bool("stat", false, "Adds a status bar at the bottom")
	flag.Parse()
	app := tview.NewApplication()
	ui := getTUI(app, *status)

	if err := app.SetRoot(ui, true).SetFocus(ui).Run(); err != nil {
		panic(err)
	}
}
