package main

type DiningStorer interface {
  GetAll() ([]Dining, error)
}
