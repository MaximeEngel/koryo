package core

import (
	"math/rand"
)

type gameBoard struct {
	players []*player
	first_player_idx int
	current_player_idx int
	deck *cards
	season int
}

func GameBoard(players_name []string) *gameBoard {
	players := make([]*player, 0, len(players_name))
	for _, name := range players_name {
		players = append(players, Player(name))
	}
	deck := Deck()
	deck.Shuffle()
	first_player := rand.Intn(len(players))

	return &gameBoard{
		players:players,
		first_player_idx: first_player,
		current_player_idx: first_player,
		deck:deck,
		season:1}
}

func (gb *gameBoard) FirstPlayer() *player {
	return gb.players[gb.first_player_idx]
}

// It finish the table turn, so the current player is also updated for the same player
func (gb *gameBoard) NextFirstPlayer() *player {
	next_idx := gb.first_player_idx + 1
	if next_idx >= len(gb.players) {
		next_idx = 0
	}

	gb.first_player_idx = next_idx
	gb.current_player_idx = next_idx
	return gb.FirstPlayer()
}

func (gb *gameBoard) CurrentPlayer() *player {
	return gb.players[gb.current_player_idx]
}

func (gb *gameBoard) NextCurrentPlayer() *player {
	next_idx := gb.current_player_idx + 1
	if next_idx >= len(gb.players) {
		next_idx = 0
	}

	if next_idx == gb.first_player_idx {
		return nil
	}

	gb.current_player_idx = next_idx
	return gb.CurrentPlayer()
}

func (gb *gameBoard) CardDistributionPhase() {
	base_nb_cards := NbSeasonCardDistribution(gb.season)
	for _, player := range gb.players {
		extra := 0 // TODO broadcasters majority power
		nb_cards := base_nb_cards + extra
		for i := 0; i < nb_cards; i++ {
			player.Draw(gb.deck.Draw())
		}
	}
}

func NbSeasonCardDistribution(season int) int {
	return 11 - season
}
