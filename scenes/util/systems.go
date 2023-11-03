package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	clientSystems "github.com/kainn9/demo/systems/client"
	loaderSystems "github.com/kainn9/demo/systems/loader"
	renderSystems "github.com/kainn9/demo/systems/render"
	simSystems "github.com/kainn9/demo/systems/sim"
)

func InitStandardSystems(scene *coldBrew.Scene) {
	// Loader Systems.
	scene.AddSystem(loaderSystems.NewPlayerSpritesLoader(scene))
	scene.AddSystem(loaderSystems.NewBackgroundLoader(scene))
	scene.AddSystem(loaderSystems.NewChatLoader(scene))
	scene.AddSystem(loaderSystems.NewUIGlobalLoader(scene))

	// Client Systems.
	scene.AddSystem(clientSystems.NewPlayerMovementInputTracker(scene))

	// Sim Systems.
	scene.AddSystem(simSystems.NewPlayerMovementInputProcessor(scene))
	scene.AddSystem(simSystems.NewChatHandler(scene))
	scene.AddSystem(simSystems.NewPlayerMovementHandler(scene))
	scene.AddSystem(simSystems.NewClearOnGroundHandler(scene))
	scene.AddSystem(simSystems.NewPlayerFloorCollisionHandler(scene))
	scene.AddSystem(simSystems.NewPlayerPlatformCollisionHandler(scene))
	scene.AddSystem(simSystems.NewLadderHandler(scene))
	scene.AddSystem(simSystems.NewCameraPositionHandler(scene))
	scene.AddSystem(simSystems.NewIndicatorCollisionHandler(scene))

	// Render Systems.
	scene.AddSystem(renderSystems.NewParallaxBackgroundRenderer(scene))
	scene.AddSystem(renderSystems.NewPlayerRenderer(scene))
	scene.AddSystem(renderSystems.NewFrontLayerRenderer(scene))
	scene.AddSystem(renderSystems.NewIndicatorRenderer(scene))
	scene.AddSystem(renderSystems.NewChatSlidesRenderer(scene))

	// This must be the last Render system to be added for now(minus debug).
	scene.AddSystem(renderSystems.NewCameraRenderer(scene))
	scene.AddSystem(renderSystems.NewDebugRigidBodyRenderer(scene))

}
