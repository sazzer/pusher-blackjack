package main

import (
	"pusher/blackjack/internal/chatter"
	"pusher/blackjack/internal/game"
	"pusher/blackjack/internal/notifier"
	"pusher/blackjack/internal/webapp"
)

func main() {
	chatter := chatter.New()
	notifier := notifier.New()

	games := game.New(&chatter)
	webapp.StartServer(&chatter, &notifier, &games)
}
