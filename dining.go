package main

/**
* 'Dining' object that holds the name of a college and its respective items.
*/
type Dining struct {
  CollegeName string `json:"collegeName"`
  Food string `json:"food"`
}

/**
* Constructor type function that returns a new instance of an object 'Dining'
*/
func NewDining(collegeName string, food string) Dining {
  return Dining {
      CollegeName: collegeName,
      Food: food,
    }
}

