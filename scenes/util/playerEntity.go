package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	playerConstants "github.com/kainn9/demo/constants/player"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddPlayerEntity(scene *coldBrew.Scene, x, y float64) tBokiComponents.RigidBody {

	// Entity Initialization.
	playerEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.InputsComponent,
		components.PlayerStateComponent,
		assetComponents.PlayerSpritesAnimMapComponent,
		tags.PlayerTag,
	)

	// RigidBody.
	playerBody := *tBokiComponents.NewRigidBodyBox(x, y, playerConstants.PLAYER_WIDTH, playerConstants.PLAYER_HEIGHT, 1, false)
	playerBody.Elasticity = 0

	components.RigidBodyComponent.SetValue(playerEntity, playerBody)

	// Inputs.
	components.InputsComponent.SetValue(playerEntity, *components.NewInputs())

	// PlayerState.
	playerState := components.NewPlayerState()
	components.PlayerStateComponent.SetValue(playerEntity, *playerState)

	// Sprites/Animations.
	playerSprites := make(map[playerConstants.AnimState]*assetComponents.Sprite, 0)

	// Idle.

	playerSprites[playerConstants.PLAYER_ANIM_STATE_IDLE] = assetComponents.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_IDLE].AnimationConfig = assetComponents.NewAnimationConfig(
		playerConstants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		playerConstants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		playerConstants.PLAYER_IDLE_FRAME_COUNT,
		playerConstants.PLAYER_IDLE_ANIM_SPEED,
		false,
	)

	// Run.
	playerSprites[playerConstants.PLAYER_ANIM_STATE_RUN] = assetComponents.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_RUN].AnimationConfig = assetComponents.NewAnimationConfig(
		playerConstants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		playerConstants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		playerConstants.PLAYER_RUN_FRAME_COUNT,
		playerConstants.PLAYER_RUN_ANIM_SPEED,
		false,
	)

	// Jump.
	playerSprites[playerConstants.PLAYER_ANIM_STATE_JUMP] = assetComponents.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_JUMP].AnimationConfig = assetComponents.NewAnimationConfig(
		playerConstants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		playerConstants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		playerConstants.PLAYER_JUMP_FRAME_COUNT,
		playerConstants.PLAYER_JUMP_ANIM_SPEED,
		true,
	)

	// Fall.
	playerSprites[playerConstants.PLAYER_ANIM_STATE_FALL] = assetComponents.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_FALL].AnimationConfig = assetComponents.NewAnimationConfig(
		playerConstants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		playerConstants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		playerConstants.PLAYER_FALL_FRAME_COUNT,
		playerConstants.PLAYER_FALL_ANIM_SPEED,
		true,
	)

	// Climb Ladder.
	playerSprites[playerConstants.PLAYER_ANIM_STATE_CLIMB_LADDER_IDLE] = assetComponents.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_CLIMB_LADDER_IDLE].AnimationConfig = assetComponents.NewAnimationConfig(
		playerConstants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		playerConstants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		playerConstants.PLAYER_CLIMB_LADDER_IDLE_FRAME_COUNT,
		playerConstants.PLAYER_CLIMB_LADDER_ANIM_SPEED,
		false,
	)

	// Climb Ladder Active
	playerSprites[playerConstants.PLAYER_ANIM_STATE_CLIMB_LADDER_ACTIVE] = assetComponents.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_CLIMB_LADDER_ACTIVE].AnimationConfig = assetComponents.NewAnimationConfig(
		playerConstants.PLAYER_ANIMATIONS_SPRITE_WIDTH,
		playerConstants.PLAYER_ANIMATIONS_SPRITE_HEIGHT,
		playerConstants.PLAYER_CLIMB_LADDER_ACTIVE_FRAME_COUNT,
		playerConstants.PLAYER_CLIMB_LADDER_ANIM_SPEED,
		false,
	)

	assetComponents.PlayerSpritesAnimMapComponent.SetValue(playerEntity, playerSprites)

	return playerBody
}
