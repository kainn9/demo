package playerGlobals

import (
	"github.com/kainn9/demo/components"
	sharedCombatGlobals "github.com/kainn9/demo/globalConfig/sharedCombat"
	sharedStateGlobals "github.com/kainn9/demo/globalConfig/sharedState"
)

// Attack Hitboxes.
var AttackHitboxesData = map[components.CharState]sharedCombatGlobals.AttackHitboxesData{}

// Attack Primary.
var _ = func() interface{} {

	noBox := sharedCombatGlobals.EMPTY_HITBOXES_FRAME

	attackHitboxesFrameOne := sharedCombatGlobals.NewAttackHitboxesFrameData(
		sharedCombatGlobals.NewAttackHitboxData(50, 10, 0, 30, -4),
	)

	attackHitboxesAllFrames := sharedCombatGlobals.NewAttackHitboxesData(
		noBox,
		noBox,
		noBox,
		attackHitboxesFrameOne,
		attackHitboxesFrameOne,
		attackHitboxesFrameOne,
		noBox,
		noBox,
		noBox,
	)

	AttackHitboxesData[sharedStateGlobals.CHAR_STATE_ATTACK_PRIMARY] = attackHitboxesAllFrames

	return nil
}()

// Attack Data.
var AttackDataMap = map[components.CharState]*sharedCombatGlobals.AttackData{}

// Attack Primary.
var _ = func() interface{} {

	AttackDataMap[sharedStateGlobals.CHAR_STATE_ATTACK_PRIMARY] = &sharedCombatGlobals.AttackData{
		TotalTickLength: PLAYER_ATTACK_PRIMARY_FRAME_COUNT * PLAYER_ATTACK_PRIMARY_ANIM_SPEED,
		TicksPerFrame:   PLAYER_ATTACK_PRIMARY_ANIM_SPEED,
	}

	return nil
}()
