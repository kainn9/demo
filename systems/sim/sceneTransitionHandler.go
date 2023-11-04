package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	scenesUtil "github.com/kainn9/demo/scenes/util"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type SceneTransitionHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewSceneTransitionHandler(scene *coldBrew.Scene) *SceneTransitionHandlerSystem {
	return &SceneTransitionHandlerSystem{
		scene: scene,
	}
}

func (sys SceneTransitionHandlerSystem) TransitionQuery() *donburi.Query {
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.IndicatorStateAndConfigComponent),
			filter.Contains(components.RigidBodyComponent),
			filter.Contains(components.SceneTransitionStateAndConfigComponent),
		),
	)
}

func (sys SceneTransitionHandlerSystem) Run(dt float64, transitionEntity *donburi.Entry) {

	world := sys.scene.World

	if systemsUtil.IsChatActive(world) {
		return
	}

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	sys.TransitionQuery().Each(world, func(transitionEntity *donburi.Entry) {

		transitionBody := components.RigidBodyComponent.Get(transitionEntity)

		if isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, transitionBody, true); isColliding {
			sys.handleTransition(transitionEntity, playerState)
		}

	})

}

func (sys SceneTransitionHandlerSystem) handleTransition(transitionEntity *donburi.Entry, playerState *components.PlayerState) {

	m := sys.scene.Manager
	transitionStateAndConfig := components.SceneTransitionStateAndConfigComponent.Get(transitionEntity)
	newScene := transitionStateAndConfig.Config.TargetScene

	if !playerState.IsInteracting && transitionStateAndConfig.Config.ClickBased {
		return
	}

	playerState.IsInteracting = false
	scenesUtil.ChangeScene(
		m,
		newScene,
		transitionStateAndConfig.Config.SpawnX,
		transitionStateAndConfig.Config.SpawnY,
		transitionStateAndConfig.Config.CamX,
		transitionStateAndConfig.Config.CamY,
	)

}
