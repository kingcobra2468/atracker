package main

import (
	"errors"

	"github.com/kingcobra2468/atracker/internal/cache"
	ptc "github.com/kingcobra2468/atracker/internal/config"
	"github.com/kingcobra2468/atracker/internal/filter"

	"github.com/spf13/viper"
)

// Validation on whether the config is correct. Check to see
// if cordinates are valid and if there exists at least one plane
// that is being filtered.
func checkConfig() error {
	err := ptc.CheckBounds()
	if err != nil {
		panic(err)
	}
	if planes := viper.GetStringSlice("aircraft"); len(planes) == 0 {
		return errors.New("failed to get list of planes to filter by")
	}

	return nil
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if err := checkConfig(); err != nil {
		panic(err)
	}

	cache.Connect(ptc.LoadRedisAddr())
	filter.FetchAircraftFilter()
	bounds := ptc.LoadBounds()
	done := make(chan bool)
	bounds.TrackAll(done, true, filter.FilterAircraft)
}
