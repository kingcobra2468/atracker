// Client for the \R\a\d\a/r/b/o/x/ API. Get planes in the area
// and metadata about each aircraft.
package radar

import (
	"fmt"
	"math/rand"
	"time"
)

// Callback function that is called each time an aircraft is
// detected within the bounds.
type onDetected func(fid FlightID, ad AircraftData)

// Track all of the aircraft in the area and return the aircraft data
// to a callback function that matches the onDetected signature.
func (rb RadarBounds) TrackAll(done <-chan bool, prevCaptcha bool, od onDetected) {
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
			fmt.Println(flights)
			for _, fid := range *flights {
				concurrentDetections <- struct{}{}
				go func(fid FlightID, pCaptcha bool) {
					if pCaptcha {
						r := rand.Intn(10)
						time.Sleep(time.Duration(r) * time.Second)
					}
					rb.processDetected(fid, concurrentDetections, od)
				}(fid, prevCaptcha)
			}
		}
	}
}

// Goroutine that is called whenever an aircraft is detected. Wraps the onDetect
// callback around the worker pool.
func (rb RadarBounds) processDetected(fid FlightID, hold <-chan struct{}, od onDetected) {
	defer func() { <-hold }()
	fi, err := rb.FlightInfo(fid)
	if err != nil {
		fmt.Println(err)
		return
	}

	od(fid, *fi)
}
