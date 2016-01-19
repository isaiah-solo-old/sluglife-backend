package main

import (
   "fmt"
   "log"
   "net/http"
)

type Event struct {
   Name string `json:"name"`
   Summary string `json:"summary"`
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
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

func main() {
   http.HandleFunc("/", handlerFunc)
   fmt.Println("Listening...")
   log.Fatal(http.ListenAndServe(":8080", nil))
}
