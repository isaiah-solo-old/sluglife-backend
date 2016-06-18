package main

type Event struct {
  Name string `json:"name"`
  Summary string `json:"summary"`
  Image string `json:"image"`
}

func NewEvent(name string, summary string, image string) Event {
  return Event {
      Name: name,
      Summary: summary,
      Image: image,
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
