package main

import (
   "fmt"
   "log"
   "net/http"
   "encoding/json"
   "os/exec"
)

type Event struct {
   Name string `json:"name"`
   Summary string `json:"summary"`
}

func getHandlerFunc() func(w http.ResponseWriter, r *http.Request) {
   handlerFunc := func(w http.ResponseWriter, r *http.Request) {
      switch r.Method {
         case "GET":
            w.Write([]byte("get\n"))
         case "POST":
            w.Write([]byte("post\n"))
         case "PUT":
            w.Write([]byte("put\n"))
         case "DELETE":
            w.Write([]byte("delete\n"))
         default:
            w.Write([]byte("default\n"))
      }
   }
   return handlerFunc
}

func main() {
   var _, _ = json.Marshal(true)
   http.HandleFunc("/", getHandlerFunc())
   fmt.Println("Listening...")
   err := exec.Command("cd", "..").Run()
   out, err2 := exec.Command("ls", ".").Output()
   if err != nil || err2 != nil {
      fmt.Println("ERROR")
      fmt.Printf("%s\n", err)
   }
   fmt.Printf("%s\n", out)
   log.Fatal(http.ListenAndServe(":8080", nil))
}
