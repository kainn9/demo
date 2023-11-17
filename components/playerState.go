package components

import (
	"github.com/yohamta/donburi"
)

type PlayerCollisionState struct {
	OnGround bool
	Climbing bool
}
type PlayerTransformState struct {
	direction               float64
	JumpTriggered           bool
	Jumping                 bool
	BasicHorizontalMovement bool
	Up                      bool
	Down                    bool
	PhaseThroughPlatforms   bool
}

type PlayerCombatState struct {
	Attacking, Hit, Defeated, Invincible bool
	CurrentAttack                        CharState
	Hits                                 map[int]int

	Health,
	AttackStartTick,
	LastHitTick,
	InvincibleStartTick,
	DefeatedStartTick int
}

type PlayerState struct {
	Collision     *PlayerCollisionState
	Transform     *PlayerTransformState
	Combat        *PlayerCombatState
	Animation     CharState
	IsInteracting bool
}

var PlayerStateComponent = donburi.NewComponentType[PlayerState]()

func NewPlayerState() *PlayerState {

	return &PlayerState{
		Collision: &PlayerCollisionState{},
		Transform: &PlayerTransformState{
			direction: 1,
		},
		Combat: &PlayerCombatState{
			AttackStartTick:     -1,
			LastHitTick:         -1,
			InvincibleStartTick: 1,
			Health:              10,
			Hits:                make(map[int]int),
		},
	}
}

func (ps *PlayerState) Direction() float64 {
	return ps.Transform.direction
}

func (ps *PlayerState) SetDirectionLeft() {
	ps.Transform.direction = -1
}

func (ps *PlayerState) SetDirectionRight() {
	ps.Transform.direction = 1
}
