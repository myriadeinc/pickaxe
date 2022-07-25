package updater

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func UpdateWebhooks() {
	url := viper.GetString("WEBHOOK_URL")
	_, err := http.Get(url)
	if err != nil {
		log.Error().Err(err).Str("url", url).Msg("Host did not respond")
	}

}
