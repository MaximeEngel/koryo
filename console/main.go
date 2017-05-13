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
	first_card := deck.Draw()
	second_card := deck.Draw()
	third_card := deck.Draw()
	fmt.Println(deck)
	fmt.Println(second_card)
	deck.Add(first_card)
	fmt.Println(deck)
	deck.Add(second_card, third_card)
	fmt.Println(deck)
}
