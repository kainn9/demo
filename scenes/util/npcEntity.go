package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"
	sharedStateGlobals "github.com/kainn9/demo/globalConfig/sharedState"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	"github.com/yohamta/donburi"
)

type npcEntityFactoryStruct struct{}

var NpcEntityFactory = npcEntityFactoryStruct{}

// Todo: DRY out common code between AddStaticNpcEntity and AddCombatNpcEntity.

func (npcEntityFactoryStruct) addStaticNpcEntity(
	scene *coldBrew.Scene,
	x, y float64,
	name components.NpcName,
	physicsMod *components.PhysicsModConfig,
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
	state := components.NewNpcState(false, 0, 0, 0, 0, 0)
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
	npcSprites[sharedStateGlobals.CHAR_STATE_IDLE] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)

	idleConfig := animationConfigs[sharedStateGlobals.CHAR_STATE_IDLE]
	npcSprites[sharedStateGlobals.CHAR_STATE_IDLE].AnimationConfig = &idleConfig

	components.SpritesCharStateMapComponent.SetValue(npcEntity, npcSprites)

	return npcEntity
}

func (f npcEntityFactoryStruct) addCombatNpcEntity(
	scene *coldBrew.Scene,
	x, y float64,
	name components.NpcName,
	physicsMod *components.PhysicsModConfig,
	attackRange, patrolRange, maxLeft, maxRight, speed float64,
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
	state := components.NewNpcState(true, attackRange, patrolRange, maxLeft, maxRight, speed)
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
	npcSprites[sharedStateGlobals.CHAR_STATE_IDLE] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)

	idleConfig := animationConfigs[sharedStateGlobals.CHAR_STATE_IDLE]
	npcSprites[sharedStateGlobals.CHAR_STATE_IDLE].AnimationConfig = &idleConfig

	f.prepCombatSpriteSheets(spriteOffset, npcSprites, animationConfigs)
	f.prepTransformSprites(spriteOffset, npcSprites, animationConfigs)

	components.SpritesCharStateMapComponent.SetValue(npcEntity, npcSprites)

	return npcEntity
}

func (npcEntityFactoryStruct) prepCombatSpriteSheets(
	spriteOffset tBokiVec.Vec2,
	npcSprites map[components.CharState]*components.Sprite,
	animationConfigs map[components.CharState]components.AnimationConfig,
) {

	// Attack
	npcSprites[sharedStateGlobals.CHAR_STATE_ATTACK_PRIMARY] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)
	attackConfig := animationConfigs[sharedStateGlobals.CHAR_STATE_ATTACK_PRIMARY]
	npcSprites[sharedStateGlobals.CHAR_STATE_ATTACK_PRIMARY].AnimationConfig = &attackConfig

	// Hurt
	npcSprites[sharedStateGlobals.CHAR_STATE_HURT] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)

	hurtConfig := animationConfigs[sharedStateGlobals.CHAR_STATE_HURT]
	npcSprites[sharedStateGlobals.CHAR_STATE_HURT].AnimationConfig = &hurtConfig

	// Defeated

	npcSprites[sharedStateGlobals.CHAR_STATE_DEFEATED] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)

	defeatedConfig := animationConfigs[sharedStateGlobals.CHAR_STATE_DEFEATED]
	npcSprites[sharedStateGlobals.CHAR_STATE_DEFEATED].AnimationConfig = &defeatedConfig

}

func (npcEntityFactoryStruct) prepTransformSprites(
	spriteOffset tBokiVec.Vec2,
	npcSprites map[components.CharState]*components.Sprite,
	animationConfigs map[components.CharState]components.AnimationConfig,
) {

	npcSprites[sharedStateGlobals.CHAR_STATE_RUN] = components.NewSprite(
		spriteOffset.X,
		spriteOffset.Y,
	)

	runConfig := animationConfigs[sharedStateGlobals.CHAR_STATE_RUN]
	npcSprites[sharedStateGlobals.CHAR_STATE_RUN].AnimationConfig = &runConfig

}

func (f npcEntityFactoryStruct) AddNpcThug(scene *coldBrew.Scene, x, y, maxLeft, maxRight float64) *donburi.Entry {
	return f.addCombatNpcEntity(scene, x, y, npcGlobals.NPC_NAME_THUG, nil, npcGlobals.THUG_ATTACK_RANGE, npcGlobals.THUG_PATROL_RANGE, maxLeft, maxRight, npcGlobals.THUG_SPEED)
}

func (f npcEntityFactoryStruct) AddTherapistTwo(scene *coldBrew.Scene, x, y float64, physicsMod *components.PhysicsModConfig) *donburi.Entry {
	return f.addStaticNpcEntity(scene, -200, -200, npcGlobals.NPC_NAME_THERAPIST_TWO, physicsMod)
}
