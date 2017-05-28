package core

import "fmt"

type CardId uint8

const (
	OMNISCIENT CardId = iota + 1
	SPY
	PRIEST
	SHIP_OWNER
	BANKER
	GUARDIAN
	SENATOR
	BROADCASTER
	MERCHANT
	// EVENTS
	BARBARIANS
	LOBBYING
)

var id_to_name = map[CardId]string{
	OMNISCIENT: "The Omniscient",
	SPY: "Spy",
	PRIEST: "Priest",
	SHIP_OWNER: "Ship Owner",
	BANKER: "Banker",
	GUARDIAN: "Guardian",
	SENATOR: "Senator",
	BROADCASTER: "Broadcast",
	MERCHANT: "Merchant",
	BARBARIANS: "Barbarians",
	LOBBYING: "Lobbying",
}


type card struct {
	influence int
	name string
	id CardId
}

func Card(id CardId) *card {
	_influence := -1
	if (IsCharacter(id)) {
		_influence = int(id)
	}
	return &card{
		influence:_influence,
		name: id_to_name[id],
		id: id}
}

func IsEvent(id CardId) bool {
	return id == LOBBYING || id == BARBARIANS;
}

func IsCharacter(id CardId) bool {
	return !IsEvent(id);
}


func (cardPtr *card) String() (s string) {
	return fmt.Sprintf("%v (%v)", cardPtr.name, cardPtr.influence)
}
