package core

import "fmt"

type CardId uint8

const (
	OMNISCIENT CardId = iota + 1
	SPY
	SENATOR
	PRIEST
	SHIP_OWNER
	BANKER
	GUARDIAN
	BROADCASTER
	MERCHANT
	// EVENTS
	BARBARIANS
	LOBBYING
)

var id_to_name = map[CardId]string{
	OMNISCIENT:  "The Omniscient",
	SPY:         "Spy",
	PRIEST:      "Priest",
	SHIP_OWNER:  "Ship Owner",
	BANKER:      "Banker",
	GUARDIAN:    "Guardian",
	SENATOR:     "Senator",
	BROADCASTER: "Broadcaster",
	MERCHANT:    "Merchant",
	BARBARIANS:  "Barbarians",
	LOBBYING:    "Lobbying",
}

type card struct {
	influence int
	name      string
	id        CardId
}

func Card(id CardId) *card {
	return &card{
		influence: CardInfluence(id),
		name:      id_to_name[id],
		id:        id}
}

func CardInfluence(id CardId) int {
	if IsCharacter(id) {
		return int(id)
	}
	// all events
	return -1
}

func CardIdsIterator() func()(CardId, bool) {
	id := 0
	count := len(id_to_name)
	return func() (CardId, bool) {
		id++
		if id >= count {
			return CardId(id), false
		}
		return CardId(id), true

	}
}

func (cardPtr *card) Id() CardId {
	return cardPtr.id
}

func IsEvent(id CardId) bool {
	return id == LOBBYING || id == BARBARIANS
}

func IsCharacter(id CardId) bool {
	return !IsEvent(id)
}

func (cardPtr *card) String() (s string) {
	return fmt.Sprintf("%v (%v)", cardPtr.name, cardPtr.influence)
}

func IdToName(id CardId) string {
	name, ok := id_to_name[id]
	if ok {
		return name
	}
	return "Unknown"
}
