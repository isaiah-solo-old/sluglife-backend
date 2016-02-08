package main

type DiningStorer interface {
  Put(e Dining) error
  GetAll() ([]Dining, error)
}
