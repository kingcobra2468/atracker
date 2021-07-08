// Client for the \R\a\d\a/r/b/o/x/ API. Get planes in the area
// and metadata about each aircraft.
package radar

import (
	"fmt"
	"time"
)

// Callback function that is called each time an aircraft is
// detected within the bounds
type onDetected func(fid FlightID, ad AircraftData)

func (rb RadarBounds) TrackAll(done <-chan bool, od onDetected) {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	concurrentDetections := make(chan struct{}, 5)

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			flights, err := rb.DetectFlights()
			if err != nil {
				fmt.Println("unable to scan for aircraft")
				continue
			}
			for _, fid := range *flights {
				concurrentDetections <- struct{}{}
				go rb.processDetected(fid, concurrentDetections, od)
			}
		}
	}
}

// Goroutine that is called whenever an aircraft is detected. Wraps the onDetect
// callback around the worker pool.
func (rb RadarBounds) processDetected(fid FlightID, hold <-chan struct{}, od onDetected) {
	defer func() { <-hold }()
	fi, _ := rb.FlightInfo(fid)
	od(fid, *fi)
}
