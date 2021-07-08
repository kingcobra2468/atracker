package main

import (
	"errors"

	ptc "github.com/kingcobra2468/atracker/internal/config"

	"github.com/spf13/viper"
)

func loadConfig() (*ptc.Config, error) {
	viperTR := viper.Sub("radarBounds.topRightCords")
	viperBL := viper.Sub("radarBounds.bottomLeftCords")

	bounds, err := ptc.LoadBounds(viperTR, viperBL)
	if err != nil {
		return nil, err
	}
	planes := viper.GetStringSlice("planes")
	if err != nil {
		return nil, errors.New("failed to get list of planes to filter by")
	}

	return &ptc.Config{Bounds: *bounds, Planes: planes}, nil
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}
}
