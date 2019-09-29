package PollingServices

import (
	"fmt"
	"time"
)
/**
*
*
camelCase -> private fields
PascalCase -> Packages, or public methods
snake_case -> 
kebab-case -> don't use
*
*/

var lastBlockHeight uint64

type TemplateFetcher struct {
	Ticker 			*time.Ticker
	BlockHeight		uint64
}

type blockTemplate struct {
	
}

func (t *templateFetcher) run() {
	fmt.Println("Run called")
}

func (t *TemplateFetcher) NewTemplateFetcher (){ 
	var tf *templateFetcher = &templateFetcher{
		time.NewTicker(2000 * time.Millisecond),
		currentHeight,
	}
	return tf
	// go tf.run()
} 



// func Start() {
//     ticker := time.NewTicker(500 * time.Millisecond)
//     done := make(chan bool)
//     go func() {
//         for {
//             select {
//             case <-done:
//                 return
//             case t := <-ticker.C:
//                 fmt.Println("Tick at", t)
//             }
//         }
//     }()

//     time.Sleep(1600 * time.Millisecond)
//     ticker.Stop()
//     done <- true
//     fmt.Println("Ticker stopped")
// }

func processLatestBlock(newBlockHeight uint64, newBlockTemplate blockTemplate) {
	if (newBlockHeight > lastBlockHeight){
		lastBlockHeight = newBlockHeight
			
	}
}
