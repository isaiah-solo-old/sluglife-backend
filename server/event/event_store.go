package event

// Storer allows for events to be stored and an array of all stored events
// to be retrieved.
type Storer interface {
  // Put an event into storage
  Put(e Event) error
  // GetAll retrieves all events
  GetAll() ([]Event, error)
}
