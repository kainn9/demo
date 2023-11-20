package components

import (
	"github.com/yohamta/donburi"
)

var NpcStateComponent = donburi.NewComponentType[NpcState]()

type NpcState struct {
	Animation CharState
	Transform *NpcTransformState
	Combat    *NpcCombatState
}
type NpcTransformState struct {
	direction               float64
	BasicHorizontalMovement bool
	Speed                   float64
}

type NpcCombatState struct {
	Hits                                                    map[int]int
	LatestHitAttackName                                     string
	Health, LastHitTick, DefeatedStartTick, AttackStartTick int
	Hittable, IsHit, Defeated                               bool
	AttackRange                                             float64
	PatrolRange                                             float64
	CurrentAttack                                           CharState
}

func NewNpcState(hittable bool, attackRange, patrolRange, speed float64) *NpcState {

	return &NpcState{
		Transform: &NpcTransformState{
			direction: -1,
			Speed:     speed,
		},
		Combat: &NpcCombatState{
			Health:      3,
			Hits:        make(map[int]int),
			Hittable:    hittable,
			AttackRange: attackRange,
			PatrolRange: patrolRange,
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
