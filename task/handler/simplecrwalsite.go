package handler

import (
    "fmt"
)

type WebHandler struct {
    Task string    
}

func (this WebHandler) DoTask() (string) {
   fmt.Println(this.Task) 
   return this.Task + "done"
}
