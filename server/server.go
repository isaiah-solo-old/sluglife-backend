package server

import (
  "encoding/json"
  "net/http"
  "./event"
  "./mux"
  "./dining"
  "html/template"
  "time"
)

// The time between the dining job of a Server retrieving fresh data.
const (
  defaultDiningTimeStep = time.Hour
)

type Server struct {
  eventStore event.Storer
  diningJob dining.Job
  mx mux.Mux
}

func New(eventStore event.Storer) Server {
  server := Server {
      eventStore: eventStore,
      diningJob: dining.NewJob(defaultDiningTimeStep),
    }
  mx := mux.New()
  mx.BindFn(mux.Endpoint {
      Method: "POST",
      Path: "/event",
    }, server.postEvent)

  mx.BindFn(mux.Endpoint {
      Method: "GET",
      Path: "/event",
    }, server.getEvents)

  mx.BindFn(mux.Endpoint {
      Method: "GET",
      Path: "/dining",
    }, server.getDiningHalls)

  mx.BindFn(mux.Endpoint {
      Method: "GET",
      Path: "/dining/menu",
    }, server.getMenu)
  server.mx = mx
  return server
}

func (server Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  // Will be removed once we get a domain and no longer have to worry about
  // CORS
  // BEGIN REMOVE
  w.Header().Set("Access-Control-Allow-Origin", "*")
  if r.Method == "OPTIONS" {
    return
  }
  // END REMOVE
  server.mx.ServeHTTP(w, r)
}

func (server Server) postEvent(w http.ResponseWriter, r *http.Request) {
  // Escapes the strings to avoid XSS attacks
  name := template.HTMLEscapeString(r.FormValue("name"))
  desc := template.HTMLEscapeString(r.FormValue("description"))
  if name == "" {
    w.WriteHeader(400)
    w.Write([]byte("Name parameter not provided"))
    return
  }
  if desc == "" {
    w.WriteHeader(400)
    w.Write([]byte("Description parameter not provided"))
    return
  }

  putErr := server.eventStore.Put(event.Event{
      Name: name,
      Summary: desc,
    })
  if putErr != nil {
    w.WriteHeader(500)
    w.Write([]byte("Server failed to store event"))
    return
  }
}

func (server Server) getEvents(w http.ResponseWriter, r *http.Request) {
  events, getErr := server.eventStore.GetAll()
  if getErr != nil {
    w.WriteHeader(500)
    w.Write([]byte("Server failed to get events"))
    return
  }

  eventsJson, jsonErr := json.Marshal(events)
  if jsonErr != nil {
    w.WriteHeader(500)
    w.Write([]byte("Failed to turn events into JSON"))
    return
  }
  w.Write(eventsJson)
}

func (server Server) getDiningHalls(w http.ResponseWriter, r *http.Request) {
  names := server.diningJob.GetNames()
  namesJson, jsonErr := json.Marshal(names)
  if jsonErr != nil {
    w.WriteHeader(500)
    w.Write([]byte("Failed to turn names into JSON"))
    return
  }
  w.Write(namesJson)
}

func (server Server) getMenu(w http.ResponseWriter, r *http.Request) {
  diningName := r.FormValue("name")
  menu, found := server.diningJob.GetMenu(diningName)
  if !found {
    w.WriteHeader(400)
    w.Write([]byte("Failed to find menu for dining hall: " + diningName))
    return
  }
  menuJson, jsonErr := json.Marshal(menu)
  if jsonErr != nil {
    w.WriteHeader(500)
    w.Write([]byte("Failed to turn menu into JSON"))
    return
  }
  w.Write(menuJson)
}
