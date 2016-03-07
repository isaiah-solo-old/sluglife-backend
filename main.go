package main

import (
  "net/http"
  "./server"
  "./server/event"
)

func main() {
  http.ListenAndServe(":8080", server.New(event.NewSimpleStore()))
}
