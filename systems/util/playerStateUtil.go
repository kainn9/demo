package systemsUtil

import (
	"github.com/kainn9/demo/components"
)

type playerStateHelper struct{}

var PlayerStateHelper = playerStateHelper{}

// These will be more robust in the future.

func (playerStateHelper) PlayerCannotAcceptInputs(state *components.PlayerState) bool {
	return false
}

func (playerStateHelper) PlayerCanJump(state *components.PlayerState) bool {
	return state.OnGround
}
