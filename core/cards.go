package core

import (
	"fmt"
	"math/rand"
)

const NB_CHARS = 45
const NB_EVENTS = 10
const NB_CARDS = NB_CHARS + NB_EVENTS

type cards []*card

func Deck() *cards {
	d := make(cards, 0, NB_CARDS)
	for i := 1 ; i < 10 ; i++ {
		for j := 0 ; j < i; j++ {
			d = append(d, Card(CardId(i)))
		}
	}
	return &d
}

func (cardsPtr *cards) Shuffle() {
	c := *cardsPtr
	for i := range c{
		j := rand.Intn(i + 1)
		c[i], c[j] = c[j], c[i]
	}
}

func (cardsPtr *cards) Draw() *card {
	c := *cardsPtr
	drawed_card := c[len(c) - 1]
	c[len(c) -1 ] = nil
	*cardsPtr = c[:len(c) - 1]
	return drawed_card
}

func (cardsPtr *cards) Add(c ...*card) {
	*cardsPtr = append(*cardsPtr, c...)
}

func (cardsPtr *cards) String() (s string) {
	s = ""
	for _, v := range *cardsPtr {
		s += fmt.Sprintf("%v ", v)
	}
	return
}

