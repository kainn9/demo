package simPlayerSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type PlayerNpcAttackHitHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerNpcAttackHitHandler(scene *coldBrew.Scene) *PlayerNpcAttackHitHandlerSystem {
	return &PlayerNpcAttackHitHandlerSystem{
		scene: scene,
	}
}

func (PlayerNpcAttackHitHandlerSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys PlayerNpcAttackHitHandlerSystem) Run(dt float64, playerEntity *donburi.Entry) {
	world := sys.scene.World
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	if playerState.Combat.Defeated {
		return
	}

	queries.AttackQuery.Each(world, func(attackEntity *donburi.Entry) {

		attackHitboxes := components.AttackBoxesComponent.Get(attackEntity)
		attackState := components.AttackStateComponent.Get(attackEntity)

		for _, attackHitbox := range *attackHitboxes {

			if isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, attackHitbox, true); isColliding {
				sys.handleHit(*playerState, playerBody, attackHitbox, attackState)
			}
		}
	})

}

func (sys PlayerNpcAttackHitHandlerSystem) handleHit(playerState components.PlayerState, playerBody, attackHitbox *tBokiComponents.RigidBody, attackState *components.AttackState) {

	if attackState.PlayerAttack {
		return
	}

	id := attackState.ID
	atkName := attackState.Name

	if playerState.Combat.Hits[id] != 0 {
		return
	}

	playerState.Combat.Health -= 1
	playerState.Combat.Hits[id] = id
	playerState.Combat.Hit = true

	playerState.Combat.LastHitTick = sys.scene.Manager.TickHandler.CurrentTick()

	sys.handleDisplacement(playerBody, attackHitbox)
	log.Println("playerHit hit! id:", id, "name:", atkName)
	log.Println("health:", playerState.Combat.Health)
}

func (sys PlayerNpcAttackHitHandlerSystem) handleDisplacement(playerBody, npcBody *tBokiComponents.RigidBody) {
	var xFactor = 1.0

	if playerBody.Pos.X < npcBody.Pos.X {
		xFactor = -1
	}

	playerBody.Vel.X = 0
	playerBody.Vel.Y = 0
	tBokiPhysics.Transformer.ApplyImpulseLinear(playerBody, tBokiVec.Vec2{X: 140 * xFactor, Y: -70})

}
