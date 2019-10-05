package PollingServices

import (
	"fmt"
	"time"
	"github.com/myriadeinc/pickaxe/src/api/monero"
	""
)

/**
*
*
@Documentation
Move this text here in a later build
camelCase -> private fields
PascalCase -> Packages, or public methods
snake_case ->
kebab-case -> don't use
*
*/

var lastBlockHeight uint64

type TemplateFetcher struct {
	Ticker      *time.Ticker
	BlockHeight uint64
	MoneroApi // Instance of monero api connection
}



var singleton TemplateFetcher

var once sync.Once

func GetInstance() *TemplateFetcher {
    once.Do(func() {
		// Instatiate the TemplateFetcher and Monero API
        singleton = &TemplateFetcher{
			Ticker time.Ticker()
			BlockHeight uint64
			}
    })
	return singleton
}

func (t *TemplateFetcher) run() {
	for {
		select {
		case <-t.Ticker.C:
			fmt.Println("call GetJobTemplate here")
			MoneroApi.GetJobTemplate()
			// TODO: Add notify subscriber after checking blockheight
		}
	}
}


func start() {
	tf := TemplateFetcher.GetInstance()
	go tf.run()
}

func stop(){


}
