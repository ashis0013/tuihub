package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func (todo *Todo) readTodos() {
    home, _ := os.UserHomeDir()
    filepath := fmt.Sprintf("%s/.config/tuihub/todo.csv", home)
    file, err := ioutil.ReadFile(filepath)
    if err != nil {
        e := createFile(home, filepath)
        if e != nil {
            panic(e)
        }
    }
    todo.todos = nil
    for _, line := range strings.Split(string(file), "\n") {
        details := strings.Split(line, ",")
        if len(details) < 2 {
            continue
        }
        unix, parseErr := strconv.ParseInt(details[1], 10, 64)
        ts := time.Unix(unix, 0)
        if parseErr != nil {
            ts = time.Now()
        }
        todo.appendTodo(details[0], ts)
    }
}

func (todo *Todo) syncBack() {
    home, _ := os.UserHomeDir()
    filepath := fmt.Sprintf("%s/.config/tuihub/todo.csv", home)
    content := ""
    for _, item := range todo.todos {
        content = fmt.Sprintf("%s%s,%v\n", content, item.task, item.timestamp.Unix())
    }
    err := ioutil.WriteFile(filepath, []byte(content), 0644)
    if err != nil {
        return
    }

}

func createFile(home string, filepath string) error {
    direrr := os.MkdirAll(fmt.Sprintf("%s/.config/tuihub", home), 0750)
    if direrr != nil {
        return direrr
    }
    f, createErr := os.Create(filepath)
    if createErr != nil {
        return createErr
    }
    f.Close()
    return nil
}
