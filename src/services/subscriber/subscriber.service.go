package SubscriberService

type Subscriber struct {
	Webhook		string
}

// Implementing a pseudo observer pattern

var subscribers []*Subscriber

func Subscribe(webhook string) {
	subscribers = append(subscribers, &Subscriber{webhook})
}

func Notify(notifyFn func(Subscriber)) {
	for _, subscriber := range subscribers{
		notifyFn(*subscriber)
	}
}