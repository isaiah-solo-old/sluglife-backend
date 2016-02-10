package main

type Event struct {
  Name string `json:"name"`
  Summary string `json:"summary"`
}

func NewEvent(name, summary string) Event {
  return Event {
      Name: name,
      Summary: summary,
    }
}
