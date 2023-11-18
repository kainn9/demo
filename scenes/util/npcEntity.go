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

func AddNpcEntity(
	scene *coldBrew.Scene,
	x, y float64,
	name components.NpcName,
	physicsMod *components.PhysicsConfig,
	hittable, moveable bool,
) *donburi.Entry {

	tag := npcGlobals.TAG_MAP[name]

	if tag == nil {
		tag = tags.EmptyTag
	}

	npcEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.SpritesAnimMapComponent,
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
	state := components.NewNpcState(hittable)
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

	// Only prep combat sprites if hittable.
	if hittable {
		prepCombatSpriteSheets(spriteOffset, npcSprites, animationConfigs)
	}

	if moveable {
		prepTransformSprites(spriteOffset, npcSprites, animationConfigs)

	}

	components.SpritesAnimMapComponent.SetValue(npcEntity, npcSprites)

	return npcEntity
}

func prepCombatSpriteSheets(
	spriteOffset tBokiVec.Vec2,
	npcSprites map[components.CharState]*components.Sprite,
	animationConfigs map[components.CharState]components.AnimationConfig,
) {
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
