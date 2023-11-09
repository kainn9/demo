package simPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type PlayerPlatformCollisionHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerPlatformCollisionHandler(scene *coldBrew.Scene) *PlayerPlatformCollisionHandlerSystem {
	return &PlayerPlatformCollisionHandlerSystem{
		scene: scene,
	}
}

func (sys PlayerPlatformCollisionHandlerSystem) CustomQuery() *donburi.Query {
	return queries.PlatformQuery
}

func (sys PlayerPlatformCollisionHandlerSystem) Run(dt float64, _ *donburi.Entry) {

	world := sys.scene.World
	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	if playerState.Collision.Climbing {
		return
	}

	query := sys.CustomQuery()

	noPlatformCollisions := true

	query.Each(world, func(platformEntity *donburi.Entry) {

		platform := components.RigidBodyComponent.Get(platformEntity)

		if (playerBody.Pos.Y+playerBody.Polygon.Height/2 > platform.Pos.Y+platform.Polygon.Height/2) || playerBody.Vel.Y < 0 {
			return
		}

		if isColliding, contacts := tBokiPhysics.Detector.Detect(playerBody, platform, true); isColliding {
			sys.handleCollision(playerBody, platform, playerState, contacts)
			noPlatformCollisions = false
		}

	})

	if noPlatformCollisions {
		playerState.Transform.PhaseThroughPlatforms = false
	}

}

func (PlayerPlatformCollisionHandlerSystem) handleCollision(
	playerBody *tBokiComponents.RigidBody,
	platform *tBokiComponents.RigidBody,
	playerState *components.PlayerState,
	contacts []tBokiComponents.Contact,
) {

	if playerState.Transform.PhaseThroughPlatforms {
		return
	}

	tBokiPhysics.Resolver.Resolve(playerBody, platform, contacts[0])
	playerState.Collision.OnGround = true
}
