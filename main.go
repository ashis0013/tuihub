package main

import (
    "github.com/rivo/tview"
)

func main() {
    app := tview.NewApplication()
    ui := getTUI(app)

    if err := app.SetRoot(ui, true).SetFocus(ui).Run(); err != nil {
		panic(err)
	}
}
