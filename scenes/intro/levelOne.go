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
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

type LevelOneScene struct{}

const (
	SCENE_WIDTH  = 1920
	SCENE_HEIGHT = 622
	SCENE_NAME   = "levelOne"
)

func (LevelOneScene) Index() string {
	return SCENE_NAME
}

func (LevelOneScene) New(m *coldBrew.Manager) *coldBrew.Scene {

	scene := coldBrew.NewScene(m, SCENE_WIDTH, SCENE_HEIGHT)

	// Systems ----------------------------------------------------------------------------------

	// Loader Systems.
	scene.AddSystem(loaderSystems.NewPlayerSpritesLoader(scene))
	scene.AddSystem(loaderSystems.NewParallaxBackgroundLoader(scene))

	// Client Systems.
	scene.AddSystem(clientSystems.NewInputTracker())

	// Sim Systems.
	scene.AddSystem(simSystems.NewPlayerInputHandler(scene))
	scene.AddSystem(simSystems.NewPlayerMovementHandler(scene))
	scene.AddSystem(simSystems.NewPlayerBlockCollisionHandler(scene))
	scene.AddSystem(simSystems.NewCameraPositionHandler(scene))

	// Render Systems.
	scene.AddSystem(renderSystems.NewParallaxBackgroundRenderer(scene))
	scene.AddSystem(renderSystems.NewPlayerRenderer(scene))
	scene.AddSystem(renderSystems.NewDebugRigidBodyRenderer(scene))

	// Entities ----------------------------------------------------------------------------------
	parallaxBackgroundSubPath := "intro/levelOne/"
	AddParallaxBackground(scene, []*assetComponents.ParallaxLayerConfig{
		// Sky.
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 0, 22, 22, false),

		// City Far Shadow.
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 1, 18, 18, false),

		// City Lights.
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 2, 8, 8, false),

		// Mountains
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 3, 0, 0, false),

		// Green Trees.
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 4, 10, 10, false),

		// Statues.
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 5, 5, 12, false),

		// Red Shrubs.
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 6, 4, 8, false),

		// Gears Close.
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 7, 2, 0, false),
		// Main layer.
		assetComponents.NewParallaxLayerConfig(parallaxBackgroundSubPath, 8, 0, 0, false),
	})

	AddPlayerEntity(scene, 100, 600)
	AddFloorEntity(scene, float64(scene.Width/2), float64(scene.Height-40), float64(scene.Width), 20, -0.04)
	AddCameraEntity(scene, 0, 262)

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
			*assetComponents.NewSprite(0, 0),
		)

	}

}

func AddPlayerEntity(scene *coldBrew.Scene, x, y float64) tBokiComponents.RigidBody {

	// Entity Initialization.
	playerEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.InputsComponent,
		components.PlayerStateComponent,
		assetComponents.SpritesMapComponent,
		tags.PlayerTag,
		tags.StandardCollisionTag,
	)

	// RigidBody.
	playerBody := *tBokiComponents.NewRigidBodyBox(x, y, constants.PLAYER_WIDTH, constants.PLAYER_HEIGHT, 1, false)
	playerBody.Elasticity = 0

	components.RigidBodyComponent.SetValue(playerEntity, playerBody)

	// Inputs.
	components.InputsComponent.SetValue(playerEntity, *components.NewInputs())

	// PlayerState.
	playerState := components.NewPlayerState()
	components.PlayerStateComponent.SetValue(playerEntity, *playerState)

	// Sprites/Animations.
	playerSprites := make(map[string]*assetComponents.Sprite, 0)

	// Idle.
	playerSprites[constants.PLAYER_ANIM_STATE_IDLE] = assetComponents.NewSprite(
		constants.PLAYER_SPRITE_OFFSET_X,
		constants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[constants.PLAYER_ANIM_STATE_IDLE].AnimationData = assetComponents.NewAnimationData(
		constants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		constants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		constants.PLAYER_IDLE_FRAME_COUNT,
		constants.PLAYER_IDLE_ANIM_SPEED,
		false,
	)

	// Run.
	playerSprites[constants.PLAYER_ANIM_STATE_RUN] = assetComponents.NewSprite(
		constants.PLAYER_SPRITE_OFFSET_X,
		constants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[constants.PLAYER_ANIM_STATE_RUN].AnimationData = assetComponents.NewAnimationData(
		constants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		constants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		constants.PLAYER_RUN_FRAME_COUNT,
		constants.PLAYER_RUN_ANIM_SPEED,
		false,
	)

	// Jump.
	playerSprites[constants.PLAYER_ANIM_STATE_JUMP] = assetComponents.NewSprite(
		constants.PLAYER_SPRITE_OFFSET_X,
		constants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[constants.PLAYER_ANIM_STATE_JUMP].AnimationData = assetComponents.NewAnimationData(
		constants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		constants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		constants.PLAYER_JUMP_FRAME_COUNT,
		constants.PLAYER_JUMP_ANIM_SPEED,
		true,
	)

	// Fall.
	playerSprites[constants.PLAYER_ANIM_STATE_FALL] = assetComponents.NewSprite(
		constants.PLAYER_SPRITE_OFFSET_X,
		constants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[constants.PLAYER_ANIM_STATE_FALL].AnimationData = assetComponents.NewAnimationData(
		constants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		constants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		constants.PLAYER_FALL_FRAME_COUNT,
		constants.PLAYER_FALL_ANIM_SPEED,
		true,
	)

	assetComponents.SpritesMapComponent.SetValue(playerEntity, playerSprites)

	return playerBody
}

func AddCameraEntity(scene *coldBrew.Scene, x, y float64) {

	cameraEntity := scene.AddEntity(
		components.CameraComponent,
	)

	components.CameraComponent.SetValue(
		cameraEntity,
		*components.NewCamera(x, y, constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT),
	)

}

func AddFloorEntity(scene *coldBrew.Scene, x, y, w, h, rotation float64) {

	boundsEntity := scene.AddEntity(
		components.RigidBodyComponent,
		tags.StandardCollisionTag,
	)

	boundsBody := tBokiComponents.NewRigidBodyBox(x, y, w, h, 0, true)
	boundsBody.Elasticity = 0
	boundsBody.Rotation = rotation
	boundsBody.UpdateVertices()

	components.RigidBodyComponent.SetValue(
		boundsEntity,
		*boundsBody,
	)
}
