package SubscriberService

import "fmt"

type Subscriber struct {
	Webhook		string
}

// Implementing a pseudo observer pattern

var subscribers []*Subscriber

func Subscribe(webhook string) {
	subscribers = append(subscribers, &Subscriber{webhook})
}

func Notify(notifyFn func(subscriber Subscriber)) {
	fmt.Println(subscribers)
	for _, subscriber := range subscribers{
		notifyFn(*subscriber)
	}
}