package simChatSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type ChatCollisionHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewChatCollisionHandler(scene *coldBrew.Scene) *ChatCollisionHandlerSystem {
	return &ChatCollisionHandlerSystem{
		scene: scene,
	}
}

func (ChatCollisionHandlerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.ChatStateAndConfigComponent),
			filter.Contains(components.RigidBodyComponent),

			filter.Not(
				filter.Contains(components.IndicatorStateAndConfigComponent),
			),
		),
	)
}

func (sys ChatCollisionHandlerSystem) Run(dt float64, chatEntity *donburi.Entry) {
	playerEntity := systemsUtil.PlayerEntity(sys.scene.World)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	chatBody := components.RigidBodyComponent.Get(chatEntity)
	chatStateAndConfig := components.ChatStateAndConfigComponent.Get(chatEntity)

	if isColliding, _ := tBokiPhysics.Detector.Detect(playerBody, chatBody, true); isColliding {

		if chatStateAndConfig.State.HasBeenRead {
			return
		}

		chatStateAndConfig.State.HasBeenRead = true
		chatStateAndConfig.Enable()
	}

}
