package main

import (
  "encoding/json"
  "net/http"
  "fmt"
)

var eventStore EventStorer = NewSimpleEventStore()
var diningStore DiningStorer = NewSimpleDiningStore()

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


func handleDining(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  fmt.Println("Recieved dining HTTP")
  collegeFood := ""
  if r.Method == "GET" {
    data, _ := diningStore.GetAll();
    jsonDat, _ := json.Marshal(data)

    byt := []byte(jsonDat)
    var dat []map[string]interface{}

    if err := json.Unmarshal(byt, &dat); err != nil{
      panic(err)
    } 
    fmt.Println(dat)

    for i := 0; i < len(dat); i+=1{
      if dat[i]["name"].(string) == r.FormValue("collegeName") {
        collegeFood = dat[i]["items"].(string)
        break
      }
    }

    jsonDat, _ = json.Marshal(collegeFood)
    w.Write(jsonDat)
  } else {
    w.WriteHeader(404)
  }
}

func main() {
  //http.HandleFunc("/event", handleEvent);
  http.HandleFunc("/dining", handleDining);
  http.ListenAndServe(":8080", nil)
}
