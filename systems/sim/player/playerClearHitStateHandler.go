package simPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

type PlayerClearHitStateHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerClearHitStateHandler(scene *coldBrew.Scene) *PlayerClearHitStateHandlerSystem {
	return &PlayerClearHitStateHandlerSystem{
		scene: scene,
	}
}

func (PlayerClearHitStateHandlerSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys PlayerClearHitStateHandlerSystem) Run(dt float64, playerEntity *donburi.Entry) {

	tickHandler := sys.scene.Manager.TickHandler

	playerState := components.PlayerStateComponent.Get(playerEntity)

	if tickHandler.TicksSinceNTicks(playerState.Combat.LastHitTick) > playerGlobals.PLAYER_HURT_DURATION_TICKS {
		playerState.Combat.IsHit = false
	}

}
