package game

import "pusher/blackjack/internal/chatter"

var CARD_VALUES = [...]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

var CARD_SUITS = [...]string{"clubs", "diamonds", "hearts", "spades"}

type Card struct {
	Suit string `json:"suit"`
	Face string `json:"face"`
}

type Player struct {
	ID    string `json:"id"`
	Bet   uint16 `json:"bet"`
	Stack uint16 `json:"stack"`
	Cards []Card `json:"cards"`
}

type Game struct {
	ID       string    `json:"id"`
	State    string    `json:"state"`
	Turn     uint16    `json:"turn"`
	Players  [6]Player `json:"players"`
	Croupier []Card    `json:"croupier"`
	deck     []Card
	chatter  *chatter.Chatter
}

type Games struct {
	games   []*Game
	chatter *chatter.Chatter
}
