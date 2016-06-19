package server

import (
  "encoding/json"
  "net/http"
  "./event"
  "./map"
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
  mapStore maps.Storer
  mx mux.Mux
}

func New(eventStore event.Storer, mapStore maps.Storer) Server {
  server := Server {
      eventStore: eventStore,
      diningJob: dining.NewJob(defaultDiningTimeStep),
      mapStore: mapStore,
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
      Method: "POST",
      Path: "/map",
    }, server.postMaps)  

  mx.BindFn(mux.Endpoint {
      Method: "GET",
      Path: "/map",
    }, server.getMaps)

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
  image := template.HTMLEscapeString(r.FormValue("image"))
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
  if image == "" {
    w.WriteHeader(400)
    w.Write([]byte("Image parameter not provided"))
    return
  }

  putErr := server.eventStore.Put(event.Event{
      Name: name,
      Summary: desc,
      Image: image,
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

func (server Server) postMaps(w http.ResponseWriter, r *http.Request) {
  // Escapes the strings to avoid XSS attacks
  name := template.HTMLEscapeString(r.FormValue("name"))
  longitude := template.HTMLEscapeString(r.FormValue("longitude"))
  latitude := template.HTMLEscapeString(r.FormValue("latitude"))
  if name == "" {
    w.WriteHeader(400)
    w.Write([]byte("name parameter not provided"))
    return
  }
  if longitude == "" {
    w.WriteHeader(400)
    w.Write([]byte("longitude parameter not provided"))
    return
  }
  if latitude == "" {
    w.WriteHeader(400)
    w.Write([]byte("latitude parameter not provided"))
    return
  }

  putErr := server.mapStore.Put(maps.Map{
      Name: name,
      Longitude: longitude,
      Latitude: latitude,
    })
  if putErr != nil {
    w.WriteHeader(500)
    w.Write([]byte("Server failed to store location"))
    return
  }
}

func (server Server) getMaps(w http.ResponseWriter, r *http.Request) {
  maps, getErr := server.mapStore.GetAll()
  if getErr != nil {
    w.WriteHeader(500)
    w.Write([]byte("Server failed to get locations"))
    return
  }

  mapsJson, jsonErr := json.Marshal(maps)
  if jsonErr != nil {
    w.WriteHeader(500)
    w.Write([]byte("Failed to turn locations into JSON"))
    return
  }
  w.Write(mapsJson)
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
