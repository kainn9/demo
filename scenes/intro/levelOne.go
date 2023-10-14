package intro

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
	clientSystems "github.com/kainn9/demo/systems/client"
	loaderSystems "github.com/kainn9/demo/systems/loader"
	renderSystems "github.com/kainn9/demo/systems/render"
	simSystems "github.com/kainn9/demo/systems/sim"
	"github.com/kainn9/demo/tags"
)

type LevelOneScene struct{}

func (LevelOneScene) Index() string {
	return "levelOne"
}

func (LevelOneScene) New(m *coldBrew.Manager) *coldBrew.Scene {

	scene := coldBrew.NewScene(m, 640, 1618)

	// Systems ----------------------------------------------------------------------------------
	scene.AddSystem(loaderSystems.NewParallaxBackgroundLoader(scene))
	scene.AddSystem(clientSystems.NewInputTracker())

	scene.AddSystem(simSystems.NewPlayerMovementHandler())

	scene.AddSystem(renderSystems.NewCameraPositionHandler(scene))
	scene.AddSystem(renderSystems.NewParallaxBackgroundRenderer(scene))
	scene.AddSystem(renderSystems.NewPlayerRenderer(scene))

	// Entities ----------------------------------------------------------------------------------
	AddCameraEntity(scene)

	parallaxBackgroundSubPath := "intro/levelOne/"
	AddParallaxBackground(scene, []*assetComponents.ParallaxLayerConfig{
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 0, 0, 0, false),
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 1, 10, 0, true),

		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 2, 15, 0, false),
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 3, 10, 0, false),
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 4, 8, 0, false),

		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 5, 0, 0, false),
	})

	AddPlayerEntity(scene)

	return scene
}

func AddParallaxBackground(scene *coldBrew.Scene, layers []*assetComponents.ParallaxLayerConfig) {

	for _, layer := range layers {

		bgLayerEntity := scene.AddEntity(assetComponents.ParallaxLayerConfigComponent, assetComponents.SpriteComponent)

		assetComponents.ParallaxLayerConfigComponent.SetValue(
			bgLayerEntity,
			*layer,
		)

		assetComponents.SpriteComponent.SetValue(
			bgLayerEntity,
			*assetComponents.NewSprite(),
		)

	}

}

func AddPlayerEntity(scene *coldBrew.Scene) {

	playerEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.InputBufferComponent,
		tags.PlayerTag,
	)

	components.RigidBodyComponent.SetValue(playerEntity, *components.NewRigidBody(0, 0))
	components.InputBufferComponent.SetValue(playerEntity, *components.NewInputBuffer())

}

func AddCameraEntity(scene *coldBrew.Scene) {

	cameraEntity := scene.AddEntity(
		components.CameraComponent,
	)

	components.CameraComponent.SetValue(
		cameraEntity,
		*components.NewCamera(0, 0, constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT),
	)

}
