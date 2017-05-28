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
	for i := 0; i < 8; i++ {
		player1.Draw(deck.Draw())
	}
	fmt.Println(deck)
	fmt.Println(player1)
	fmt.Println(player1.SelectPlayCard(0))
	fmt.Println(player1.SelectPlayCard(0))
	fmt.Println(player1.SelectPlayCard(1))
	fmt.Println(player1.SelectPlayCard(2))
	fmt.Println(player1)
}
