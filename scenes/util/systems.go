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
	scene.AddSystem(loaderSystems.NewParallaxBackgroundLoader(scene))
	scene.AddSystem(loaderSystems.NewChatLoader(scene))

	// Client Systems.
	scene.AddSystem(clientSystems.NewPlayerMovementInputTracker())

	// Sim Systems.
	scene.AddSystem(simSystems.NewPlayerInputHandler(scene))
	scene.AddSystem(simSystems.NewChatHandler(scene))
	scene.AddSystem(simSystems.NewPlayerMovementHandler(scene))
	scene.AddSystem(simSystems.NewPlayerBlockCollisionHandler(scene))
	scene.AddSystem(simSystems.NewCameraPositionHandler(scene))

	// Render Systems.
	scene.AddSystem(renderSystems.NewParallaxBackgroundRenderer(scene))
	scene.AddSystem(renderSystems.NewChatSlidesRenderer(scene))
	scene.AddSystem(renderSystems.NewPlayerRenderer(scene))
	scene.AddSystem(renderSystems.NewDebugRigidBodyRenderer(scene))
}
