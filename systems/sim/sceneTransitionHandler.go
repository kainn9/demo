package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	scenesUtil "github.com/kainn9/demo/scenes/util"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type SceneTransitionHandlerSystem struct {
	scene               *coldBrew.Scene
	PermissionCallbacks []SceneTransitionPermissionCallback
}

type SceneTransitionPermissionCallback interface {
	Index() string
	AllowedToTransition(*coldBrew.Scene) bool
	ChatEntity() *donburi.Entry
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

	chatIsActive, _ := systemsUtil.IsChatActive(sys.scene.World)
	if chatIsActive {
		return
	}

	playerEntity := systemsUtil.PlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	sys.TransitionQuery().Each(world, func(transitionEntity *donburi.Entry) {

		transitionBody := components.RigidBodyComponent.Get(transitionEntity)

		if isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, transitionBody, tBokiComponents.ResolverType); isColliding {
			sys.handleTransition(transitionEntity, playerState)
		}

	})

}

func (sys SceneTransitionHandlerSystem) handleTransition(transitionEntity *donburi.Entry, playerState *components.PlayerState) {

	m := sys.scene.Manager
	transitionStateAndConfig := components.SceneTransitionStateAndConfigComponent.Get(transitionEntity)
	newScene := transitionStateAndConfig.Config.TargetScene

	if !playerState.IsInteracting {
		return
	}

	if !sys.allowedToTransition(transitionStateAndConfig) {
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

func (sys SceneTransitionHandlerSystem) allowedToTransition(transitionStateAndConfig *components.SceneTransitionStateAndConfig) bool {

	for _, callback := range sys.PermissionCallbacks {
		index := callback.Index()

		if index == transitionStateAndConfig.Config.TargetScene.Index() && !callback.AllowedToTransition(sys.scene) {
			sys.showRestrictedMessage(callback)
			return false
		}

	}

	return true
}

func (sys SceneTransitionHandlerSystem) showRestrictedMessage(callback SceneTransitionPermissionCallback) {
	chatEntity := callback.ChatEntity()
	chatStateAndConfig := components.ChatStateAndConfigComponent.Get(chatEntity)
	chatStateAndConfig.Enable()
}
