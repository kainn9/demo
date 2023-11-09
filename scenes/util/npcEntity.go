package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"
	sharedAnimationGlobals "github.com/kainn9/demo/globalConfig/sharedAnimation"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddNpcEntity(scene *coldBrew.Scene, x, y float64, name components.NpcName, hittable bool) {

	// Entity initialization
	npcEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.SpritesAnimMapComponent,
		components.NpcConfigComponent,
		components.NpcStateComponent,
		tags.NpcTag,
	)

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

	components.SpritesAnimMapComponent.SetValue(npcEntity, npcSprites)
}
