package main

import (
        "fmt"
        "strings"
        "encoding/json"
       )


type ReqData struct {
     Taskid string `json:"taskid,omitempty"`
     Need_detail bool `json:"need_detail,omitempty"`
     Alias []string   `json:"alias,omitempty"`
}

func main() {
      v:="2233806,13444,9226"
      fmt.Println(v) 
      s:=strings.Split(v,",")
      fmt.Println(s)
      a,_:= json.Marshal(s)
      fmt.Println(string (a[:]))
        
      data:=new(ReqData)
      data.Taskid="13333"
      data.Need_detail=true
      data.Alias=s
      b,_:= json.Marshal(data)
      fmt.Println(string (b[:]))
}  
