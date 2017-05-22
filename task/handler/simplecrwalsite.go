package handler

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strings"
)

type WebHandler struct {
    Task string    
}

func (this WebHandler) DoTask() (string) {
   fmt.Println(this.Task) 
   resp, err := http.Get(strings.Replace(this.Task, "\n", "", -1))
   if err != nil {
       fmt.Println("http get error", err)
   }
   defer resp.Body.Close()
   
   body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
       fmt.Println("http read error")
   }

   sourcePage := string(body)
   return sourcePage
}
