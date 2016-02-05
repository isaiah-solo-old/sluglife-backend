package main

import (
  "reflect"
  "testing"
)

func TestSimpleEventStorer(t *testing.T) {
  eventStorerTests(t, NewSimpleEventStore())
}

func eventStorerTests(t *testing.T, store EventStorer) {
  if store.GetAll() != []Event{} {
    t.Fatalf("Invalid start state")
  }

  expected := []Event {
      Event{name: "a", summary: "d"},
      Event{name: "b", summary: "e"},
      Event{name: "c", summary: "f"},
    }

  for _, value := range expected {
    store.Put(value)
  }

}

func containSameValues(this []interface, that []interface) {

}

func contains(haystack []interface{}, needle interface{}) {

}
