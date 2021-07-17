// Utilities that are used within atracker.
package util

import "github.com/kingcobra2468/atracker/internal/radar"

// Validate that both of a cordinate's latitude and longitude
// are within bounds.
func ValidateCord(c radar.Cords) bool {
	return ValidateLatitude(c) && ValidateLongitude(c)
}

// Validate that a cordinate's latitude is within bounds.
func ValidateLatitude(c radar.Cords) bool {
	if -90 > c.Latitude || c.Latitude > 90 {
		return false
	}
	return true
}

// Validate that a cordinate's longitude is within bounds.
func ValidateLongitude(c radar.Cords) bool {
	if -180 > c.Latitude || c.Latitude > 180 {
		return false
	}
	return true
}
