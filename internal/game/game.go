package game

import (
	"errors"
	"fmt"
	"strconv"
)

const GAME_STATE_BETTING = "betting"
const GAME_STATE_PLAYING = "playing"

func (g *Game) Bet(id string, amount uint16) error {
	if g.State != GAME_STATE_BETTING {
		return errors.New("The betting phase is over")
	}

	for index, s := range g.Players {
		if s.ID == id {
			if g.Players[index].Bet != 0 {
				return errors.New("You have already placed a bet")
			}
			if amount > g.Players[index].Stack {
				return errors.New("You can not bet that much")
			}
			g.Players[index].Bet = amount
			g.Players[index].Stack -= amount
			g.Players[index].Cards = []Card{}
			g.Croupier = []Card{}
			g.chatter.Message(g.ID, fmt.Sprintf("%s has bet %d", id, amount))
		}
	}

	stillBetting := false
	players := 0

	for _, player := range g.Players {
		if player.Bet == 0 && player.ID != "" {
			stillBetting = true
			players++
		}
	}

	if stillBetting == false {
		g.State = GAME_STATE_PLAYING

		if len(g.deck) < 5*(players+1) { // We need at least 5 cards per player to be able to play
			g.shuffle()
		}

		for index, player := range g.Players {
			if player.ID != "" {
				g.Players[index].Cards = append(g.Players[index].Cards, g.deal())
				g.Players[index].Cards = append(g.Players[index].Cards, g.deal())
			}
		}

		g.Croupier = append(g.Croupier, g.deal())

		for index, player := range g.Players {
			if player.ID != "" {
				g.Turn = uint16(index)
				g.chatter.Message(g.ID, fmt.Sprintf("%s's turn to act", player.ID))
				break
			}
		}
	}

	return nil
}

func (g *Game) Hit(id string) error {
	if g.State != GAME_STATE_PLAYING {
		return errors.New("We are not currently acting")
	}

	player := &(g.Players[g.Turn])
	if player.ID != id {
		return errors.New("It is not your turn")
	}

	player.Cards = append(player.Cards, g.deal())

	newScore := calculateScore(player.Cards)
	if newScore > 21 {
		// Bust. Next players turn
		g.chatter.Message(g.ID, fmt.Sprintf("%s is bust", player.ID))
		return g.Stick(id)
	} else if len(player.Cards) == 5 {
		// 5 Cards. Next players turn
		g.chatter.Message(g.ID, fmt.Sprintf("%s has 5 cards", player.ID))
		return g.Stick(id)
	}

	return nil
}

func (g *Game) Stick(id string) error {
	if g.State != GAME_STATE_PLAYING {
		return errors.New("We are not currently acting")
	}

	player := g.Players[g.Turn]
	if player.ID != id {
		return errors.New("It is not your turn")
	}

	foundNext := false
	for index, player := range g.Players {
		if uint16(index) > g.Turn && player.ID != "" {
			g.Turn = uint16(index)
			g.chatter.Message(g.ID, fmt.Sprintf("%s's turn to act", player.ID))
			foundNext = true
			break
		}
	}

	if foundNext == false {
		// The round has finished
		g.State = GAME_STATE_BETTING

		for calculateScore(g.Croupier) < 16 {
			g.Croupier = append(g.Croupier, g.deal())
		}

		croupierScore := calculateScore(g.Croupier)
		if croupierScore > 21 {
			g.chatter.Message(g.ID, "The house is bust")
		}
		for index, player := range g.Players {
			if player.ID != "" {
				playerScore := calculateScore(g.Players[index].Cards)
				if playerScore == 21 {
					// Blackjack
					g.chatter.Message(g.ID, fmt.Sprintf("%s got Blackjack!", player.ID))
					g.Players[index].Stack += g.Players[index].Bet * 3
				} else if playerScore <= 21 {
					if croupierScore > 21 || playerScore > croupierScore {
						// Player has won
						g.chatter.Message(g.ID, fmt.Sprintf("%s has won", player.ID))
						g.Players[index].Stack += g.Players[index].Bet * 2
					} else if playerScore == croupierScore {
						// Scores match
						g.chatter.Message(g.ID, fmt.Sprintf("%s has drawn. Money back.", player.ID))
						g.Players[index].Stack += g.Players[index].Bet
					}
				}
				g.Players[index].Bet = 0
			}
		}

		g.chatter.Message(g.ID, "Bets please.")

	}

	return nil
}

func calculateScore(cards []Card) uint16 {
	var score uint64
	aces := 0

	for _, card := range cards {
		switch card.Face {
		case "A":
			score += 11
			aces++
		case "J", "Q", "K":
			score += 10
		default:
			parsedScore, _ := strconv.ParseUint(card.Face, 10, 16)
			score += parsedScore
		}
	}

	for i := 0; i < aces; i++ {
		if score > 21 {
			score -= 10
		}
	}

	return uint16(score)
}
