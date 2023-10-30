package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
)

type PlayerClearOnGroundHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerClearOnGroundHandler(scene *coldBrew.Scene) *PlayerClearOnGroundHandlerSystem {
	return &PlayerClearOnGroundHandlerSystem{
		scene: scene,
	}
}

func (sys PlayerClearOnGroundHandlerSystem) Run(dt float64, _ *donburi.Entry) {

	world := sys.scene.World

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	playerState.OnGround = false

}
