package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	sharedAnimationGlobals "github.com/kainn9/demo/globalConfig/sharedAnimation"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddPlayerEntity(scene *coldBrew.Scene, x, y float64) tBokiComponents.RigidBody {

	// Entity Initialization.
	playerEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.InputsComponent,
		components.PlayerStateComponent,
		components.SpritesAnimMapComponent,
		components.AttackHitboxConfigComponent,
		tags.PlayerTag,
	)

	// RigidBody.
	playerBody := *tBokiComponents.NewRigidBodyBox(x, y, playerGlobals.PLAYER_WIDTH, playerGlobals.PLAYER_HEIGHT, 1, false)
	playerBody.Elasticity = 0

	components.RigidBodyComponent.SetValue(playerEntity, playerBody)

	// Inputs.
	components.InputsComponent.SetValue(playerEntity, *components.NewInputs())

	// PlayerState.
	playerState := components.NewPlayerState()
	components.PlayerStateComponent.SetValue(playerEntity, *playerState)

	// Sprites/Animations.
	playerSprites := make(map[components.CharState]*components.Sprite, 0)

	// Idle.
	playerSprites[sharedAnimationGlobals.CHAR_STATE_IDLE] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[sharedAnimationGlobals.CHAR_STATE_IDLE].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[sharedAnimationGlobals.CHAR_STATE_IDLE]

	// Run.
	playerSprites[sharedAnimationGlobals.CHAR_STATE_RUN] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[sharedAnimationGlobals.CHAR_STATE_RUN].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[sharedAnimationGlobals.CHAR_STATE_RUN]

	// Jump.
	playerSprites[sharedAnimationGlobals.CHAR_STATE_JUMP] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[sharedAnimationGlobals.CHAR_STATE_JUMP].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[sharedAnimationGlobals.CHAR_STATE_JUMP]

	// Fall.
	playerSprites[sharedAnimationGlobals.CHAR_STATE_FALL] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[sharedAnimationGlobals.CHAR_STATE_FALL].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[sharedAnimationGlobals.CHAR_STATE_FALL]

	// Climb Ladder.
	playerSprites[playerGlobals.PLAYER_CHAR_STATE_CLIMB_LADDER_IDLE] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerGlobals.PLAYER_CHAR_STATE_CLIMB_LADDER_IDLE].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[playerGlobals.PLAYER_CHAR_STATE_CLIMB_LADDER_IDLE]

	// Climb Ladder Active.
	playerSprites[playerGlobals.PLAYER_CHAR_STATE_CLIMB_LADDER_ACTIVE] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerGlobals.PLAYER_CHAR_STATE_CLIMB_LADDER_ACTIVE].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[playerGlobals.PLAYER_CHAR_STATE_CLIMB_LADDER_ACTIVE]

	// Attack Primary.
	playerSprites[sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY]

	// Hurt.
	playerSprites[sharedAnimationGlobals.CHAR_STATE_HURT] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[sharedAnimationGlobals.CHAR_STATE_HURT].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[sharedAnimationGlobals.CHAR_STATE_HURT]

	// Defeated.

	playerSprites[sharedAnimationGlobals.CHAR_STATE_DEFEATED] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[sharedAnimationGlobals.CHAR_STATE_DEFEATED].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[sharedAnimationGlobals.CHAR_STATE_DEFEATED]

	components.SpritesAnimMapComponent.SetValue(playerEntity, playerSprites)

	// Hitboxes
	noBox := []components.HitboxData{components.NewHitboxData(0, 0, 0, 0, 0)}

	hitboxesDataFrame1 := []components.HitboxData{
		components.NewHitboxData(50, 10, 0, 30, -4),
	}

	hitboxes := components.NewAttackHitboxConfig(
		noBox,
		noBox,
		hitboxesDataFrame1,
		hitboxesDataFrame1,
		hitboxesDataFrame1,
		noBox, // 5
		noBox, // 6
		noBox, // 7
		noBox, // 8
	)

	components.AttackHitboxConfigComponent.SetValue(playerEntity, *hitboxes)

	return playerBody
}
