package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
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
		components.PlayerSpritesAnimMapComponent,
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
	playerSprites := make(map[components.AnimState]*components.Sprite, 0)

	// Idle.

	playerSprites[playerConstants.PLAYER_ANIM_STATE_IDLE] = components.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_IDLE].AnimationConfig = playerConstants.PLAYER_ANIMATION_CONFIGS[playerConstants.PLAYER_ANIM_STATE_IDLE]

	// Run.
	playerSprites[playerConstants.PLAYER_ANIM_STATE_RUN] = components.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_RUN].AnimationConfig = playerConstants.PLAYER_ANIMATION_CONFIGS[playerConstants.PLAYER_ANIM_STATE_RUN]

	// Jump.
	playerSprites[playerConstants.PLAYER_ANIM_STATE_JUMP] = components.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_JUMP].AnimationConfig = playerConstants.PLAYER_ANIMATION_CONFIGS[playerConstants.PLAYER_ANIM_STATE_JUMP]

	// Fall.
	playerSprites[playerConstants.PLAYER_ANIM_STATE_FALL] = components.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_FALL].AnimationConfig = playerConstants.PLAYER_ANIMATION_CONFIGS[playerConstants.PLAYER_ANIM_STATE_FALL]

	// Climb Ladder.
	playerSprites[playerConstants.PLAYER_ANIM_STATE_CLIMB_LADDER_IDLE] = components.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_CLIMB_LADDER_IDLE].AnimationConfig = playerConstants.PLAYER_ANIMATION_CONFIGS[playerConstants.PLAYER_ANIM_STATE_CLIMB_LADDER_IDLE]

	// Climb Ladder Active
	playerSprites[playerConstants.PLAYER_ANIM_STATE_CLIMB_LADDER_ACTIVE] = components.NewSprite(
		playerConstants.PLAYER_SPRITE_OFFSET_X,
		playerConstants.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerConstants.PLAYER_ANIM_STATE_CLIMB_LADDER_ACTIVE].AnimationConfig = playerConstants.PLAYER_ANIMATION_CONFIGS[playerConstants.PLAYER_ANIM_STATE_CLIMB_LADDER_ACTIVE]

	components.PlayerSpritesAnimMapComponent.SetValue(playerEntity, playerSprites)

	return playerBody
}
