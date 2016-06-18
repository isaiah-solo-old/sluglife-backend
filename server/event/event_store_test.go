package event

import (
  "reflect"
  "testing"
)

func TestSimpleStorer(t *testing.T) {
  eventStorerTests(t, NewSimpleStore())
}

func eventStorerTests(t *testing.T, store Storer) {
  initEvents, _ := store.GetAll()
  if !reflect.DeepEqual(initEvents, []Event{}) {
    t.Fatalf("Invalid start state")
  }

  expected := []Event {
      Event{Name: "a", Summary: "d"},
      Event{Name: "b", Summary: "e"},
      Event{Name: "c", Summary: "f"},
    }

  for _, value := range expected {
    store.Put(value)
  }

  events, _ := store.GetAll()
  if !containSameValues([]interface{}{events}, []interface{}{expected}) {
    t.Fatalf("GetAll does not return expected events")
  }
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
