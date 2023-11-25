package components

import "github.com/yohamta/donburi"

var AttackDataComponent = donburi.NewComponentType[AttackData]()

type AttackData struct {
	Initiator *donburi.Entry
	Name      CharState
}

func NewAttackState(name string, initiator *donburi.Entry) *AttackData {

	return &AttackData{
		Name:      CharState(name),
		Initiator: initiator,
	}
}
