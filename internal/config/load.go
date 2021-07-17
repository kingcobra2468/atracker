// Configuration load and validation help methods to be used
// for atracker.
package config

import (
	"errors"
	"fmt"

	"github.com/kingcobra2468/atracker/internal/radar"
	"github.com/spf13/viper"
)

// Read cordinate bounds from the config file and process it as a
// bounds box.
func LoadBounds() *radar.RadarBounds {
	viperTR := viper.Sub("radarBounds.topRightCords")
	viperBL := viper.Sub("radarBounds.bottomLeftCords")

	topRightCords, _ := loadCords(viperTR)
	bottomLeftCords, _ := loadCords(viperBL)
	bounds := radar.RadarBounds{TopRight: *topRightCords, BottomLeft: *bottomLeftCords}

	return &bounds
}

// Read Redis endpoint information and configure the address
// to be used by the redis client.
func LoadRedisAddr() string {
	addr := viper.GetString("redis.addr")
	port := viper.GetUint("redis.port")

	return fmt.Sprintf("%s:%d", addr, port)
}

// Read in all filter aircraft.
func LoadAircraft() *[]string {
	filter := viper.GetStringSlice("aircraft")

	return &filter
}

// Pull cordinate data from the config file if it exists.
func loadCords(v *viper.Viper) (*radar.Cords, error) {
	var cord radar.Cords
	if err := v.Unmarshal(&cord); err != nil {
		return nil, errors.New("unable to get cordinates for top right of the bounds")
	}

	return &cord, nil
}
