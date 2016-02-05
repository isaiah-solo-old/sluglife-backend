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
/*
func (e Event) Name() string {
  return e.Name
}

func (e Event) Summary() string {
  return e.Summary
}
*/
