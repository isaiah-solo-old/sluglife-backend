package main

type simpleDiningStore struct {
  dining []Dining
}

/**
* Constructor type function that returns a new instance of an array of 'Dining' objects
*/
func NewSimpleDiningStore() DiningStorer {
  return &simpleDiningStore {
      dining: make([]Dining, 0),
    }
}

/**
* Function that returns the array of 'Dining' objects (each Dining holding a name of a college and its items)
*/
func (store *simpleDiningStore) GetAll() ([]Dining, error) {
  return store.dining, nil
}


/**
* Function that returns the array of 'Dining' objects (each Dining holding a name of a college and its items)
*/
func (store *simpleDiningStore) Put(e Dining) error {
  store.dining = append(store.dining, e)
  return nil
}
