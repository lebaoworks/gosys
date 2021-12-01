package main

import (
    "fmt"
    "github.com/lebaoworks/shared/go/plugin"
    "github.com/lebaoworks/shared/go/fault"
    "runtime/debug"
)

func PluginRoutine(path string) (err error) {
    old := debug.SetPanicOnFault(true)
    defer debug.SetPanicOnFault(old)
    defer fault.CatchFault(&err)

    p1, e := plugin.Load(path)
    if e != nil {
        fmt.Println("load plugin ", path, " error: ", e)
        return e
    }
    fmt.Println("loaded plugin ", p1.Name(), " version: ", p1.Version())
    var pointer *int
    *pointer = 5

    return
}

func main() {
    err := PluginRoutine("sample.so")
    if err != nil {
        fmt.Println("Run got error: ", err)
    }
    fmt.Println("OK")
}
