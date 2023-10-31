package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

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

	// Climb Ladder.
	playerSprites[constants.PLAYER_ANIM_STATE_CLIMB_LADDER_IDLE] = assetComponents.NewSprite(
		constants.PLAYER_SPRITE_OFFSET_X,
		constants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[constants.PLAYER_ANIM_STATE_CLIMB_LADDER_IDLE].AnimationConfig = assetComponents.NewAnimationConfig(
		constants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		constants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		constants.PLAYER_CLIMB_LADDER_IDLE_FRAME_COUNT,
		constants.PLAYER_CLIMB_LADDER_ANIM_SPEED,
		false,
	)

	// Climb Ladder Active
	playerSprites[constants.PLAYER_ANIM_STATE_CLIMB_LADDER_ACTIVE] = assetComponents.NewSprite(
		constants.PLAYER_SPRITE_OFFSET_X,
		constants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[constants.PLAYER_ANIM_STATE_CLIMB_LADDER_ACTIVE].AnimationConfig = assetComponents.NewAnimationConfig(
		constants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		constants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		constants.PLAYER_CLIMB_LADDER_ACTIVE_FRAME_COUNT,
		constants.PLAYER_CLIMB_LADDER_ANIM_SPEED,
		false,
	)

	assetComponents.SpritesMapComponent.SetValue(playerEntity, playerSprites)

	return playerBody
}
