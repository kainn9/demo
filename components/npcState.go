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
	Hits                                                    map[int]bool
	LatestHitAttackName                                     CharState
	Health, LastHitTick, DefeatedStartTick, AttackStartTick int
	Hittable, IsHit, Defeated                               bool
	AttackRange                                             float64
	PatrolRange                                             float64
	MaxLeft, MaxRight                                       float64
	CurrentAttack                                           CharState
}

func NewNpcState(hittable bool, attackRange, patrolRange, maxLeft, maxRight, speed float64) *NpcState {

	return &NpcState{
		Transform: &NpcTransformState{
			direction: -1,
			Speed:     speed,
		},
		Combat: &NpcCombatState{
			Health:      3,
			Hits:        make(map[int]bool),
			Hittable:    hittable,
			AttackRange: attackRange,
			PatrolRange: patrolRange,
			MaxLeft:     maxLeft,
			MaxRight:    maxRight,
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

func (ns *NpcState) IsLeft() bool {
	return ns.Transform.direction == -1
}

func (ns *NpcState) IsRight() bool {
	return ns.Transform.direction == 1
}
