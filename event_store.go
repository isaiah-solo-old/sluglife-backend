package main

type EventStorer interface {
  Put(e Event) error
  GetAll() ([]Event, error)
}
