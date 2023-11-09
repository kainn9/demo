package simPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
)

// We make this a its own system, because we want to resetOnGround
// to false before the floor and platform collision handlers run.
type ClearOnGroundHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewClearOnGroundHandler(scene *coldBrew.Scene) *ClearOnGroundHandlerSystem {
	return &ClearOnGroundHandlerSystem{
		scene: scene,
	}
}

func (sys ClearOnGroundHandlerSystem) Run(dt float64, _ *donburi.Entry) {

	world := sys.scene.World

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	playerState.Collision.OnGround = false
}
