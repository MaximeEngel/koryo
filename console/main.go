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
}
