package mux

import (
  "net/http"
)

type Endpoint struct {
  Path string
  Method string
}

type Mux struct {
  fnMap map[Endpoint]http.Handler
}

func New() Mux {
  return Mux {
      fnMap: make(map[Endpoint]http.Handler),
    }
}

func (mux Mux) Bind(endpoint Endpoint, handler http.Handler) {
  mux.fnMap[endpoint] = handler
}

func (mux Mux) BindFn(
    endpoint Endpoint, fn func(w http.ResponseWriter, r *http.Request)) {
  mux.Bind(endpoint, http.HandlerFunc(fn))
}

func (mux Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  endpoint := Endpoint {
      Path: r.URL.Path,
      Method: r.Method,
    }

  val, found := mux.fnMap[endpoint]
  if !found {
    w.WriteHeader(404)
    return
  }
  val.ServeHTTP(w, r)
}
