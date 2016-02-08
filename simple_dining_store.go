package main

type simpleDiningStore struct {
  dining []Dining
}

func NewSimpleDiningStore() DiningStorer {
  return &simpleDiningStore {
      dining: make([]Dining, 0),
    }
}

func (store *simpleDiningStore) Put(e Dining) error {
  store.dining = append(store.dining, e)
  return nil
}

func (store *simpleDiningStore) GetAll() ([]Dining, error) {
  return store.dining, nil
}
