package maps

// Storer allows for locations to be stored and an array of all stored locations
// to be retrieved.
type Storer interface {
  // Put a map into storage
  Put(e Map) error
  // GetAll retrieves all maps
  GetAll() ([]Map, error)
}
