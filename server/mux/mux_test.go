package mux

import (
  "net/http"
  "net/http/httptest"
  "reflect"
  "testing"
)

// TODO cut down repetition
func TestEndpointNotFound(t *testing.T) {
  recorder := httptest.NewRecorder()
  fakeReq, err := http.NewRequest("POST", "http://www.triton.cloud/test", nil)
  if err != nil {
    t.Fatalf("Failed to create Request")
  }

  mux := New()
  mux.ServeHTTP(recorder, fakeReq)
  if recorder.Code != 404 {
    t.Fatalf("Invalid endpoint does not return 404")
  }
}

func TestGoodEndpoint(t *testing.T) {
  testStatusCode := 205
  testWrite := []byte("Works")
  fn := func(w http.ResponseWriter, r *http.Request) {
      w.WriteHeader(testStatusCode)
      w.Write(testWrite)
    }

  testEndpoint := Endpoint {
      Path: "/test",
      Method: "POST",
    }

  recorder := httptest.NewRecorder()
  fakeReq, err := http.NewRequest("POST", "http://www.triton.cloud/test", nil)
  if err != nil {
    t.Fatalf("Failed to create Request")
  }

  mux := New()
  mux.BindFn(testEndpoint, fn)
  mux.ServeHTTP(recorder, fakeReq)
  if recorder.Code != testStatusCode {
    t.Fatalf("Unexpected status code.")
  }
  if !reflect.DeepEqual(recorder.Body.Bytes(), testWrite) {
    t.Fatalf("Unexpected body.")
  }
}
