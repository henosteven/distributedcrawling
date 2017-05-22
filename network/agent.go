package network

import (
    "net"
    "task/handler"
    "fmt"
)

func InitAgent() {
    service := "127.0.0.1:6768"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    fmt.Println(tcpAddr)
    if err != nil {
        fmt.Println(err)
    }
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    fmt.Println(conn)
    if err != nil {
        fmt.Println(err)
    }
    agentHandler(*conn)
}

func agentHandler(conn net.TCPConn) {
    buf := make([]byte, 1024)
    for {
        len, err := conn.Read(buf)
        if err != nil {
            fmt.Println(err)
            break //server close , err -> EOF
        }
        task := string(buf[0:len])
        
        var taskHandler handler.WebHandler
        taskHandler.Task = task
        fmt.Println(task, ":", taskHandler.DoTask())
    }
}
