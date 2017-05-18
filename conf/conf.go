package conf

import (
    "os"
    "io/ioutil"
    "encoding/json"
)

type ServerConf struct {
    Host string
    Port int
}

type AgentConf struct {
    Server_host string
    Server_port int
}

func LoadServerConfig(confPath string) ServerConf{
    var conf ServerConf
    fp, err := os.Open(confPath)
    if err != nil {
        panic("failed access to server-config-file")
    }
    content, _ := ioutil.ReadAll(fp)

    err = json.Unmarshal(content, &conf)
    if err != nil {
        panic(err)
    }
    return conf
}

func LoadAgentConfig(confPath string) AgentConf{
    var conf AgentConf
    return conf
}
