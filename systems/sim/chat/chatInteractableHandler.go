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

	if systemsUtil.IsChatActive(world) {
		return
	}

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	sys.ChatInteractableQuery().Each(world, func(chatEntity *donburi.Entry) {
		chatBody := components.RigidBodyComponent.Get(chatEntity)
		chatState := components.ChatStateAndConfigComponent.Get(chatEntity).State

		if isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, chatBody, true); isColliding {
			sys.handleInteraction(playerState, chatState)
		}

	})

}

func (sys ChatInteractableHandlerSystem) handleInteraction(
	playerState *components.PlayerState,
	chatState *components.ChatState,
) {

	if !playerState.IsInteracting {
		return
	}

	playerState.IsInteracting = false
	chatState.Active = true
	chatState.PopUpMode = true

}
