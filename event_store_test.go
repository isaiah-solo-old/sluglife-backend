package main

import (
  "reflect"
  "testing"
)

func TestSimpleEventStorer(t *testing.T) {
  eventStorerTests(t, NewSimpleEventStore())
}

func eventStorerTests(t *testing.T, store EventStorer) {
  initEvents, _ := store.GetAll()
  if !reflect.DeepEqual(initEvents, []Event{}) {
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

  events, _ := store.GetAll()
  containSameValues([]interface{}(events, expected)
}

func containSameValues(this, that []interface{}) bool {
  if len(this) != len(that) {
    return false
  }

  for _, value := range this {
    if !contains(that, value) {
      return false
    }
  }
  return true
}

func contains(haystack []interface{}, needle interface{}) bool {
  for _, value := range haystack {
    if reflect.DeepEqual(value, needle) {
      return true
    }
  }
  return false
}
