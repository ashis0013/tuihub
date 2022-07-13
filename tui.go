package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func getTUI(app *tview.Application) *tview.Flex {
    todoUI := getTodo(app)

    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch(event.Key()) {
        case tcell.KeyCtrlN:
            todoUI.openInput()
        case tcell.KeyEnter:
            todoUI.closeInput()
        }
        return event
    })

    return tview.NewFlex().SetDirection(tview.FlexRow).AddItem(
        todoUI.ui, 0, 1, true,
    ).AddItem(
        getStatusUI(app), 1, 1, false,
    )
}
