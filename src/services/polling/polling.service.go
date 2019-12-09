package PollingService

import (
	"time"
	"sync"
	"github.com/myriadeinc/pickaxe/src/api/monero"
	"github.com/myriadeinc/pickaxe/src/services/subscriber"
	"github.com/myriadeinc/pickaxe/src/util/config"
	"github.com/myriadeinc/pickaxe/src/util/logger"
	"encoding/json"
	"bytes"
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
				BlockHeight: 1,
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
			var jobTemplate *MoneroApi.JobTemplateResponse
			//"Blocktemplate:  %v", ConfigUtil.Get("pool.wallet_address").(string))
			jobTemplate = MoneroApi.GetJobTemplate(8, ConfigUtil.Get("pool.wallet_address").(string))
			jobPayload, err := json.Marshal(jobTemplate)
			if err != nil {
				LoggerUtil.Logger.Error("Critical json marshal error!", err.Error())
				panic("Bad marshaling")
			} 	
			if(*jobTemplate.Height > t.BlockHeight){
			// @TODO: Compare prevHash field first (in later build)
				SubscriberService.Notify(func (subscriber SubscriberService.Subscriber) {
					// If needed check return statement for tracking faulty subscribers
					go SubscriberService.SendRequest(bytes.NewBuffer(jobPayload),subscriber)
					})
				t.setBlockHeight(*jobTemplate.Height)
				}

			}
		}
}



// rc/services/polling/polling.service.go:54:6: syntax error: unexpected newline, expecting comma or )
// pickaxe_1   | src/services/polling/polling.service.go:56:5: syntax error: unexpected ) at end of statement
// pickaxe_1   | src/services/polling/polling.service.go:61:1: syntax error: non-declaration statement outside function body

func Init() {
	MoneroApi.Init(ConfigUtil.Get("pool.monero_url").(string), true)
	var tf *TemplateFetcher = GetInstance()
	go tf.run()
}

func Shutdown(){
	LoggerUtil.Logger.Info("Shutting down TemplateFetching Service")
	var tf *TemplateFetcher = GetInstance()
	tf.Ticker.Stop()
}
