package network

import (
    "net"
    "fmt"
)

var ClientList []net.Conn
var Ch = make(chan string, 10)

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
        Ch <- "hello agent"
        go writeToAgent(conn)
        go recvFromAgent(conn)
    }
}
 
func writeToAgent(conn net.Conn) {
   for {
       msg := <- Ch
       fmt.Println("in serverHandler")
       len, err := conn.Write([]byte(msg)) 
       if err != nil {
           fmt.Println(err, len)
       }
       fmt.Println("send hello to agent")
   }
}

func recvFromAgent(conn net.Conn) {
    buf := make([]byte, 1024)
    for {
        len, err := conn.Read(buf)
        fmt.Println("recv data from remote server")
        if err != nil {
            fmt.Println(err)
        }

        fmt.Println(string(buf[0:len]))
        Ch <- string(buf[0:len])
    }

} 
