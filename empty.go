package main

import 	"github.com/rivo/tview"

const text = `
    You dont have a task press [yellow:]<Ctrl-N>[white:] and press [yellow:]Enter[white:] to create a new task
    If you wish to exit the app press [yellow:]<Ctrl-C>[white:]
`

func EmptyView() *tview.Flex {
    text := tview.NewTextView().
        SetDynamicColors(true).
        SetTextAlign(tview.AlignCenter).
        SetText(text)

	return tview.NewFlex().
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(text, 0, 1, true).
			AddItem(tview.NewBox(), 0, 1, false), 0, 1, true).
		AddItem(tview.NewBox(), 0, 1, false)
}

