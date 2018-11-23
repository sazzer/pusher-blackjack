package game

import (
	"math/rand"
	"time"
)

const NUMBER_OF_DECKS = 6

func (g *Game) shuffle() {
	cards := []Card{}
	for i := 0; i < NUMBER_OF_DECKS; i++ {
		for _, suit := range CARD_SUITS {
			for _, value := range CARD_VALUES {
				cards = append(cards, Card{
					Face: value,
					Suit: suit,
				})
			}
		}
	}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	g.deck = make([]Card, len(cards))
	perm := r.Perm(len(cards))
	for i, randIndex := range perm {
		g.deck[i] = cards[randIndex]
	}
}

func (g *Game) deal() Card {
	card, remainder := g.deck[0], g.deck[1:]
	g.deck = remainder
	return card
}
