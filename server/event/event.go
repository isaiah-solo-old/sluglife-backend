// package event provides primitives to store and retrieve events.
package event

type Event struct {
  Name string `json:"name"`
  Date string `json:"date"`
  Summary string `json:"summary"`
  Image string `json:"image"`
}
