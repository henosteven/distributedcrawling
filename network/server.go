package network

import (
    "net"
    "fmt"
)

var ClientList []net.Conn

func InitServer() {
    service := ":6768"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    if err != nil {
        fmt.Println(err)
    }
    l, err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("wait for agent")

    /* 主逻辑accept */
    for {
        conn, err := l.Accept()
        fmt.Println("accept agent", conn)
        if err != nil {
            fmt.Println(err)
        }
        ClientList = append(ClientList, conn)
        go serverHandler(conn)
    }
}

func serverHandler(conn net.Conn) {
   fmt.Println("in serverHandler")
   len, err := conn.Write([]byte("hello client")) 
   if err != nil {
       fmt.Println(err, len)
   }
   fmt.Println("send hello to agent")
}
