package main

import (
   "fmt"
   "log"
   "net/http"
)

func getHandlerFunc() func(w http.ResponseWriter, r *http.Request) {
   handlerFunc := func(w http.ResponseWriter, r *http.Request) {
      switch r.Method {
         case "GET": w.Write([]byte("get\n"))
         case "POST": w.Write([]byte("post\n"))
         case "PUT": w.Write([]byte("put\n"))
         case "DELETE": w.Write([]byte("delete\n"))
         default: w.Write([]byte("default\n"))
      }
   }
   return handlerFunc
}

func main() {
   http.HandleFunc("/", getHandlerFunc())
   fmt.Println("Listening...")
   log.Fatal(http.ListenAndServe(":8080", nil))
}
