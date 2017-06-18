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
	for player := gb.FirstPlayer(); player != nil ; player = gb.NextCurrentPlayer() {
		fmt.Println(player)
	}
}
