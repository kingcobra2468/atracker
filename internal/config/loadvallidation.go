// Procesing and validation of the config file for ptracker.
package config

import (
	"errors"

	"github.com/kingcobra2468/atracker/internal/radar"
	"github.com/kingcobra2468/atracker/internal/util"
	"github.com/spf13/viper"
)

// Config data pulled from the configuration file after validation
// and preprocessing has been performed.
type Config struct {
	Bounds radar.RadarBounds
	Planes []string
}

// Read cordinate bounds from the config file and process it. Validate that the cords
// are valid and are within bounds of the correct axis.
func LoadBounds(trv, brv *viper.Viper) (*radar.RadarBounds, error) {
	topRightCords, err := loadCords(trv)
	if err != nil {
		return nil, errors.New("unable to get cordinates for top right of the bounds")
	}
	bottomLeftCords, err := loadCords(brv)
	if err != nil {
		return nil, errors.New("unable to get cordinates for bottom left of the bounds")
	}
	bounds := radar.RadarBounds{TopRight: *topRightCords, BottomLeft: *bottomLeftCords}

	if !util.ValidateCord(*topRightCords) || !util.ValidateCord(*bottomLeftCords) {
		return nil, errors.New("cordinate bounds are not within a valid range")
	}

	return &bounds, nil
}

// Pull cordinate data from the config file if it exists.
func loadCords(v *viper.Viper) (*radar.Cords, error) {
	var cord radar.Cords

	if err := v.Unmarshal(&cord); err != nil {
		return nil, errors.New("unable to get cordinates for top right of the bounds")
	}
	return &cord, nil
}
