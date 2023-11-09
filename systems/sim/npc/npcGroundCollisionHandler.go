package simNpcSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type NpcGroundCollisionHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewNpcGroundCollisionHandler(scene *coldBrew.Scene) *NpcGroundCollisionHandlerSystem {
	return &NpcGroundCollisionHandlerSystem{
		scene: scene,
	}
}

func (sys NpcGroundCollisionHandlerSystem) Query() *donburi.Query {
	return queries.NpcQuery
}

func (sys NpcGroundCollisionHandlerSystem) Run(dt float64, npcEntity *donburi.Entry) {

	world := sys.scene.World
	npcBody := components.RigidBodyComponent.Get(npcEntity)

	// TEMP HACK!
	npcBody.Vel.X = 0

	queries.FloorQuery.Each(world, func(floorEntity *donburi.Entry) {
		floorBody := components.RigidBodyComponent.Get(floorEntity)

		if isColliding, contacts := tBokiPhysics.Detector.Detect(npcBody, floorBody, true); isColliding {
			tBokiPhysics.Resolver.Resolve(npcBody, floorBody, contacts[0])
		}
	})

	queries.PlatformQuery.Each(world, func(platformEntity *donburi.Entry) {
		platformBody := components.RigidBodyComponent.Get(platformEntity)

		if isColliding, contacts := tBokiPhysics.Detector.Detect(npcBody, platformBody, true); isColliding {
			tBokiPhysics.Resolver.Resolve(npcBody, platformBody, contacts[0])
		}

	})

}
