package PollingService

import (
	"fmt"
	"time"
	"sync"
	"github.com/myriadeinc/pickaxe/src/api/monero"
	"github.com/myriadeinc/pickaxe/src/services/subscriber"
	"github.com/myriadeinc/pickaxe/src/util/config"
	"github.com/myriadeinc/pickaxe/src/util/logger"
)


type TemplateFetcher struct {
	Ticker      *time.Ticker
	BlockHeight uint64
}
var singleton *TemplateFetcher
var once sync.Once

func GetInstance() *TemplateFetcher {
    once.Do(func() {
		// Instatiate the TemplateFetcher and Monero API
			singleton = &TemplateFetcher{
			// Every 2 seconds
				Ticker: time.NewTicker(2000 * time.Millisecond),
				BlockHeight: 100,
			}
    })
	return singleton
}

func (t *TemplateFetcher) setBlockHeight(blockheight uint64){
	t.BlockHeight = blockheight
}

func (t *TemplateFetcher) run() {
	for {
		select {
		case <-t.Ticker.C:
			// fmt.Println("Ticker called")
			// fmt.Println(*t)
			var jobTemplate *MoneroApi.JobTemplateResponse 			
			jobTemplate = MoneroApi.GetJobTemplate(8, ConfigUtil.Get("pool.wallet_address").(string))
			// fmt.Println(t.BlockHeight, jobTemplate.Height)
			// @TODO: Compare prevHash field first (in later build)
			if (t.BlockHeight < jobTemplate.Height) {
				SubscriberService.Notify(func (subscriber SubscriberService.Subscriber) {
					fmt.Println("Notifying subscriber")
					fmt.Println(subscriber)
				})	
				t.setBlockHeight(jobTemplate.Height)
			}
		}
	}
}


func Init() {
	MoneroApi.Init(ConfigUtil.Get("pool.monero_url").(string), true)
	var tf *TemplateFetcher = GetInstance()
	tf.setBlockHeight(MoneroApi.GetJobTemplate(8,ConfigUtil.Get("pool.wallet_address").(string)).Height)
	go tf.run()
}

func Shutdown(){
	LoggerUtil.Logger.Info("Shutting down TemplateFetching Service")
	var tf *TemplateFetcher = GetInstance()
	tf.Ticker.Stop()
}
