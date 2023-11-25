package simChatSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type ChatInteractableHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewChatInteractableHandler(scene *coldBrew.Scene) *ChatInteractableHandlerSystem {
	return &ChatInteractableHandlerSystem{
		scene: scene,
	}
}

func (sys ChatInteractableHandlerSystem) ChatInteractableQuery() *donburi.Query {
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.ChatStateAndConfigComponent),
			filter.Contains(components.RigidBodyComponent),
			filter.Contains(components.IndicatorStateAndConfigComponent),
		),
	)
}

func (sys ChatInteractableHandlerSystem) Run(dt float64, _ *donburi.Entry) {

	world := sys.scene.World

	isChatActive, _ := systemsUtil.IsChatActive(sys.scene.World)
	if isChatActive {
		return
	}

	playerEntity := systemsUtil.PlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	sys.ChatInteractableQuery().Each(world, func(chatEntity *donburi.Entry) {
		chatBody := components.RigidBodyComponent.Get(chatEntity)
		chatStateAndConfig := components.ChatStateAndConfigComponent.Get(chatEntity)

		if isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, chatBody, true); isColliding {
			sys.handleInteraction(playerState, chatStateAndConfig)
		}

	})

}

func (sys ChatInteractableHandlerSystem) handleInteraction(
	playerState *components.PlayerState,
	chatState *components.ChatStateAndConfig,
) {

	if !playerState.IsInteracting {
		return
	}

	playerState.IsInteracting = false
	chatState.Enable()
}
