// Configuration load and validation help methods to be used
// for atracker.
package config

import (
	"errors"

	"github.com/kingcobra2468/atracker/internal/util"
	"github.com/spf13/viper"
)

// Read cordinate bounds from the config file and validate that the cords
// are valid and within bounds of the correct axis.
func CheckBounds() error {
	viperTR := viper.Sub("radarBounds.topRightCords")
	viperBL := viper.Sub("radarBounds.bottomLeftCords")

	topRightCords, err := loadCords(viperTR)
	if err != nil {
		return errors.New("unable to get cordinates for top right of the bounds")
	}
	bottomLeftCords, err := loadCords(viperBL)
	if err != nil {
		return errors.New("unable to get cordinates for bottom left of the bounds")
	}

	if !util.ValidateCord(*topRightCords) || !util.ValidateCord(*bottomLeftCords) {
		return errors.New("cordinate bounds are not within a valid range")
	}

	return nil
}
