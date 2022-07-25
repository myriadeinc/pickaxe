package config

import (
	"github.com/spf13/viper"
)

func DefaultConfigs() {

	viper.SetDefault("WEBHOOK_URL", "http://zircon:4990/new")
	viper.SetDefault("WALLET_ADDRESS", "47PAULmUFo3DPHKehGPuxXbEAB4JkRYJ49DEFs4EqaT7M2TRqqWWHAeJyEHWg8eRoWNwMAHh7bx6Eh5SR2fpdnj71fhxugC")
	viper.SetDefault("RESERVE_OFFSET", 8)

}
