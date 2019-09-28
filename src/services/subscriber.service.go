package SubscriberService

import (
	"github.com/Kairi/godash"
)

type Subscriber struct {
	Hostname		string
}
type notify func(Subscriber) bool

var subscribers []*Subscriber

func Subscribe(hostname string) () {
		append(subscribers, &Subscriber{hostname})
}

func Notify(fn notify) ([]bool) {
	var results []bool = godash.Map(subscribers, fn) 
	return results
}