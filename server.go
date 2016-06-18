package main

import (
  "encoding/json"
  "net/http"
  "fmt"
)

var store EventStorer = NewSimpleEventStore()

func handleEvent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  fmt.Println("Recieved event HTTP")
  if r.Method == "POST" {
    event := NewEvent(r.FormValue("name"), r.FormValue("description"), r.FormValue("image"))
    fmt.Printf("%+v\n", event)
    store.Put(event)
  } else {
    data, _ := store.GetAll();
    jsonDat, _ := json.Marshal(data)
    w.Write(jsonDat)
  }
}

func main() {
  http.HandleFunc("/event", handleEvent);
  http.ListenAndServe(":8080", nil)
}
