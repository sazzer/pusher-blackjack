package game

import "errors"

func (g *Game) JoinTable(seat uint16, id string) (*Player, error) {
	if g.Players[seat].ID == "" {
		g.Players[seat] = Player{
			ID:    id,
			Bet:   0,
			Stack: 100,
			Cards: []Card{},
		}
		return &(g.Players[seat]), nil
	} else {
		return nil, errors.New("That seat is already taken")
	}
}

func (g *Game) LeaveTable(id string) {
	for index, s := range g.Players {
		if s.ID == id {
			g.Players[index] = Player{}
		}
	}
}
