package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddParallaxBackgroundEntity(scene *coldBrew.Scene, layers []*assetComponents.ParallaxLayerConfig) {

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

func AddFrontLayerEntity(scene *coldBrew.Scene, sceneAssetPath string) {
	frontLayerBackgroundEntity := scene.AddEntity(assetComponents.FrontLayerComponent, assetComponents.SpriteComponent)

	assetComponents.FrontLayerComponent.SetValue(
		frontLayerBackgroundEntity,
		*assetComponents.NewFrontLayerConfig(sceneAssetPath),
	)

	assetComponents.SpriteComponent.SetValue(
		frontLayerBackgroundEntity,
		*assetComponents.NewSprite(0, 0),
	)
}

func AddPlayerEntity(scene *coldBrew.Scene, x, y float64) tBokiComponents.RigidBody {

	// Entity Initialization.
	playerEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.InputsComponent,
		components.PlayerStateComponent,
		assetComponents.SpritesMapComponent,
		tags.PlayerTag,
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

	playerSprites[constants.PLAYER_ANIM_STATE_IDLE].AnimationConfig = assetComponents.NewAnimationConfig(
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

	playerSprites[constants.PLAYER_ANIM_STATE_RUN].AnimationConfig = assetComponents.NewAnimationConfig(
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

	playerSprites[constants.PLAYER_ANIM_STATE_JUMP].AnimationConfig = assetComponents.NewAnimationConfig(
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

	playerSprites[constants.PLAYER_ANIM_STATE_FALL].AnimationConfig = assetComponents.NewAnimationConfig(
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

	floorEntity := scene.AddEntity(
		components.RigidBodyComponent,
		tags.FloorTag,
	)

	floorBody := tBokiComponents.NewRigidBodyBox(x, y, w, h, 0, true)
	floorBody.Elasticity = 0
	floorBody.Rotation = rotation
	floorBody.UpdateVertices()

	components.RigidBodyComponent.SetValue(
		floorEntity,
		*floorBody,
	)
}

// Note: Platforms do not support rotation.
func AddPlatformEntity(scene *coldBrew.Scene, x, y, w, h float64) {

	platformEntity := scene.AddEntity(
		components.RigidBodyComponent,
		tags.PlatformTag,
	)

	platformBody := tBokiComponents.NewRigidBodyBox(x, y, w, h, 0, true)
	platformBody.Elasticity = 0
	platformBody.UpdateVertices()

	components.RigidBodyComponent.SetValue(
		platformEntity,
		*platformBody,
	)
}

// Note: Ladders do not support rotation.
func AddLadderEntity(scene *coldBrew.Scene, x, y, w, h float64) {

	ladderEntity := scene.AddEntity(
		components.RigidBodyComponent,
		tags.LadderTag,
	)

	ladderBody := tBokiComponents.NewRigidBodyBox(x, y, w, h, 0, true)
	ladderBody.UpdateVertices()

	components.RigidBodyComponent.SetValue(
		ladderEntity,
		*ladderBody,
	)
}

func AddChatEntity(scene *coldBrew.Scene, slidesCount int, introChat, sceneAssetPath string, portraitNames []string) {

	chatEntity := scene.AddEntity(
		components.ChatStateComponent,
		assetComponents.SpritesSliceComponent,
		assetComponents.SpritesMapComponent,
	)

	chatState := components.NewChatState(introChat, sceneAssetPath)
	chatState.Active = true
	chatState.PopUpMode = true
	chatState.PortraitNames = portraitNames

	slideSprites := make([]*assetComponents.Sprite, slidesCount)
	portraits := make(map[string]*assetComponents.Sprite, 0)

	for i := 0; i < slidesCount; i++ {
		slideSprites[i] = assetComponents.NewSprite(0, 0)
		portraits[portraitNames[i]] = assetComponents.NewSprite(0, 0)
	}

	assetComponents.SpritesMapComponent.SetValue(
		chatEntity,
		portraits,
	)

	assetComponents.SpritesSliceComponent.SetValue(
		chatEntity,
		slideSprites,
	)

	components.ChatStateComponent.SetValue(
		chatEntity,
		*chatState,
	)

}
