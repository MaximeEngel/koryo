package core

import (
	"fmt"
)

const NB_CHARS = 45
const NB_EVENTS = 10
const NB_CARDS = NB_CHARS + NB_EVENTS

type deck struct {
	cards []*card
}

func Deck() *deck {
	d := &deck{make([]*card, 0, NB_CARDS)}
	for i := 1 ; i < 10 ; i++ {
		for j := 0 ; j < i; j++ {
			d.cards = append(d.cards, &card{Influence:i})
		}
	}
	return d
}

func (d *deck) String() (s string) {
	s = ""
	for _, v := range d.cards {
		s += fmt.Sprintf("%v", *v)
	}
	return
}

