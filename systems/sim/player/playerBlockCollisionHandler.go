package simPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type PlayerBlockCollisionHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerBlockCollisionHandler(scene *coldBrew.Scene) *PlayerBlockCollisionHandlerSystem {
	return &PlayerBlockCollisionHandlerSystem{
		scene: scene,
	}
}

func (sys PlayerBlockCollisionHandlerSystem) Query() *donburi.Query {
	return queries.BlockQuery
}

func (sys PlayerBlockCollisionHandlerSystem) Run(dt float64, blockEntity *donburi.Entry) {

	world := sys.scene.World

	playerEntity := systemsUtil.PlayerEntity(world)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	blockBody := components.RigidBodyComponent.Get(blockEntity)

	if playerState.Collision.Climbing {
		return
	}

	if isColliding, contacts := tBokiPhysics.Detector.Detect(playerBody, blockBody, true); isColliding {

		p1Left := tBokiVec.Vec2{
			X: playerBody.Polygon.WorldVertices[3].X,
			Y: playerBody.Polygon.WorldVertices[3].Y - 1,
		}

		p1Right := tBokiVec.Vec2{
			X: playerBody.Polygon.WorldVertices[2].X,
			Y: playerBody.Polygon.WorldVertices[2].Y - 1,
		}

		isIntersectingLeft, _ := tBokiPhysics.Resolver.LineIntersection(p1Left, playerBody.Polygon.WorldVertices[3], blockBody.Polygon.WorldVertices[0], blockBody.Polygon.WorldVertices[1])
		isIntersectingRight, _ := tBokiPhysics.Resolver.LineIntersection(p1Right, playerBody.Polygon.WorldVertices[2], blockBody.Polygon.WorldVertices[0], blockBody.Polygon.WorldVertices[1])

		// Note: This is not the same as:
		// playerState.Collision.OnGround = isIntersectingLeft || isIntersectingRight
		// as we don't want to set it to false if it's already true from a previous block.
		if isIntersectingLeft || isIntersectingRight {
			playerState.Collision.OnGround = true
		}

		tBokiPhysics.Resolver.Resolve(playerBody, blockBody, contacts[0])

	}

}
