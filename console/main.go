package main

import (
	"fmt"
	"github.com/MaximeEngel/koryo/core"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	gb := core.GameBoard([]string{"Maxime", "Eloise", "Nathalie"})

	gb.CardDistributionPhase()
	for i := 0; i < 3; i++ {
		for player := gb.FirstPlayer(); player != nil; player = gb.NextCurrentPlayer() {
			hand := player.HandConst()
			for i := 0; i < 4; i++ {
				player.SelectPlayCardPtr(hand[i])
			}
			player.PlayCard(0)
			player.PlayCard(1)
		}
		gb.NextFirstPlayer()
	}
	for player := gb.FirstPlayer(); player != nil; player = gb.NextCurrentPlayer() {
		fmt.Println(player)
	}
	gb.NextFirstPlayer()
	for i := 0 ; i < 11; i++ {
		id := core.CardId(i)
		for player := gb.FirstPlayer(); player != nil; player = gb.NextCurrentPlayer() {
			if gb.HasMajority(player, id, true) {
				fmt.Printf("%s has %s majority\n", player.Name, core.IdToName(id))
			}
		}
		gb.NextFirstPlayer() // trick because i dont want to make a player iterator for the test
	}
	// Debug why it doesnt print
	for player := gb.FirstPlayer(); player != nil; player = gb.NextCurrentPlayer() {
		fmt.Printf("%s score: %d\n", player.Name, gb.ScoringPlayer(player))
	}

	for _, s := range gb.Scoring() {
		fmt.Printf("%s score : %d \n", s.P.Name, s.S)
	}
}
