package network

import (
    "net"
    "fmt"
    "conf"
    "strconv"
)

var ClientList  = make(map[string]net.Conn)
var ClientCh  = make(map[string]chan string)

func InitServer() {
    
    var serverConfig conf.ServerConf
    serverConfig = conf.LoadServerConfig("../conf/server.json") 

    service := serverConfig.Host + ":" + strconv.Itoa(serverConfig.Port)

    fmt.Println(service)

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
        
        /* 记录客户端标记 */
        clientFlag := conn.RemoteAddr().String()
        ClientList[clientFlag] = conn
        ClientCh[clientFlag] = make(chan string, 10)

        ClientCh[clientFlag] <- "hello agent1"
        ClientCh[clientFlag] <- "hello agent2"
        go writeToAgent(conn)
        go recvFromAgent(conn)
    }
}
 
func writeToAgent(conn net.Conn) {
   for {
       msg := <- ClientCh[conn.RemoteAddr().String()]
       fmt.Println("in serverHandler")
       len, err := conn.Write([]byte(msg)) 
       if err != nil {
           fmt.Println(err, len)
       }
       fmt.Println("send hello to agent")
   }
}

func recvFromAgent(conn net.Conn) {
    
    defer conn.Close()

    buf := make([]byte, 1024)
    for {
        len, err := conn.Read(buf)
        fmt.Println("recv data from remote agent")
        if err != nil {
            fmt.Println("this is error", err)
            break
        }

        if len > 0 {
            buf[len] = 0
        }

        fmt.Println(string(buf[0:len]), len)
        ClientCh[conn.RemoteAddr().String()] <- string(buf[0:len])
    }
} 
