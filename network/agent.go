package network

import (
    "net"
    "task/handler"
    "fmt"
    "protocol"
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

    var ch chan []byte
    ch = make(chan []byte)
    go handleTask(ch, *conn)
    agentHandler(ch, *conn)
}

func handleTask(ch chan []byte, conn net.TCPConn) {
    for {
        task := <-ch
        var taskHandler handler.WebHandler
        taskHandler.Task = string(task)
        taskResponse := taskHandler.DoTask()
        
        backMsg := append(append(task, []byte("####")...), []byte(taskResponse)...)
        pkgMsg := protocol.Pack(backMsg)

        writeLen, err := conn.Write(pkgMsg)
        fmt.Println(writeLen, err)
    }
}

func agentHandler(ch chan []byte, conn net.TCPConn) {
    buf := make([]byte, 1024)
    for {
        len, err := conn.Read(buf)
        if err != nil {
            fmt.Println(err)
            break //server close , err -> EOF
        }

        protocol.UnPack(buf[0:len], ch)
    }
}
