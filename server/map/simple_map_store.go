package maps

// simpleStore is an Storer that is backed by an in-memory array.
type simpleStore struct {
  // maps is all the stored locations.
  maps []Map
}

// NewSimpleStore creates a new Storer that uses an in-memory array
// for storage.
func NewSimpleStore() Storer {
  return &simpleStore {
      maps: make([]Map, 0),
    }
}

// Put adds the lcoation into storage, and never returns an error.
func (store *simpleStore) Put(e Map) error {
  store.maps = append(store.maps, e)
  return nil
}

// GetAll returns the backing array for the storer, and always returns no error.
func (store *simpleStore) GetAll() ([]Map, error) {
  return store.maps, nil
}
