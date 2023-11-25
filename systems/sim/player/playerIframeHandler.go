package simPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	"github.com/kainn9/demo/queries"
	"github.com/yohamta/donburi"
)

type PlayerIframeHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerIframeHandler(scene *coldBrew.Scene) *PlayerIframeHandlerSystem {
	return &PlayerIframeHandlerSystem{
		scene: scene,
	}
}

func (sys PlayerIframeHandlerSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys PlayerIframeHandlerSystem) Run(dt float64, playerEntity *donburi.Entry) {

	playerState := components.PlayerStateComponent.Get(playerEntity)
	tickHandler := sys.scene.Manager.TickHandler

	if sys.iframeStartTickMustBeSet(playerState) {
		playerState.Combat.RecoveryIframeStartTick = tickHandler.CurrentTick()
	}

	if sys.iframeJustEnded(playerState, tickHandler) {
		playerState.Combat.IsInRecoveryIframe = false
		playerState.Combat.RecoveryIframeStartTick = -1
	}

}

func (sys PlayerIframeHandlerSystem) iframeStartTickMustBeSet(playerState *components.PlayerState) bool {

	return playerState.Combat.RecoveryIframeStartTick == -1 && playerState.Combat.IsInRecoveryIframe

}

func (sys PlayerIframeHandlerSystem) iframeJustEnded(playerState *components.PlayerState, tickHandler *coldBrew.TickHandler) bool {

	if playerState.Combat.RecoveryIframeStartTick == -1 {
		return false
	}

	return tickHandler.TicksSinceNTicks(playerState.Combat.RecoveryIframeStartTick) >= playerGlobals.PLAYER_RECOVERY_IFRAME_DURATION_TICKS

}
