package main

import (
    "fmt"
    "flag"
    "helper"
    "network"
)

func main() {
    
    defer panicRecover()

    mode := flag.String("mode", "no-mode-gave", "run-mode")
    flag.Parse()

    if *mode == "server" {
        server()
    } else if *mode == "agent" {
        agent()
    } else {
        panic("invalid node")
    }
}

func server() {
    network.InitServer() 
}

func agent() {
    network.InitAgent()
}

func panicRecover() {
    var help string = helper.HelpText()
    if err := recover(); err != nil {
        fmt.Println(err)
        fmt.Println(help)
    }
}
