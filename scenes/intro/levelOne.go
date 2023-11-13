package introScenes

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/systems/systemInitializers"

	UIGlobals "github.com/kainn9/demo/globalConfig/UI"
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"
	scenesUtil "github.com/kainn9/demo/scenes/util"
)

type LevelOneScene struct{}

const (
	LEVEL_ONE_SCENE_WIDTH      = 4030
	LEVEL_ONE_SCENE_HEIGHT     = 360
	LEVEL_ONE_SCENE_NAME       = "levelOne"
	LEVEL_ONE_SCENE_SECTION    = "intro"
	LEVEL_ONE_SCENE_ASSET_PATH = LEVEL_ONE_SCENE_SECTION + "/" + LEVEL_ONE_SCENE_NAME + "/"
)

func (LevelOneScene) Index() string {
	return LEVEL_ONE_SCENE_NAME
}

func (LevelOneScene) New(m *coldBrew.Manager) *coldBrew.Scene {

	scene := coldBrew.NewScene(m, LEVEL_ONE_SCENE_WIDTH, LEVEL_ONE_SCENE_HEIGHT)

	// Systems ----------------------------------------------------------------------------------
	systemInitializers.InitStandardSystems(scene, false)

	// Entities ----------------------------------------------------------------------------------

	scenesUtil.AddCameraEntity(scene, 0, 0)

	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		// Sky.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 0, 0, 0, false),

		// Moon
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 1, 0, 0, true),

		// Clouds Back.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 2, 16, 0, false),

		// Clouds Front.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 3, 8, 0, false),

		// City.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 4, 0, 0, false),
		// Front.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 5, 0, 0, false),
	})

	// Floors left to right.
	scenesUtil.AddFloorEntity(scene, 255, float64(scene.Height), 514, 95, 0)
	scenesUtil.AddFloorEntity(scene, 937, float64(scene.Height), 515, 95, 0)
	scenesUtil.AddFloorEntity(scene, 3445, float64(scene.Height), 1180, 95, 0)

	// Trash Bins.
	// Todo: make and change to block entity.
	scenesUtil.AddFloorEntity(scene, 994, 290, 283, 130, 0)

	// Ladder.
	scenesUtil.AddLadderEntity(scene, 1130, 180, 40, 360)

	// Platforms left to right.
	scenesUtil.AddPlatformEntity(scene, 1491, 87, 322, 10)
	scenesUtil.AddPlatformEntity(scene, 2037, 191, 320, 10)
	scenesUtil.AddPlatformEntity(scene, 2597, 201, 320, 10)

	// Tutorial indicators.
	// Movement.
	scenesUtil.AddOnCollisionIndicatorEntity(
		scene,
		147,
		275,
		10,
		15,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_MOVEMENT].X,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_MOVEMENT].Y,
		true,
		UIGlobals.INDICATOR_MOVEMENT,
	)

	// // Jump.
	scenesUtil.AddOnCollisionIndicatorEntity(
		scene,
		462,
		275,
		50,
		10,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_JUMP].X,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_JUMP].Y,
		true,
		UIGlobals.INDICATOR_JUMP,
	)

	// // Ladder.
	scenesUtil.AddOnCollisionIndicatorEntity(
		scene,
		1120,
		187,
		10,
		10,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_LADDER].X,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_LADDER].Y,
		true,
		UIGlobals.INDICATOR_LADDER,
	)

	// Transition Entity Door.
	scenesUtil.AddSceneTransitionEntity(
		scene,
		3661,
		250,
		60,
		110,
		LevelTwoScene{},
		66,
		231,
		0,
		0,
	)

	// Thugs.
	scenesUtil.AddNpcEntity(scene, 1071, 188, npcGlobals.NPC_NAME_BIG_BOI, true)
	scenesUtil.AddNpcEntity(scene, 1636, 45, npcGlobals.NPC_NAME_BIG_BOI, true)
	scenesUtil.AddNpcEntity(scene, 2095, 149, npcGlobals.NPC_NAME_BIG_BOI, true)
	scenesUtil.AddNpcEntity(scene, 2606, 159, npcGlobals.NPC_NAME_BIG_BOI, true)

	return scene
}
