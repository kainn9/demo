package components

import (
	"github.com/yohamta/donburi"
)

type PlayerState struct {
	direction               float64
	JumpWindupStart         int // Tick.
	BasicHorizontalMovement bool
	Up                      bool
	Down                    bool
	Interact                bool
	OnGround                bool
	Climbing                bool
	Jumping                 bool
	PhaseThroughPlatforms   bool

	AnimationState string
}

var PlayerStateComponent = donburi.NewComponentType[PlayerState]()

func NewPlayerState() *PlayerState {

	return &PlayerState{

		direction: 1,
	}
}

func (ps *PlayerState) Direction() float64 {
	return ps.direction
}

func (ps *PlayerState) SetDirectionLeft() {
	ps.direction = -1
}

func (ps *PlayerState) SetDirectionRight() {
	ps.direction = 1
}
