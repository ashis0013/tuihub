package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func setObserver(t *tview.TextView, routine chan string) {
    for update := range routine {
        t.SetText(update)
    }
}

func timeRoutine(status chan string) {
    for {
        status <- fmt.Sprintf("UnixDate: %s", time.Now().Format(time.UnixDate))
        time.Sleep(500 * time.Millisecond)
    }
}

func bytesToGB(bytes uint64) float64 {
    return float64(bytes)/(1<<30)
}

func machineRoutine(status chan string) {
    for {
        cpuP, _ := cpu.Percent(time.Second, false)
        v, _ := mem.VirtualMemory()
        status <- fmt.Sprintf(
            "cpu: %.2f%%  ram: %.2f%% Total: %.2f GB, Free:%.2f GB",
            cpuP[0],
            v.UsedPercent,
            bytesToGB(v.Total),
            bytesToGB(v.Used),
        )
        time.Sleep(2 * time.Second)
    }
}


func getStatus(app *tview.Application) *tview.Flex {
    createTV := func() *tview.TextView {
        return tview.NewTextView().SetWrap(false).SetChangedFunc(func() {
            app.Draw()
        })
    }
    timeStat := createTV()
    machineStat := createTV()
    status := tview.NewFlex().AddItem(timeStat, 0, 1, false).AddItem(machineStat, 0, 1, false)
    timeStatusChan := make(chan string)
    machineStatusChan := make(chan string)
    go timeRoutine(timeStatusChan)
    go machineRoutine(machineStatusChan)
    go setObserver(timeStat, timeStatusChan)
    go setObserver(machineStat, machineStatusChan)
    return status
}
