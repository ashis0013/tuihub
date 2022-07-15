package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getTUI(app *tview.Application) *tview.Flex {
    todo := getTodo(app)

    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch(event.Key()) {
        case tcell.KeyCtrlN:
            todo.openInput()
        case tcell.KeyEnter:
            todo.closeInput()
        case tcell.KeyCtrlSpace:
            todo.completeTask()
        }
        return event
    })

    return tview.NewFlex().SetDirection(tview.FlexRow).AddItem(
        todo.ui, 0, 1, true,
    ).AddItem(
        getStatusUI(app), 1, 1, false,
    )
}
