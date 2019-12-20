package SubscriberService

import (
	"io"
	"fmt"
	"strconv"
	"net/http"
	"github.com/myriadeinc/pickaxe/src/util/config"
	"github.com/myriadeinc/pickaxe/src/util/logger"
)

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

func SendRequest(data io.Reader, subscriber Subscriber) bool {

	req, err := http.NewRequest("POST", subscriber.Webhook, data)

	req.Header.Add("Authorization", ConfigUtil.Get("service.shared_secret").(string) ) 
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		LoggerUtil.Logger.Error("Subscriber Request Error: ", err)
		return false
	}

	defer resp.Body.Close()
	if(resp.StatusCode == 200){
		// No news is good news
		return true
	}
	
	LoggerUtil.Logger.Error( fmt.Sprint("Bad Status Code: ",strconv.Itoa(resp.StatusCode)," Subscriber: ",subscriber.Webhook) )
	return false

}