package main

import (
	"time"

	"github.com/rivo/tview"
)

type TodoItem struct {
    task string
    timestamp time.Time
}

type Todo struct {
    list *tview.List
    input *tview.InputField
    ui *tview.Flex
    app *tview.Application
    todos []*TodoItem
}

func (todo *Todo) init(app *tview.Application) {
    todo.readTodos()
    todo.list = tview.NewList()
    for _, item := range todo.todos {
        todo.list.AddItem(item.task, "", ' ', nil)
    }
    todo.input = tview.NewInputField()

    todo.ui = tview.NewFlex().SetDirection(tview.FlexRow).AddItem(
        todo.list, 0, 1, true,
    )
    todo.app = app
}

func (todo *Todo) appendTodo(task string, ts time.Time) {
    item := new(TodoItem)
    item.task = task
    item.timestamp = ts
    todo.todos = append(todo.todos, item)
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
        if (todo.input.GetText() != "") {
            todo.appendTodo(todo.input.GetText(), time.Now())
            todo.list.AddItem(todo.todos[len(todo.todos)-1].task, "", ' ', nil)
            todo.syncBack()
            todo.input.SetText("")
        }
        todo.ui.RemoveItem(todo.input)
    })
}

func (todo *Todo) completeTask() {
    index := todo.list.GetCurrentItem()
    todo.list.RemoveItem(index)
    todo.todos = append(todo.todos[:index], todo.todos[index+1:]...)
    todo.syncBack()
}

func getTodo(app *tview.Application) *Todo {
    todo := new(Todo)
    todo.init(app)

    return todo
}
