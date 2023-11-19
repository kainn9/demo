package components

import "github.com/yohamta/donburi"

var AttackStateComponent = donburi.NewComponentType[AttackState]()

type AttackState struct {
	Name         string
	ID           int
	OwnerId      int
	PlayerAttack bool
}

func NewAttackState(id, ownerId int, name string, playerAttack bool) *AttackState {

	return &AttackState{
		ID:           id,
		OwnerId:      ownerId,
		PlayerAttack: playerAttack,
		Name:         name,
	}
}
