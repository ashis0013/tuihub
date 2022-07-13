package main

import (
	"github.com/rivo/tview"
)

type Todo struct {
    list *tview.List
    input *tview.InputField
    ui *tview.Flex
    app *tview.Application
}

func (todo *Todo) init(app *tview.Application) {
    todo.list = tview.NewList().AddItem("bruh", "", 'o', nil).AddItem("nuh", "", 'o', nil)
    todo.input = tview.NewInputField()

    todo.ui = tview.NewFlex().SetDirection(tview.FlexRow).AddItem(
        todo.list, 0, 1, true,
    )
    todo.app = app
}

func toggleInput(app *tview.Application, redundant bool, focusable tview.Primitive, update func()) {
    if redundant {
        return
    }
    update()
    app.SetFocus(focusable)
    // app.Draw()
}

func (todo *Todo) openInput() {
    toggleInput(todo.app, todo.ui.GetItemCount() > 1, todo.input, func() {
        todo.ui.AddItem(todo.input, 1, 1, false)
    })
}

func (todo *Todo) closeInput() {
    toggleInput(todo.app, todo.ui.GetItemCount() == 1, todo.list, func() {
        todo.ui.RemoveItem(todo.input)
    })
}

func getTodo(app *tview.Application) *Todo {
    var todo Todo
    todo.init(app)

    return &todo
}
