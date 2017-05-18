package network

import (
    "net"
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
        fmt.Println("recv data from remote server")
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println(string(buf[0:len]))
    }
}
