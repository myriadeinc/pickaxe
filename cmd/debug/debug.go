package main

import (
	"fmt"

	"github.com/myriadeinc/pickaxe/internal/config"
	"github.com/myriadeinc/pickaxe/internal/nodeapi"
	"github.com/myriadeinc/pickaxe/internal/poller"

	"github.com/rs/zerolog/log"
)

func main() {
	config.DefaultConfigs()
	fmt.Println("debug mode")
	n := nodeapi.NewNodeClient([]string{"https://node.monerod.org/json_rpc"})
	template := n.GetRawBlockTemplate()

	log.Info().Msgf("template %s", debugKeys(template))

	p := poller.NewPoller()
	p.PollForever()
}

func debugKeys(m map[string]interface{}) string {
	s := ""
	for k, _ := range m {
		s = s + k + " | "
	}
	return s

}
