// Filter all aircraft and perform alerts when aircraft are found.
// The filtered aircraft are the same as the one set inside of config.yaml file.
package filter

import (
	"fmt"

	"github.com/kingcobra2468/atracker/internal/cache"
	"github.com/kingcobra2468/atracker/internal/radar"
)

// Callback method for TrackAll that filters down all aircraft
// and propogates message when a filtered aircraft is found.
func FilterAircraft(fid radar.FlightID, ad radar.AircraftData) {
	if tked, err := cache.CheckPlaneTracked(fid); tked || err != nil {
		return
	}
	if !checkAircraft(ad.Act) {
		return
	}

	fmt.Printf("Found a filtered aircraft: %s\n", ad.Act)
	cache.TrackNew(fid) // cache result to avoid duplicate alerts for same aircraft
}
