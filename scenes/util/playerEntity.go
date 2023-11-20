package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	sharedAnimationGlobals "github.com/kainn9/demo/globalConfig/sharedAnimation"
	sharedCombatGlobals "github.com/kainn9/demo/globalConfig/sharedCombat"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	"github.com/yohamta/donburi"
)

type playerEntityFactoryStruct struct{}

var PlayerEntityFactory = playerEntityFactoryStruct{}

func (f playerEntityFactoryStruct) AddPlayerEntity(scene *coldBrew.Scene, x, y float64) tBokiComponents.RigidBody {

	// Entity Initialization.
	playerEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.InputsComponent,
		components.PlayerStateComponent,
		components.SpritesCharStateMapComponent,
		components.SoundCharStateMapComponent,
		components.AttackHitboxConfigComponent,
		components.PhysicsConfigComponent,
		components.InventoryComponent,
		tags.PlayerTag,
	)

	playerBody := f.setupPhysicsComponents(x, y, playerEntity)

	f.setupStateComponents(playerEntity)
	f.setupSpritesAndAnimComponents(playerEntity)
	f.setupSoundComponents(playerEntity)
	f.setupAttackDataComponents(playerEntity)
	f.setupInventoryComponents(playerEntity)

	return playerBody
}

func (playerEntityFactoryStruct) setupPhysicsComponents(x, y float64, playerEntity *donburi.Entry) tBokiComponents.RigidBody {
	// Physics Config/modifiers.
	components.PhysicsConfigComponent.SetValue(playerEntity, *components.NewPhysicsConfig(0))

	// RigidBody.
	playerBody := *tBokiComponents.NewRigidBodyBox(x, y, playerGlobals.PLAYER_WIDTH, playerGlobals.PLAYER_HEIGHT, 1, false)
	playerBody.Elasticity = 0

	components.RigidBodyComponent.SetValue(playerEntity, playerBody)

	return playerBody

}

func (playerEntityFactoryStruct) setupStateComponents(playerEntity *donburi.Entry) {
	// Inputs.
	components.InputsComponent.SetValue(playerEntity, *components.NewInputs())

	// PlayerState.
	playerState := components.NewPlayerState()
	components.PlayerStateComponent.SetValue(playerEntity, *playerState)

}

func (playerEntityFactoryStruct) setupSpritesAndAnimComponents(playerEntity *donburi.Entry) {
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

	// Walk.
	playerSprites[sharedAnimationGlobals.CHAR_STATE_WALK] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[sharedAnimationGlobals.CHAR_STATE_WALK].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[sharedAnimationGlobals.CHAR_STATE_WALK]

	// Sit.
	playerSprites[playerGlobals.PLAYER_CHAR_STATE_SIT] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerGlobals.PLAYER_CHAR_STATE_SIT].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[playerGlobals.PLAYER_CHAR_STATE_SIT]

	// Roll.
	playerSprites[playerGlobals.PLAYER_CHAR_STATE_ROLL] = components.NewSprite(
		playerGlobals.PLAYER_SPRITE_OFFSET_X,
		playerGlobals.PLAYER_SPRITE_OFFSET_Y,
	)

	playerSprites[playerGlobals.PLAYER_CHAR_STATE_ROLL].AnimationConfig = playerGlobals.PLAYER_ANIMATION_CONFIGS[playerGlobals.PLAYER_CHAR_STATE_ROLL]

	// Set map.
	components.SpritesCharStateMapComponent.SetValue(playerEntity, playerSprites)
}

func (playerEntityFactoryStruct) setupSoundComponents(playerEntity *donburi.Entry) {
	playerSounds := make(map[components.CharState]*components.Sound, 0)

	playerSounds[sharedAnimationGlobals.CHAR_STATE_HURT] = components.NewSound(30, 1)
	playerSounds[sharedAnimationGlobals.CHAR_STATE_RUN] = components.NewSound(55, 0.5)
	playerSounds[sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY] = components.NewSound(60, 1)
	components.SoundCharStateMapComponent.SetValue(playerEntity, playerSounds)
}

func (playerEntityFactoryStruct) setupAttackDataComponents(playerEntity *donburi.Entry) {

	noBox := sharedCombatGlobals.EMPTY_HITBOX

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
}

func (playerEntityFactoryStruct) setupInventoryComponents(playerEntity *donburi.Entry) {
	// Inventory.
	components.InventoryComponent.SetValue(playerEntity, make([]*components.InventoryItem, 0))

}
