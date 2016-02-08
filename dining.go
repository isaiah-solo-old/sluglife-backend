package main

type Dining struct {
  CollegeName string `json:"collegeName"`
  Food string `json:"food"`
}

func NewDining(collegeName string, food string) Dining {
  return Dining {
      CollegeName: collegeName,
      Food: food,
    }
}

