package npcGlobals

import (
	"github.com/kainn9/demo/components"
	sharedCombatGlobals "github.com/kainn9/demo/globalConfig/sharedCombat"
	"github.com/kainn9/demo/tags"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	"github.com/yohamta/donburi"
)

type AttackDataMap = map[components.CharState]*sharedCombatGlobals.AttackData
type AttackHitboxesData = map[components.CharState]sharedCombatGlobals.AttackHitboxesData

var NPCAttackDataMaps = map[components.NpcName]AttackDataMap{}
var NPCAttackHitboxesDataMaps = map[components.NpcName]AttackHitboxesData{}

// RigidBody dimensions
var NPC_DIMENSIONS = map[components.NpcName]tBokiVec.Vec2{
	NPC_NAME_THUG:          {X: THUG_WIDTH, Y: THUG_HEIGHT},
	NPC_NAME_THERAPIST_TWO: {X: THERAPIST_TWO_WIDTH, Y: THERAPIST_TWO_HEIGHT},
}

// Sprite offsets
var NPC_SPRITE_OFFSETS = map[components.NpcName]tBokiVec.Vec2{
	NPC_NAME_THUG:          {X: THUG_SPRITE_OFFSET_X, Y: THUG_SPRITE_OFFSET_Y},
	NPC_NAME_THERAPIST_TWO: {X: THERAPIST_TWO_SPRITE_OFFSET_X, Y: THERAPIST_TWO_SPRITE_OFFSET_Y},
}

// Animation configs
var NPC_ANIMATION_CONFIGS = map[components.NpcName]map[components.CharState]components.AnimationConfig{
	NPC_NAME_THUG:          THUG_ANIMATION_CONFIGS,
	NPC_NAME_THERAPIST_TWO: THERAPIST_TWO_ANIMATION_CONFIGS,
}

// Defeated durations.
var NPC_DEFEATED_DURATIONS = map[components.NpcName]int{
	NPC_NAME_THUG: THUG_DEFEATED_DURATION,
}

// Tags.
var TAG_MAP = map[components.NpcName]*donburi.ComponentType[struct{}]{
	NPC_NAME_THUG: tags.NpcThugTag,
}
