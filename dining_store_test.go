package main

import (
  "reflect"
  "testing"
)

func TestSimpleDiningStorer(t *testing.T) {
  diningStorerTests(t, NewSimpleDiningStore())
}

func diningStorerTests(t *testing.T, store DiningStorer) {
  if store.GetAll() != []Dining{} {
    t.Fatalf("Invalid start state")
  }

  expected := []Dining {
      Dining{collegeName: "Cowell", food: "d"},
      Dining{collegeName: "Porter", food: "e"},
      Dining{collegeName: "Crown", food: "f"},
    }

  for _, value := range expected {
    store.Put(value)
  }

}

func containSameValues(this []interface, that []interface) {

}

func contains(haystack []interface{}, needle interface{}) {

}
