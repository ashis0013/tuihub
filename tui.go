package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getTUI(app *tview.Application, status bool) *tview.Flex {
	todo := getTodo(app)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlN:
			todo.openInput()
		case tcell.KeyEnter, tcell.KeyEsc:
			todo.closeInput()
		case tcell.KeyRune:
			switch event.Rune() {
			case 'q':
				app.Stop()
			case 'j':
				todo.MoveDown()
			case 'k':
				todo.MoveUp()
			case ' ':
				if !todo.IsInputOpen() {
					todo.completeTask()
				}
			}
		}
		return event
	})

	tui := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(
		todo.ui, 0, 1, true,
	)
	if status {
		tui.AddItem(
			getStatusUI(app), 1, 1, false,
		)
	}
	return tui
}
