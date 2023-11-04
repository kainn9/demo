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
	JumpWindupStart         int // Tick.
	Jumping                 bool
	BasicHorizontalMovement bool
	Up                      bool
	Down                    bool
	PhaseThroughPlatforms   bool
}

type PlayerState struct {
	Collision     *PlayerCollisionState
	Transform     *PlayerTransformState
	Animation     AnimState
	IsInteracting bool
}

var PlayerStateComponent = donburi.NewComponentType[PlayerState]()

func NewPlayerState() *PlayerState {

	return &PlayerState{
		Collision: &PlayerCollisionState{},
		Transform: &PlayerTransformState{
			direction: 1,
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
