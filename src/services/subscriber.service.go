package Services

import (
	"github.com/Kairi/godash"
)

type Subscriber struct {
	Webhook		string
}

type notify func(Subscriber) bool

var subscribers []*Subscriber

func InitSubscriberService() (*SubscriberService) {
	
}

func Subscribe(webhook string) () {
		append(subscribers, &Subscriber{webhook})
}

func Notify(fn notify) ([]bool) {
	var results []bool = godash.Map(subscribers, fn) 
	return results
}