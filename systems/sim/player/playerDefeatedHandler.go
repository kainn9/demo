package simPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	"github.com/yohamta/donburi"
)

type PlayerDefeatedHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerDefeatedHandler(scene *coldBrew.Scene) *PlayerDefeatedHandlerSystem {
	return &PlayerDefeatedHandlerSystem{
		scene: scene,
	}
}

func (sys PlayerDefeatedHandlerSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys PlayerDefeatedHandlerSystem) Run(dt float64, npcEntity *donburi.Entry) {
	ticksHandler := sys.scene.Manager.TickHandler
	state := components.PlayerStateComponent.Get(npcEntity)
	body := components.RigidBodyComponent.Get(npcEntity)

	if state.Combat.Health <= 0 && !state.Combat.Defeated {
		state.Combat.Defeated = true
		state.Combat.DefeatedStartTick = ticksHandler.CurrentTick()

	}

	if state.Combat.Defeated {
		state.Transform.BasicHorizontalMovement = false
		body.Vel = tBokiVec.Vec2{X: 0, Y: 0}
	}

	// ticksSinceDefeated := ticksHandler.TicksSinceNTicks(state.Combat.DefeatedStartTick)
	// if ticksSinceDefeated > playerGlobals.PLAYER_DEFEAT_DURATION_TICKS {
	// 	// TODO: HANDLE GAME OVER!
	// }

}
