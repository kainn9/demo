package components

import (
	"github.com/yohamta/donburi"
)

type NpcTransformState struct {
	direction               float64
	BasicHorizontalMovement bool
}

type NpcCombatState struct {
	Hits                                   map[int]int
	Health, LastHitTick, DefeatedStartTick int
	Hittable, IsHit, Defeated              bool
}

type NpcState struct {
	Animation CharState
	Transform *NpcTransformState
	Combat    *NpcCombatState
}

var NpcStateComponent = donburi.NewComponentType[NpcState]()

func NewNpcState(hittable bool) *NpcState {

	return &NpcState{
		Transform: &NpcTransformState{
			direction: -1,
		},
		Combat: &NpcCombatState{
			Health:   3,
			Hits:     make(map[int]int),
			Hittable: hittable,
		},
	}

}

func (ns *NpcState) Direction() float64 {
	return ns.Transform.direction
}

func (ns *NpcState) SetDirectionLeft() {
	ns.Transform.direction = -1
}

func (ns *NpcState) SetDirectionRight() {
	ns.Transform.direction = 1
}
