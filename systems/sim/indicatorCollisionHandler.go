package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
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
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.IndicatorStateAndConfigComponent),
			filter.Contains(components.RigidBodyComponent),
		),
	)
}

func (sys IndicatorCollisionHandlerSystem) Run(dt float64, _ *donburi.Entry) {

	world := sys.scene.World

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	sys.IndicatorQuery().Each(world, func(indicatorEntity *donburi.Entry) {

		body := components.RigidBodyComponent.Get(indicatorEntity)
		indicatorState := components.IndicatorStateAndConfigComponent.Get(indicatorEntity)

		if isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, body, true); isColliding {

			indicatorState.State.Active = true

		} else {

			indicatorState.State.Active = false

		}
	})

}
