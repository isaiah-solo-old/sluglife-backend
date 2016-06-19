package event

// simpleStore is an Storer that is backed by an in-memory array.
type simpleStore struct {
  // events is all the stored events.
  events []Event
}

// NewSimpleStore creates a new Storer that uses an in-memory array
// for storage.
func NewSimpleStore() Storer {
  return &simpleStore {
      events: make([]Event, 0),
    }
}

// Put adds the event into storage, and never returns an error.
func (store *simpleStore) Put(e Event) error {
  store.events = append(store.events, e)
  return nil
}

// GetAll returns the backing array for the storer, and always returns no error.
func (store *simpleStore) GetAll() ([]Event, error) {
  return store.events, nil
}

// Put adds the event into storage, and never returns an error.
func (store *simpleStore) Delete(s string) error {
  for index,event := range store.events {
    if(event.Name == s) {
      store.events = append(store.events[:index], store.events[index + 1:]...)
      return nil
    }
  }
  return nil
}
