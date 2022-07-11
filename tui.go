package main

import (
    "github.com/rivo/tview"
)

func getTUI(app *tview.Application) *tview.Flex {
    return tview.NewFlex().SetDirection(tview.FlexRow).AddItem(tview.NewBox().SetBorder(true), 0, 1, false).AddItem(getStatus(app), 1, 1, true)
}
