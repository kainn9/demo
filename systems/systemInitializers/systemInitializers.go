package systemInitializers

import (
	"github.com/kainn9/coldBrew"
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
	simNpcSystems "github.com/kainn9/demo/systems/sim/npc"
	simPlayerSystems "github.com/kainn9/demo/systems/sim/player"
)

func InitStandardSystems(scene *coldBrew.Scene, zoomed bool) {
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
	scene.AddSystem(simSystems.NewChatHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerMovementHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerNpcHitHandler(scene))
	scene.AddSystem(simSystems.NewGravityAndIntegrationHandler(scene))
	scene.AddSystem(simPlayerSystems.NewClearOnGroundHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerFloorCollisionHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerPlatformCollisionHandler(scene))
	scene.AddSystem(simNpcSystems.NewNpcGroundCollisionHandler(scene))
	scene.AddSystem(simPlayerSystems.NewLadderHandler(scene))
	scene.AddSystem(simSystems.NewCameraPositionHandler(scene, zoomed))
	scene.AddSystem(simPlayerSystems.NewIndicatorCollisionHandler(scene))
	scene.AddSystem(simSystems.NewSceneTransitionHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerMeleeAttackHandler(scene))
	scene.AddSystem(simPlayerSystems.NewPlayerDefeatedHandler(scene))
	scene.AddSystem(simNpcSystems.NewNpcAttackedHandler(scene))
	scene.AddSystem(simNpcSystems.NewNpcDefeatedHandler(scene))

	// Render Systems.
	scene.AddSystem(renderSystems.NewParallaxBackgroundRenderer(scene))
	scene.AddSystem(renderNpcSystems.NewNpcRenderer(scene))
	scene.AddSystem(renderPlayerSystems.NewPlayerRenderer(scene))
	scene.AddSystem(renderSystems.NewFrontLayerRenderer(scene))
	scene.AddSystem(renderSystems.NewIndicatorRenderer(scene))
	scene.AddSystem(renderSystems.NewChatSlidesRenderer(scene))

	// This must be the last Render system to be added for now(minus debug).
	scene.AddSystem(renderSystems.NewCameraRenderer(scene))
	scene.AddSystem(renderDebugSystems.NewDebugRigidBodyRenderer(scene))
	scene.AddSystem(renderDebugSystems.NewHitBoxPreviewer(scene))

}
