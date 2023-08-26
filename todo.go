package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TodoItem struct {
	task      string
	timestamp time.Time
}

type Todo struct {
	list  *tview.List
	input *tview.InputField
	ui    *tview.Flex
	app   *tview.Application
	todos []*TodoItem
}

func (todo *Todo) init(app *tview.Application) {
	todo.readTodos()
	todo.list = tview.NewList()
	todo.list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && (event.Rune() == ' ' || event.Rune() == 'o') {
			//do nothing
			return nil
		}
		return event
	})
	for _, item := range todo.todos {
		ch := ' '
		text := item.task
		if time.Now().Day() != item.timestamp.Day() || time.Now().Sub(item.timestamp).Hours() > 24.0 {
			ch = 'o'
			text = fmt.Sprintf("[:red]%s[:black]", text)
		}
		todo.list.AddItem(text, "", ch, nil)
	}
	todo.input = tview.NewInputField()

	todo.ui = tview.NewFlex().SetDirection(tview.FlexRow)
	todo.ui.AddItem(
		func() tview.Primitive {
			if len(todo.todos) == 0 {
				return EmptyView()
			}
			return todo.list
		}(),
		0, 1, true,
	)
	todo.app = app
}

func (todo *Todo) MoveUp() {
	todo.list.SetCurrentItem((todo.list.GetItemCount() + todo.list.GetCurrentItem() - 1) % todo.list.GetItemCount())
}

func (todo *Todo) MoveDown() {
	todo.list.SetCurrentItem((todo.list.GetCurrentItem() + 1) % todo.list.GetItemCount())
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
		if todo.input.GetText() == "" {
			todo.ui.RemoveItem(todo.input)
			return
		}
		isFirst := len(todo.todos) == 0
		todo.appendTodo(todo.input.GetText(), time.Now())
		todo.syncBack()
		todo.input.SetText("")
		if isFirst {
			todo.ui.Clear()
			todo.ui.AddItem(todo.list, 0, 1, true)
		}
		todo.list.AddItem(todo.todos[len(todo.todos)-1].task, "", ' ', nil)
		todo.ui.RemoveItem(todo.input)
	})
}

func (todo *Todo) completeTask() {
	index := todo.list.GetCurrentItem()
	todo.list.RemoveItem(index)
	todo.todos = append(todo.todos[:index], todo.todos[index+1:]...)
	todo.syncBack()
	if len(todo.todos) == 0 {
		todo.ui.Clear()
		todo.ui.AddItem(EmptyView(), 0, 1, true)
	}
}

func getTodo(app *tview.Application) *Todo {
	todo := new(Todo)
	todo.init(app)

	return todo
}
