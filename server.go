package main

import (
  "encoding/json"
  "net/http"
  "fmt"
  "container/list"
)

var eventStore EventStorer = NewSimpleEventStore()
var diningStore DiningStorer = NewSimpleDiningStore()

/**
* Handler Function for a 'GET' and 'POST' request for getting all the events or posting an event.
*/
func handleEvent(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  fmt.Println("Recieved event HTTP")
  if r.Method == "POST" {
    event := NewEvent(r.FormValue("name"), r.FormValue("description"))
    fmt.Printf("%+v\n", event)
    eventStore.Put(event)
  } else {
    data, _ := eventStore.GetAll();
    jsonDat, _ := json.Marshal(data)
    w.Write(jsonDat)
  }
}

/**
* Handler Function for a 'GET' request returning the food for the specific college given.
*/
func handleDiningCollege(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  fmt.Println("Recieved dining HTTP")
  collegeFood := ""
  if r.Method == "GET" {
    data, _ := diningStore.GetAll();

    for i := 0; i < len(data); i+=1{
      if data[i].CollegeName == r.FormValue("collegeName") {
        collegeFood = data[i].Food
        break
      }
    }

    jsonDat, _ := json.Marshal(collegeFood)
    w.Write(jsonDat)
  } else {
    w.WriteHeader(404)
  }
}

/**
* Handler Function for a 'GET' request returning the names of all the colleges.
*/
func handleDiningAll(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  fmt.Println("Recieved dining HTTP")
  collegeNames := list.New()
  if r.Method == "GET" {
    data, _ := diningStore.GetAll();

    for i := 0; i < len(data); i+=1{
      collegeNames.PushBack(data[i].CollegeName)
    }

    jsonDat, _ := json.Marshal(collegeNames)
    w.Write(jsonDat)
  } else {
    w.WriteHeader(404)
  }
}


func main() {
  http.HandleFunc("/event", handleEvent);
  http.HandleFunc("/dining/college", handleDiningCollege);
  http.HandleFunc("/dining/all", handleDiningAll);
  http.ListenAndServe(":8080", nil)
}
