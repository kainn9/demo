package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"
	sharedAnimationGlobals "github.com/kainn9/demo/globalConfig/sharedAnimation"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	"github.com/yohamta/donburi"
)

// TODO: DRY out common code between AddStaticNpcEntity and AddCombatNpcEntity.

func AddStaticNpcEntity(
	scene *coldBrew.Scene,
	x, y float64,
	name components.NpcName,
	physicsMod *components.PhysicsConfig,
) *donburi.Entry {

	tag := npcGlobals.TAG_MAP[name]

	if tag == nil {
		tag = tags.EmptyTag
	}

	npcEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.SpritesCharStateMapComponent,
		components.NpcConfigComponent,
		components.NpcStateComponent,
		components.PhysicsConfigComponent,
		tags.NpcTag,
		tag,
	)

	// Physics Config/modifiers.
	if physicsMod != nil {
		components.PhysicsConfigComponent.SetValue(npcEntity, *physicsMod)
	}

	// Config.
	config := components.NewNpcConfig(name)
	components.NpcConfigComponent.SetValue(npcEntity, *config)

	// State.
	state := components.NewNpcState(false, 0, 0, 0)
	components.NpcStateComponent.SetValue(npcEntity, *state)

	// RigidBody.
	bodyDimensions := npcGlobals.NPC_DIMENSIONS[name]
	npcBody := *tBokiComponents.NewRigidBodyBox(x, y, bodyDimensions.X, bodyDimensions.Y, 1, false)
	npcBody.Elasticity = 0

	components.RigidBodyComponent.SetValue(npcEntity, npcBody)

	// Sprites/Animations.
	npcSprites := make(map[components.CharState]*components.Sprite, 0)
	spriteOffset := npcGlobals.NPC_SPRITE_OFFSETS[name]
	animationConfigs := npcGlobals.NPC_ANIMATION_CONFIGS[name]

	// Idle.
	npcSprites[sharedAnimationGlobals.CHAR_STATE_IDLE] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)

	idleConfig := animationConfigs[sharedAnimationGlobals.CHAR_STATE_IDLE]
	npcSprites[sharedAnimationGlobals.CHAR_STATE_IDLE].AnimationConfig = &idleConfig

	components.SpritesCharStateMapComponent.SetValue(npcEntity, npcSprites)

	return npcEntity
}

func AddCombatNpcEntity(
	scene *coldBrew.Scene,
	x, y float64,
	name components.NpcName,
	physicsMod *components.PhysicsConfig,
	attackRange, patrolRange, speed float64,
) *donburi.Entry {

	tag := npcGlobals.TAG_MAP[name]

	if tag == nil {
		tag = tags.EmptyTag
	}

	npcEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.SpritesCharStateMapComponent,
		components.NpcConfigComponent,
		components.NpcStateComponent,
		components.PhysicsConfigComponent,
		components.AttackHitboxConfigComponent,
		tags.NpcTag,
		tag,
	)

	// Physics Config/modifiers.
	if physicsMod != nil {
		components.PhysicsConfigComponent.SetValue(npcEntity, *physicsMod)
	}

	// Config.
	config := components.NewNpcConfig(name)
	components.NpcConfigComponent.SetValue(npcEntity, *config)

	// State.
	state := components.NewNpcState(true, attackRange, patrolRange, speed)
	components.NpcStateComponent.SetValue(npcEntity, *state)

	// RigidBody.
	bodyDimensions := npcGlobals.NPC_DIMENSIONS[name]
	npcBody := *tBokiComponents.NewRigidBodyBox(x, y, bodyDimensions.X, bodyDimensions.Y, 1, false)
	npcBody.Elasticity = 0

	components.RigidBodyComponent.SetValue(npcEntity, npcBody)

	// Sprites/Animations.
	npcSprites := make(map[components.CharState]*components.Sprite, 0)
	spriteOffset := npcGlobals.NPC_SPRITE_OFFSETS[name]
	animationConfigs := npcGlobals.NPC_ANIMATION_CONFIGS[name]

	// Idle.
	npcSprites[sharedAnimationGlobals.CHAR_STATE_IDLE] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)

	idleConfig := animationConfigs[sharedAnimationGlobals.CHAR_STATE_IDLE]
	npcSprites[sharedAnimationGlobals.CHAR_STATE_IDLE].AnimationConfig = &idleConfig

	prepCombatSpriteSheets(spriteOffset, npcSprites, animationConfigs)
	prepTransformSprites(spriteOffset, npcSprites, animationConfigs)

	components.SpritesCharStateMapComponent.SetValue(npcEntity, npcSprites)

	// Hitboxes.
	prepHitboxes(npcEntity)

	return npcEntity
}

func prepCombatSpriteSheets(
	spriteOffset tBokiVec.Vec2,
	npcSprites map[components.CharState]*components.Sprite,
	animationConfigs map[components.CharState]components.AnimationConfig,
) {

	// Attack
	npcSprites[sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)
	attackConfig := animationConfigs[sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY]
	npcSprites[sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY].AnimationConfig = &attackConfig

	// Hurt
	npcSprites[sharedAnimationGlobals.CHAR_STATE_HURT] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)

	hurtConfig := animationConfigs[sharedAnimationGlobals.CHAR_STATE_HURT]
	npcSprites[sharedAnimationGlobals.CHAR_STATE_HURT].AnimationConfig = &hurtConfig

	// Defeated

	npcSprites[sharedAnimationGlobals.CHAR_STATE_DEFEATED] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)

	defeatedConfig := animationConfigs[sharedAnimationGlobals.CHAR_STATE_DEFEATED]
	npcSprites[sharedAnimationGlobals.CHAR_STATE_DEFEATED].AnimationConfig = &defeatedConfig

}

func prepTransformSprites(
	spriteOffset tBokiVec.Vec2,
	npcSprites map[components.CharState]*components.Sprite,
	animationConfigs map[components.CharState]components.AnimationConfig,
) {

	npcSprites[sharedAnimationGlobals.CHAR_STATE_RUN] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)

	runConfig := animationConfigs[sharedAnimationGlobals.CHAR_STATE_RUN]
	npcSprites[sharedAnimationGlobals.CHAR_STATE_RUN].AnimationConfig = &runConfig

}

func prepHitboxes(npcEntity *donburi.Entry) {

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

	components.AttackHitboxConfigComponent.SetValue(npcEntity, *hitboxes)

}
