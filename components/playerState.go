package components

import (
	"github.com/yohamta/donburi"
)

type PlayerState struct {
	direction               float64
	JumpWindupStart         int // tick when jump windup started.
	BasicHorizontalMovement bool
	OnGround                bool
	Jumping                 bool
	AnimationState          string
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
