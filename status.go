package main

import (
	"fmt"
	"time"
    "strings"

	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func setObserver(t *tview.TextView, routine <-chan string) {
    for update := range routine {
        t.SetText(update)
    }
}

func writerRoutine(status chan<- string, delay time.Duration, update func() string) {
    for {
        status <- update()
        time.Sleep(delay)
    }
}

func timeRoutine(status chan<- string) {
    writerRoutine(status, 500 * time.Millisecond, func() string {
        curTime := time.Now()
        weekday := fmt.Sprint(curTime.Weekday())[:3]
        _, month, day := curTime.Date()
        clock := strings.Split(curTime.Format(time.Layout), " ")[1]
        return fmt.Sprintf("  [:darkmagenta]%s %d %s %s", weekday, day, month, clock)
    })
}

func bytesToGB(bytes uint64) float64 {
    return float64(bytes)/(1<<30)
}

func machineRoutine(status chan<- string) {
    writerRoutine(status, 2 * time.Second, func() string {
        cpuP, _ := cpu.Percent(time.Second, false)
        v, _ := mem.VirtualMemory()
        bgColor := "darkmagenta"
        if v.UsedPercent > 70.0 || cpuP[0] > 55.0 {
            bgColor = "red"
        }
        return fmt.Sprintf(
            "[:%s]cpu: %.2f%%  ram: %.2f%% Total: %.2f GB, Free:%.2f GB",
            bgColor,
            cpuP[0],
            v.UsedPercent,
            bytesToGB(v.Total),
            bytesToGB(v.Used),
        )
    })
}


func getStatus(app *tview.Application) *tview.Flex {
    createTV := func(align int) *tview.TextView {
        return tview.NewTextView().SetDynamicColors(true).SetWrap(false).SetTextAlign(align).SetChangedFunc(func() {
            app.Draw()
        })
    }
    timeStat := createTV(0)
    machineStat := createTV(2)
    status := tview.NewFlex().AddItem(timeStat, 0, 1, false).AddItem(machineStat, 0, 2, false)
    timeStatusChan := make(chan string)
    machineStatusChan := make(chan string)
    go timeRoutine(timeStatusChan)
    go machineRoutine(machineStatusChan)
    go setObserver(timeStat, timeStatusChan)
    go setObserver(machineStat, machineStatusChan)
    return status
}
