package main

import (
    "fmt"
    "flag"
    "helper"
    "network"
    "task/generator"
    "net"
)

func main() {
    
    defer panicRecover()

    mode := flag.String("mode", "no-mode-gave", "run-mode")
    flag.Parse()
   
    if *mode == "server" {
        
        var taskch chan string 
        taskch = make(chan string) //note for this
        var fileGenerator generator.FileGenerator
        fileGenerator = (fileGenerator.GeneratorInit("./taskfile")).(generator.FileGenerator)
        go fileGenerator.GetAllTask(taskch)
        go handleTask(taskch)

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

func handleTask(taskch chan string) {
   var doneTask map[string]string
   doneTask = make(map[string]string)
   for {
       if len(network.ClientList) == 0 {
           continue
       }

       task := <- taskch
       fmt.Println(task)
       if doneTask[task] != ""{
           continue
       }

       handlerAgent := getHandleAgent(network.ClientList)
       if handlerAgent == "" {
           continue
       }
       fmt.Println(task, handlerAgent)
       network.ClientCh[handlerAgent] <- task
       doneTask[task] = task
   }
}

func getHandleAgent(clientList map[string]net.Conn) string{
    var flag string
    for clientFlag, _ := range(clientList) {
        flag = clientFlag
    }
    return flag
}
