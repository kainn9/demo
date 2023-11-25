package components

import (
	"github.com/yohamta/donburi"
)

var PlayerStateComponent = donburi.NewComponentType[PlayerState]()

type PlayerState struct {
	Collision     *PlayerCollisionState
	Transform     *PlayerTransformState
	Combat        *PlayerCombatState
	Animation     CharState
	IsInteracting bool
}

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

	DodgeTriggered    bool
	Dodging           bool
	DodgeFinishedTick int
}

type PlayerCombatState struct {
	Attacking, IsHit, Defeated, IsInRecoveryIframe bool
	CurrentAttack                                  CharState
	Hits                                           map[int]bool

	Health,
	AttackStartTick,
	LastHitTick,
	RecoveryIframeStartTick,
	DefeatedStartTick int
}

func NewPlayerState() *PlayerState {

	return &PlayerState{
		Collision: &PlayerCollisionState{},
		Transform: &PlayerTransformState{
			direction: 1,
		},

		Combat: &PlayerCombatState{
			AttackStartTick:         -1,
			LastHitTick:             -1,
			RecoveryIframeStartTick: -1,
			DefeatedStartTick:       -1,
			Health:                  10,
			Hits:                    make(map[int]bool),
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

func (cs *PlayerCombatState) ClearAttackState() {
	cs.Attacking = false
	cs.AttackStartTick = -1
	cs.CurrentAttack = ""
}
