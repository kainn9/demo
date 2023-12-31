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

type IndicatorCollisionHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewIndicatorCollisionHandler(scene *coldBrew.Scene) *IndicatorCollisionHandlerSystem {
	return &IndicatorCollisionHandlerSystem{
		scene: scene,
	}
}

func (sys IndicatorCollisionHandlerSystem) IndicatorQuery() *donburi.Query {
	return queries.IndicatorQuery
}

func (sys IndicatorCollisionHandlerSystem) Run(dt float64, _ *donburi.Entry) {

	world := sys.scene.World

	playerEntity := systemsUtil.PlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	sys.IndicatorQuery().Each(world, func(indicatorEntity *donburi.Entry) {

		body := components.RigidBodyComponent.Get(indicatorEntity)
		indicatorState := components.IndicatorStateAndConfigComponent.Get(indicatorEntity)

		if isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, body, tBokiComponents.ResolverType); isColliding {
			indicatorState.State.Active = true

		} else {
			indicatorState.State.Active = false
		}
	})

}
