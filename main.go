package main

import (
  "fmt"
  "net/http"
  "./server"
  "./server/event"
  "./server/map"
)

func main() {
  fmt.Printf("Listening...\n")
  http.ListenAndServe(":8080", server.New(event.NewSimpleStore(), maps.NewSimpleStore()))
}
