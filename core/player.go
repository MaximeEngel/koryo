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
	selected_to_play cards
	victoryPoints int
	gb *gameBoard
}

func Player(name string, gb *gameBoard) *player {
	return &player{
		Name:name,
		hand:make(cards, 0, MAX_HAND_CARDS),
		played:make(cards, 0, 14), // 14 seems to be a good average
		selected_to_play:make(cards, 0, 9), // 9 merchant at max
		victoryPoints:0,
		gb:gb}
}

func (p *player) Draw(c *card) {
	p.hand = append(p.hand, c)
}

func (p* player) PlayCard(idx uint) bool {
	if idx >= uint(len(p.selected_to_play)) {
		return false
	}

	// Add
	p.played = append(p.played, p.selected_to_play[idx])

	// Delete from hand
	copy(p.selected_to_play[idx:], p.selected_to_play[idx+1:])
	last := len(p.selected_to_play)-1
	p.selected_to_play[last] = nil
	p.selected_to_play = p.selected_to_play[:last]
	return true
}

func (p* player) CanSelectPlayCardPtr(c *card) (uint, bool) {
	for idx, other := range p.hand {
		if other == c {
			u_idx := uint(idx)
			return u_idx, p.CanSelectPlayCard(u_idx)
		}
	}
	return 0, false
}

func (p *player) CanSelectPlayCard(hand_idx uint) bool {
	nb_selected := len(p.selected_to_play)
	if nb_selected == 0 {
		return true
	}

	two_different := p.gb.HasMajority(p, SHIP_OWNER, false)
	if nb_selected == 1 && two_different {
		return true
	}

	c := p.hand[hand_idx]
	selected_ids := make(map[CardId]bool)
	for _, selected := range p.selected_to_play {
		selected_ids[selected.Id()] = true
	}
	if len(selected_ids) == 1 && selected_ids[c.Id()] {
		return true
	}

	return false
}

func (p *player) SelectPlayCardPtr(c *card) bool{
	idx, found := p.CanSelectPlayCardPtr(c)
	if found {
		return p.SelectPlayCard(idx)
	}
	return false
}

func (p *player) SelectPlayCard(hand_idx uint) bool {
	if !p.CanSelectPlayCard(hand_idx) {
		return false
	}

	c := p.hand[hand_idx]
	p.hand = append(p.hand[:hand_idx], p.hand[hand_idx + 1:]...)
	p.selected_to_play = append(p.selected_to_play, c)
	return true
}

// Only for reading
func (p *player) HandConst() cards{
	return p.hand
}

func (p *player) NbHand() int{
	return len(p.hand)
}

// Only for reading
func (p *player) PlayedConst() cards {
	return p.played
}

func (p *player) HasPlayed(id CardId) int {
	count := 0
	for _, card := range p.played {
		if card.id == id {
			count++
		}
	}
	return count
}

func (p *player) NbPlayed() int {
	return len(p.played)
}

func (p *player) NbSelected() int {
	return len(p.selected_to_play)
}

func (p *player) String() (s string) {
	s = p.Name
	s += "\n Hand : \n"
	for _, c := range p.hand {
		s += fmt.Sprintf(" %v ", c)
	}
	s += "\n Played : \n"
	for _, c := range p.played {
		s += fmt.Sprintf("%s ", c)
	}
	s += fmt.Sprintf("\n Victory Points : %v", p.victoryPoints)
	return
}
