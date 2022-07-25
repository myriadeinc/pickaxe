package poller

import (
	"time"

	"github.com/rs/zerolog/log"

	"github.com/myriadeinc/pickaxe/internal/cache"
	"github.com/myriadeinc/pickaxe/internal/nodeapi"
	"github.com/myriadeinc/pickaxe/internal/updater"
)

type Poller struct {
	client      nodeapi.NodeApi
	cache       cache.CacheService
	blockHeight uint64
}

func NewPoller() Poller {
	nodes := []string{"https://node.monerod.org/json_rpc"}

	client := nodeapi.NewNodeClient(nodes)
	cache := cache.NewDummyClient()

	return Poller{
		client:      client,
		cache:       cache,
		blockHeight: 0,
	}
}

func (p *Poller) PollForever() {
	for {
		template := p.client.GetRawBlockTemplate()
		// Case of bad startup
		if len(template) == 0 {
			log.Error().Msg("Could not get block template from node")
			time.Sleep(10 * time.Second)
			continue
		}

		height := uint64(template["height"].(float64))

		if height <= p.blockHeight {
			time.Sleep(10 * time.Second)
			log.Debug().Uint64("height", height).Uint64("pollerheight", height).Msg("skipping lower height")
			continue
		}

		log.Info().Uint64("height", height).Msg("new blockheight detected")
		p.blockHeight = height

		err := p.cache.SaveNewTemplate(template)
		if err != nil {
			log.Error().Err(err).Msg("could not save newtemplate")
			time.Sleep(100 * time.Second)
			continue
		}
		log.Info().Msg("Saved new blocktemplate")
		// Fire and forget
		updater.UpdateWebhooks()

		time.Sleep(10 * time.Second)

	}

}
