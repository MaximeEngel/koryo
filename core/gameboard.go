package core

import (
	"math/rand"
)

type gameBoard struct {
	players []*player
	first_player_idx int
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

	return &gameBoard{
		players:players,
		first_player_idx:rand.Intn(len(players)),
		deck:deck,
		season:1}
}
