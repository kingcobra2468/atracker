package filter

import ptc "github.com/kingcobra2468/atracker/internal/config"

var aircraft *[]string

// Save the filter aircraft into a protected global instance. This
// method could update the filtered list of aircraft during runtime.
func FetchAircraftFilter() {
	aircraft = ptc.LoadAircraft()
}

// Perform search to see if a given aircraft is part of a the list
// of filtered aircraft.
func checkAircraft(aft string) bool {
	for _, a := range *aircraft {
		if aft == a {
			return true
		}
	}

	return false
}
