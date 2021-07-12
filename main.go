package main

import (
	"errors"

	ptc "github.com/kingcobra2468/atracker/internal/config"

	"github.com/spf13/viper"
)

func checkConfig() error {
	err := ptc.CheckBounds()
	if err != nil {
		panic(err)
	}
	if planes := viper.GetStringSlice("planes"); len(planes) == 0 {
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
}
