package main

import "fmt"

const (
    PluginName    = "sampleplugin"
    PluginVersion = 0x00010000
)

//export Hello
func Hello() {
    fmt.Println("Hello from plugin sample")
}

func main() {}