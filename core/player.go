package core

import (
	"fmt"
)

// Not a gameplay limit but just a good hint for slice capacity indeed
// by logic we have at maximum 10 cards at season 1
// 9 or less in next seasons but with a potential bonus of one for the broadaster
const MAX_HAND_CARDS = 10

type player struct {
	Name string
	hand cards
	played cards
	victoryPoints int
}

func Player(name string) *player {
	return &player{
		Name:name,
		hand:make(cards, 0, MAX_HAND_CARDS),
		played:make(cards, 0, 14),
		victoryPoints:0}
}

func (p *player) Draw(c *card) {
	p.hand = append(p.hand, c)
}

func (p *player) String() (s string) {
	s = p.Name
	s += "\n Hand : \n"
	for _, c := range p.hand {
		s += fmt.Sprintf(" %v", *c)
	}
	s += "\n Played : \n"
	for _, c := range p.played {
		s += fmt.Sprintf("% v", *c)
	}
	s += fmt.Sprintf("\n Victory Points : %v", p.victoryPoints)
	return
}