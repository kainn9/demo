package intro

import (
	"strconv"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
	clientSystems "github.com/kainn9/demo/systems/client"
	loaderSystems "github.com/kainn9/demo/systems/loader"
	renderSystems "github.com/kainn9/demo/systems/render"
	simSystems "github.com/kainn9/demo/systems/sim"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

type LevelOneScene struct{}

const (
	SCENE_WIDTH  = 1618
	SCENE_HEIGHT = 360
	SCENE_NAME   = "levelOne"
)

func (LevelOneScene) Index() string {
	return SCENE_NAME
}

func (LevelOneScene) New(m *coldBrew.Manager) *coldBrew.Scene {

	scene := coldBrew.NewScene(m, SCENE_WIDTH, SCENE_HEIGHT)

	// Systems ----------------------------------------------------------------------------------
	scene.AddSystem(loaderSystems.NewPlayerSpritesLoader(scene))
	scene.AddSystem(loaderSystems.NewParallaxBackgroundLoader(scene))

	scene.AddSystem(clientSystems.NewInputTracker())

	scene.AddSystem(simSystems.NewPlayerMovementHandler())
	scene.AddSystem(simSystems.NewPlayerBlockCollisionHandler(scene))

	scene.AddSystem(renderSystems.NewCameraPositionHandler(scene))
	scene.AddSystem(renderSystems.NewParallaxBackgroundRenderer(scene))
	scene.AddSystem(renderSystems.NewPlayerRenderer(scene))
	scene.AddSystem(renderSystems.NewDebugRigidBodyRenderer(scene))

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
	AddLinearBoundsEntity(scene, float64(scene.Width/2), float64(scene.Height-10), float64(scene.Width), 20)

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
			*assetComponents.NewSprite(strconv.Itoa(layer.ZIndex)),
		)

	}

}

func AddPlayerEntity(scene *coldBrew.Scene) {

	playerEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.InputBufferComponent,
		assetComponents.SpriteComponents,
		tags.PlayerTag,
		tags.StandardCollisionTag,
	)

	// RigidBody
	playerBody := *tBokiComponents.NewRigidBodyBox(0, 200, 32, 64, 1, false)
	playerBody.Elasticity = 0

	components.RigidBodyComponent.SetValue(playerEntity, playerBody)

	// InputBuffer
	components.InputBufferComponent.SetValue(playerEntity, *components.NewInputBuffer())

	// Sprites
	playerSprites := make([]*assetComponents.Sprite, 0)

	playerSprites = append(playerSprites, assetComponents.NewSprite("temp"))

	assetComponents.SpriteComponents.SetValue(playerEntity, playerSprites)

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

func AddLinearBoundsEntity(scene *coldBrew.Scene, x, y, w, h float64) {

	boundsEntity := scene.AddEntity(
		components.RigidBodyComponent,
		tags.StandardCollisionTag,
	)
	boundsBody := tBokiComponents.NewRigidBodyBox(x, y, w, h, 0, false)
	boundsBody.Elasticity = 0

	components.RigidBodyComponent.SetValue(
		boundsEntity,
		*boundsBody,
	)

}
