// internal/notifier/notifier.go
package notifier

import (
	"github.com/pusher/pusher-http-go"
)

type Notifier struct {
	notifyChannel chan<- string
}

func notifier(notifyChannel <-chan string) {
	client := pusher.Client{
		AppId:   "PUSHER_APP_ID",
		Key:     "PUSHER_KEY",
		Secret:  "PUSHER_SECRET",
		Cluster: "PUSHER_CLUSTER",
		Secure:  true,
	}
	for {
		table := <-notifyChannel
		client.Trigger("table-"+table, "update", table)
	}
}
func New() Notifier {
	notifyChannel := make(chan string)
	go notifier(notifyChannel)
	return Notifier{notifyChannel}
}
func (notifier *Notifier) Notify(table string) {
	notifier.notifyChannel <- table
}
