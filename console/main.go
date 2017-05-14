package main

import (
	"fmt"
	"github.com/MaximeEngel/koryo/core"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	deck := core.Deck()
	deck.Shuffle()
	fmt.Println(deck)
	player1 := core.Player("Maxime")
	for i := 0; i < 3; i++ {
		player1.Draw(deck.Draw())
	}
	fmt.Println(deck)
	fmt.Println(player1)
}
