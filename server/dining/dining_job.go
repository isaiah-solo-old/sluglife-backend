package dining

import (
  "time"
)

// Job continuesly scrapes the dining hall websites with a given timestep.
type Job struct {
  // timeStep is the duration which the job will wait before getting fresh data.
  // Do not modify this value directly, it is immutable.
  timeStep time.Duration
  // lastData is the most recent data scraped by the job. Do not use this directly,
  // use getDataRequest instead.
  lastData []Location
  // getDataRequest is a channel that accepts a channel of []Location.
  // It will return the more recently scraped data through the channel it was
  // passed. This is avoid race conditions with lastData.
  getDataRequest chan chan []Location
  // refreshData will make the job scrape the data again.
  refreshData chan bool
}

// NewJob creates a job that will scrape the dining hall data every timeStep.
// Note that it will scrape immediately to initialize.
func NewJob(timeStep time.Duration) Job {
  job := Job {
      timeStep: timeStep,
      lastData: ParseAll(),
      getDataRequest: make(chan chan []Location),
      refreshData: make(chan bool),
    }
  // begins running the job in the background.
  go job.run()
  return job
}

func (job Job) run() {
  // This goroutine will refresh the data after every time step.
  go func() {
      ticker := time.NewTicker(job.timeStep)
      for range ticker.C {
        job.refreshData <- true
      }
    }()
  for {
    // This select guarntees that lastData will never be read and written at the
    // same time.
    select {
      // returns the most recent data
      case returnChan := <-job.getDataRequest:
        returnChan <- job.lastData
      case <-job.refreshData:
        job.lastData = ParseAll()
    }
  }
}

// GetData returns the last list of Location pulled by this job.
func (job Job) GetData() []Location {
  returnChan := make(chan []Location)
  job.getDataRequest <- returnChan
  return <-returnChan
}

// GetMenu returns the menu of the provided dining hall, and a boolean
// indicating if it was found.
func (job Job) GetMenu(diningHall string) (Menu, bool) {
  returnChan := make(chan []Location)
  job.getDataRequest <- returnChan
  dlocs := <-returnChan
  for _, dloc := range dlocs {
    if dloc.Name == diningHall {
      return dloc.Menu, true
    }
  }
  return Menu{}, false
}

// GetNames returns a list of all the dining hall names.
func (job Job) GetNames() []string {
  returnChan := make(chan []Location)
  job.getDataRequest <- returnChan
  dlocs := <-returnChan
  names := make([]string, len(dlocs))
  for i := 0; i < len(dlocs); i++ {
    names[i] = dlocs[i].Name
  }
  return names
}
