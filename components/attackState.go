package components

import "github.com/yohamta/donburi"

var AttackStateComponent = donburi.NewComponentType[AttackState]()

type AttackState struct {
	ID           int
	OwnerId      int
	PlayerAttack bool
}

func NewAttackState(id, ownerId int, playerAttack bool) *AttackState {

	return &AttackState{
		ID:           id,
		OwnerId:      ownerId,
		PlayerAttack: playerAttack,
	}
}
