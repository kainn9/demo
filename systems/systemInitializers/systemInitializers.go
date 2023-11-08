package systemInitializers

import (
	"github.com/kainn9/coldBrew"
	clientSystems "github.com/kainn9/demo/systems/client"
	loaderSystems "github.com/kainn9/demo/systems/loader"
	renderSystems "github.com/kainn9/demo/systems/render"
	renderDebugSystems "github.com/kainn9/demo/systems/render/debug"
	simSystems "github.com/kainn9/demo/systems/sim"
)

func InitStandardSystems(scene *coldBrew.Scene) {
	// Loader Systems.
	scene.AddSystem(loaderSystems.NewPlayerSpritesLoader(scene))
	scene.AddSystem(loaderSystems.NewBackgroundLoader(scene))
	scene.AddSystem(loaderSystems.NewChatLoader(scene))
	scene.AddSystem(loaderSystems.NewUIGlobalLoader(scene))

	// Client Systems.
	scene.AddSystem(clientSystems.NewInputTracker(scene))

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
	scene.AddSystem(simSystems.NewSceneTransitionHandler(scene))
	scene.AddSystem(simSystems.NewPlayerAttackHandler(scene))

	// Render Systems.
	scene.AddSystem(renderSystems.NewParallaxBackgroundRenderer(scene))
	scene.AddSystem(renderSystems.NewPlayerRenderer(scene))
	scene.AddSystem(renderSystems.NewFrontLayerRenderer(scene))
	scene.AddSystem(renderSystems.NewIndicatorRenderer(scene))
	scene.AddSystem(renderSystems.NewChatSlidesRenderer(scene))

	// This must be the last Render system to be added for now(minus debug).
	scene.AddSystem(renderSystems.NewCameraRenderer(scene))
	scene.AddSystem(renderDebugSystems.NewDebugRigidBodyRenderer(scene))
	scene.AddSystem(renderDebugSystems.NewHitBoxPreviewer(scene))

}

func InitDrivingSystems(scene *coldBrew.Scene) {
	// Loader Systems.
	scene.AddSystem(loaderSystems.NewBackgroundLoader(scene))
	scene.AddSystem(loaderSystems.NewUIGlobalLoader(scene))
	scene.AddSystem(loaderSystems.NewPlayerCarLoader(scene))

	// Client Systems.
	// scene.AddSystem(clientSystems.NewInputTracker(scene))
	// Replace^

	// Sim Systems.
	scene.AddSystem(simSystems.NewCameraPositionHandler(scene))
	scene.AddSystem(simSystems.NewSceneTransitionHandler(scene))
	scene.AddSystem(simSystems.NewPlayerCarMovementHandler())

	// Render Systems.
	scene.AddSystem(renderSystems.NewParallaxBackgroundRenderer(scene))
	scene.AddSystem(renderSystems.NewPlayerCarRenderer(scene))
	scene.AddSystem(renderSystems.NewFrontLayerRenderer(scene))

	// This must be the last Render system to be added for now(minus debug).
	scene.AddSystem(renderSystems.NewCameraRenderer(scene))
	scene.AddSystem(renderDebugSystems.NewDebugRigidBodyRenderer(scene))

}
