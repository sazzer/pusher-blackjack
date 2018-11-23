package game

import "pusher/blackjack/internal/chatter"

func New(chatter *chatter.Chatter) Games {
	return Games{
		games:   []*Game{},
		chatter: chatter,
	}
}

func (g *Games) newGame(id string) *Game {
	return &Game{
		ID:      id,
		State:   GAME_STATE_BETTING,
		Turn:    0,
		Players: [6]Player{},
		deck:    []Card{},
		chatter: g.chatter,
	}
}

func (g *Games) GetGame(id string) *Game {
	var result *Game

	for _, v := range g.games {
		if v.ID == id {
			result = v
			break
		}
	}

	if result == nil {
		result = g.newGame(id)
		g.games = append(g.games, result)
	}

	return result
}
