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
		// sys.collisionHandler(playerBody, blockBody, contacts)
		tBokiPhysics.Resolver.Resolve(playerBody, blockBody, contacts)
		sys.groundedHandler(playerBody, blockBody, playerState)
	}

}

func (sys *PlayerBlockCollisionHandlerSystem) collisionHandler(playerBody, blockBody *tBokiComponents.RigidBody, contacts tBokiComponents.Contacts) {

	// incidentEdge := contacts.Data[0].IncidentEdge[blockBody]

	// log.Println(incidentEdge)

	playerTopLeftVert := playerBody.Polygon.WorldVertices[0]
	playerBottomLeftVert := playerBody.Polygon.WorldVertices[3]

	playerTopRightVert := playerBody.Polygon.WorldVertices[1]
	playerBottomRightVert := playerBody.Polygon.WorldVertices[2]

	blockTopLeftVert := blockBody.Polygon.WorldVertices[0]
	blockTopRightVert := blockBody.Polygon.WorldVertices[1]

	contactVertTop := playerTopLeftVert
	contactVertBottom := playerBottomLeftVert

	if blockTopLeftVert.Y < blockTopRightVert.Y {
		contactVertTop = playerTopRightVert
		contactVertBottom = playerBottomRightVert
	}

	isIntersectingSide, sideIntersection := tBokiPhysics.Resolver.LineIntersection(contactVertTop, contactVertBottom, blockTopLeftVert, blockTopRightVert)
	if !isIntersectingSide {
		return
	}

	playerBody.Pos.Y = sideIntersection.Y - playerBody.Polygon.Height/2 + 0.01
	playerBody.Vel.Y = 0

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
