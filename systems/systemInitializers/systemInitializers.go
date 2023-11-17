package systemInitializers

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	sharedAnimationGlobals "github.com/kainn9/demo/globalConfig/sharedAnimation"
	clientSystems "github.com/kainn9/demo/systems/client"
	clientDebugSystems "github.com/kainn9/demo/systems/client/debug"
	loaderSystems "github.com/kainn9/demo/systems/loader"
	loaderNpcSystems "github.com/kainn9/demo/systems/loader/npc"
	loaderPlayerSystems "github.com/kainn9/demo/systems/loader/player"
	renderSystems "github.com/kainn9/demo/systems/render"
	renderDebugSystems "github.com/kainn9/demo/systems/render/debug"
	renderNpcSystems "github.com/kainn9/demo/systems/render/npc"
	renderPlayerSystems "github.com/kainn9/demo/systems/render/player"
	simSystems "github.com/kainn9/demo/systems/sim"
	simChatSystems "github.com/kainn9/demo/systems/sim/chat"
	simNpcSystems "github.com/kainn9/demo/systems/sim/npc"
	simPlayerSystems "github.com/kainn9/demo/systems/sim/player"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

func InitStandardSystems(scene *coldBrew.Scene, title string, indoor bool) {
	// Loader Systems.
	scene.AddSystem(loaderPlayerSystems.NewPlayerAssetsLoader(scene))
	scene.AddSystem(loaderNpcSystems.NewNpcAssetLoader(scene))
	scene.AddSystem(loaderSystems.NewBackgroundLoader(scene))
	scene.AddSystem(loaderSystems.NewChatLoader(scene))
	scene.AddSystem(loaderSystems.NewUIGlobalLoader(scene))

	// Client Systems.
	scene.AddSystem(clientSystems.NewInputTracker(scene))
	scene.AddSystem(clientDebugSystems.NewDebugClickCoordsTracker(scene))

	// Sim Systems.
	scene.AddSystem(simPlayerSystems.NewPlayerPhysicsInputProcessor(scene))

	scene.AddSystem(simChatSystems.NewChatCollisionHandler(scene))
	scene.AddSystem(simChatSystems.NewChatHandler(scene))
	scene.AddSystem(simChatSystems.NewChatInteractableHandler(scene))

	scene.AddSystem(simPlayerSystems.NewPlayerMovementHandler(scene, indoor))

	scene.AddSystem(simSystems.NewGravityAndIntegrationHandler(scene))
	scene.AddSystem(simPlayerSystems.NewClearOnGroundHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerBlockCollisionHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerPlatformCollisionHandler(scene))
	scene.AddSystem(simNpcSystems.NewNpcGroundCollisionHandler(scene))
	scene.AddSystem(simPlayerSystems.NewLadderHandler(scene))
	scene.AddSystem(simSystems.NewCameraPositionHandler(scene))
	scene.AddSystem(simPlayerSystems.NewIndicatorCollisionHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerMeleeAttackHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerNpcHitHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerIframeHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerDefeatedHandler(scene))
	scene.AddSystem(simNpcSystems.NewNpcAttackedHandler(scene))
	scene.AddSystem(simNpcSystems.NewNpcDefeatedHandler(scene))
	scene.AddSystem(simNpcSystems.NewNpcSimpleAiHandler(scene))

	scene.AddSystem(simSystems.NewSceneTransitionHandler(scene)) // Keep last.

	// Render Systems.
	scene.AddSystem(renderSystems.NewParallaxBackgroundRenderer(scene))
	scene.AddSystem(renderNpcSystems.NewNpcRenderer(scene))
	scene.AddSystem(renderPlayerSystems.NewPlayerRenderer(scene, indoor))
	scene.AddSystem(renderSystems.NewFrontLayerRenderer(scene))
	scene.AddSystem(renderSystems.NewIndicatorRenderer(scene))
	scene.AddSystem(renderSystems.NewCameraRenderer(scene))

	// Post camera render systems.
	scene.AddSystem(renderSystems.NewChatSlidesRenderer(scene))
	scene.AddSystem(renderSystems.NewSceneTitleRenderer(scene, title))
	scene.AddSystem(renderDebugSystems.NewDebugRigidBodyRenderer(scene))
	scene.AddSystem(renderDebugSystems.NewHitBoxPreviewer(scene))

}

func AttachChatCallback(scene *coldBrew.Scene, callback simChatSystems.ChatCallBackSystem) {

	for _, sys := range scene.Systems {
		switch sys.(type) {

		case *simChatSystems.ChatHandlerSystem:
			chatSys := sys.(*simChatSystems.ChatHandlerSystem)
			chatSys.CallBackSystems = append(chatSys.CallBackSystems, callback)
		}
	}
}

type SitCallBackStart struct {
	name  string
	index int
}
type SitCallBackEnd struct {
	name  string
	index int
}

func (cb SitCallBackStart) ChatName() string {
	return cb.name
}

func (cb SitCallBackStart) SlideIndex() int {
	return cb.index
}

func (cb SitCallBackStart) Callback(scene *coldBrew.Scene) {
	playerEntity := systemsUtil.GetPlayerEntity(scene.World)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerState.Animation = playerGlobals.PLAYER_CHAR_STATE_SIT
}

func (cb SitCallBackEnd) ChatName() string {
	return cb.name
}

func (cb SitCallBackEnd) SlideIndex() int {
	return cb.index
}

func (cb SitCallBackEnd) Callback(scene *coldBrew.Scene) {
	playerEntity := systemsUtil.GetPlayerEntity(scene.World)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerState.Animation = sharedAnimationGlobals.CHAR_STATE_IDLE
}

func AttachSitCallbackToChat(scene *coldBrew.Scene, chatName string, slideCount int) {
	AttachChatCallback(scene, SitCallBackStart{chatName, 0})
	AttachChatCallback(scene, SitCallBackEnd{chatName, slideCount})
}
