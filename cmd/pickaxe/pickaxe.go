package main

import (
	"os"

	"github.com/spf13/viper"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/myriadeinc/pickaxe/internal/config"
	"github.com/myriadeinc/pickaxe/internal/poller"
)

func main() {
	config.DefaultConfigs()
	viper.AutomaticEnv()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	p := poller.NewPoller()

	log.Info().Msg("Initialized pickaxe")

	p.PollForever()
}
