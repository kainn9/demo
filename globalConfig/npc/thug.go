package npcGlobals

import (
	"github.com/kainn9/demo/components"
	sharedCombatGlobals "github.com/kainn9/demo/globalConfig/sharedCombat"
	sharedStateGlobals "github.com/kainn9/demo/globalConfig/sharedState"
)

const (
	NPC_NAME_THUG components.NpcName = "thug"

	THUG_ATTACK_RANGE = 55
	THUG_PATROL_RANGE = 400
	THUG_SPEED        = 80

	THUG_WIDTH  = 30
	THUG_HEIGHT = 75

	THUG_SPRITE_OFFSET_X = -60
	THUG_SPRITE_OFFSET_Y = -93

	THUG_ANIMATIONS_SPRITE_WIDTH  = 128
	THUG_ANIMATIONS_SPRITE_HEIGHT = 128

	THUG_IDLE_FRAME_COUNT = 4
	THUG_IDLE_ANIM_SPEED  = 10

	THUG_RUN_FRAME_COUNT = 4
	THUG_RUN_ANIM_SPEED  = 10

	THUG_HURT_FRAME_COUNT = 1
	THUG_HURT_ANIM_SPEED  = 20

	THUG_DEFEATED_FRAME_COUNT = 4
	THUG_DEFEATED_ANIM_SPEED  = 6

	THUG_DEFEATED_DURATION = 240

	THUG_ATTACK_PRIMARY_FRAME_COUNT = 7
	THUG_ATTACK_PRIMARY_ANIM_SPEED  = 6
)

// Anim Configs
var THUG_ANIMATION_CONFIGS = map[components.CharState]components.AnimationConfig{
	sharedStateGlobals.CHAR_STATE_IDLE: *components.NewAnimationConfig(
		THUG_ANIMATIONS_SPRITE_WIDTH,
		THUG_ANIMATIONS_SPRITE_HEIGHT,
		THUG_IDLE_FRAME_COUNT,
		THUG_IDLE_ANIM_SPEED,
		false,
	),

	sharedStateGlobals.CHAR_STATE_RUN: *components.NewAnimationConfig(
		THUG_ANIMATIONS_SPRITE_WIDTH,
		THUG_ANIMATIONS_SPRITE_HEIGHT,
		THUG_RUN_FRAME_COUNT,
		THUG_RUN_ANIM_SPEED,
		false,
	),

	sharedStateGlobals.CHAR_STATE_HURT: *components.NewAnimationConfig(
		THUG_ANIMATIONS_SPRITE_WIDTH,
		THUG_ANIMATIONS_SPRITE_HEIGHT,
		THUG_HURT_FRAME_COUNT,
		THUG_HURT_ANIM_SPEED,
		false,
	),

	sharedStateGlobals.CHAR_STATE_DEFEATED: *components.NewAnimationConfig(
		THUG_ANIMATIONS_SPRITE_WIDTH,
		THUG_ANIMATIONS_SPRITE_HEIGHT,
		THUG_DEFEATED_FRAME_COUNT,
		THUG_DEFEATED_ANIM_SPEED,
		true,
	),

	sharedStateGlobals.CHAR_STATE_ATTACK_PRIMARY: *components.NewAnimationConfig(
		THUG_ANIMATIONS_SPRITE_WIDTH,
		THUG_ANIMATIONS_SPRITE_HEIGHT,
		THUG_ATTACK_PRIMARY_FRAME_COUNT,
		THUG_ATTACK_PRIMARY_ANIM_SPEED,
		false,
	),
}

// Attack Primary.

// Data.
var _ = func() interface{} {

	NPCAttackDataMaps[NPC_NAME_THUG] = map[components.CharState]*sharedCombatGlobals.AttackData{
		sharedStateGlobals.CHAR_STATE_ATTACK_PRIMARY: {
			TotalTickLength: THUG_ATTACK_PRIMARY_FRAME_COUNT * THUG_ATTACK_PRIMARY_ANIM_SPEED,
			TicksPerFrame:   THUG_ATTACK_PRIMARY_ANIM_SPEED,
		},
	}

	return nil
}()

// Hitbox
var _ = func() interface{} {

	noBox := sharedCombatGlobals.EMPTY_HITBOXES_FRAME

	attackHitboxesFrameOne := sharedCombatGlobals.NewAttackHitboxesFrameData(
		sharedCombatGlobals.NewAttackHitboxData(60, 10, -0.1, 35, -30),
	)
	attackHitboxesFrameTwo := sharedCombatGlobals.NewAttackHitboxesFrameData(
		sharedCombatGlobals.NewAttackHitboxData(50, 10, -0.1, 30, -30),
	)

	attackHitboxesAllFrames := sharedCombatGlobals.NewAttackHitboxesData(
		noBox,
		noBox,
		noBox,
		noBox,
		noBox,
		attackHitboxesFrameOne,
		attackHitboxesFrameTwo,
	)

	NPCAttackHitboxesDataMaps[NPC_NAME_THUG] = map[components.CharState]sharedCombatGlobals.AttackHitboxesData{
		sharedStateGlobals.CHAR_STATE_ATTACK_PRIMARY: attackHitboxesAllFrames,
	}

	return nil
}()
