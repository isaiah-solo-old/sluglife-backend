package main

type simpleEventStore struct {
  events []Event
}

func NewSimpleEventStore() EventStorer {
  return &simpleEventStore {
      events: make([]Event, 0),
    }
}

func (store *simpleEventStore) Put(e Event) error {
  store.events = append(store.events, e)
  return nil
}

func (store *simpleEventStore) GetAll() ([]Event, error) {
  return store.events, nil
}
