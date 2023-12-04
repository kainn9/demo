package simPlayerSystems

import (
	"math"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
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

	if isColliding, contacts := tBokiPhysics.Detector.Detect(playerBody, blockBody, tBokiComponents.ResolverType); isColliding {
		sys.verticalCollisionHandler(playerBody, blockBody, contacts)
		sys.groundedHandler(playerBody, blockBody, playerState)
	}

}

func (PlayerBlockCollisionHandlerSystem) verticalCollisionHandler(playerBody, blockBody *tBokiComponents.RigidBody, contacts tBokiComponents.Contacts) {

	pBottomLeftVert := playerBody.Polygon.WorldVertices[2]
	pBottomRightVert := playerBody.Polygon.WorldVertices[3]

	bTopLeftVert := blockBody.Polygon.WorldVertices[0]
	bTopRightVert := blockBody.Polygon.WorldVertices[1]

	pBottomIntersectionBTop, _ := tBokiPhysics.Resolver.LineIntersection(pBottomLeftVert, pBottomRightVert, bTopLeftVert, bTopRightVert)

	pBottomBelowBTopWithNoRotation := blockBody.Rotation == 0 && pBottomLeftVert.Y <= bTopLeftVert.Y

	if pBottomIntersectionBTop || pBottomBelowBTopWithNoRotation {
		ox := playerBody.Pos.X
		ovx := playerBody.Vel.X

		tBokiPhysics.Resolver.Resolve(playerBody, blockBody, contacts)

		playerBody.Pos.X = ox
		playerBody.Vel.X = ovx

		if playerBody.Vel.X > 0 {
			playerBody.Vel.Y = 0
		}
	} else {
		tBokiPhysics.Resolver.Resolve(playerBody, blockBody, contacts)
	}
}

func (PlayerBlockCollisionHandlerSystem) groundedHandler(playerBody, blockBody *tBokiComponents.RigidBody, playerState *components.PlayerState) {
	pBottomLeftVert := playerBody.Polygon.WorldVertices[2]
	pBottomRightVert := playerBody.Polygon.WorldVertices[3]
	pLowestY := math.Min(pBottomLeftVert.Y, pBottomRightVert.Y)

	bTopLeftVert := blockBody.Polygon.WorldVertices[0]
	bTopRightVert := blockBody.Polygon.WorldVertices[1]

	bHighestY := math.Max(bTopLeftVert.Y, bTopRightVert.Y)

	if pLowestY <= bHighestY {
		playerState.Collision.OnGround = true
	}

}
