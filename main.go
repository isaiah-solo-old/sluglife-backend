package main

import (
  "fmt"
  "net/http"
  "./server"
  "./server/event"
)

func main() {
  fmt.Printf("Listening...\n")
  http.ListenAndServe(":8080", server.New(event.NewSimpleStore()))
}
