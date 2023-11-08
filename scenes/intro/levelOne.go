package introScenes

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	UIConstants "github.com/kainn9/demo/constants/UI"
	"github.com/kainn9/demo/systems/systemInitializers"

	scenesUtil "github.com/kainn9/demo/scenes/util"
)

type LevelOneScene struct{}

const (
	LEVEL_ONE_SCENE_WIDTH      = 1920
	LEVEL_ONE_SCENE_HEIGHT     = 622
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
	systemInitializers.InitStandardSystems(scene)

	// Entities ----------------------------------------------------------------------------------

	scenesUtil.AddCameraEntity(scene, 0, 262)

	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		// Sky.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 0, 22, 22, false),

		// City Far Shadow.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 1, 18, 18, false),

		// City Lights.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 2, 8, 8, false),

		// Mountains
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 3, 0, 0, false),

		// Green Trees.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 4, 8, 10, false),

		// Statues.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 5, 5, 12, false),

		// Red Shrubs.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 6, 4, 8, false),

		// Gears Close.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 7, 2, 0, false),
		// Main layer.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 8, 0, 0, false),
	})

	// Floor.
	scenesUtil.AddFloorEntity(scene, float64(scene.Width/2), float64(scene.Height-40), float64(scene.Width), 20, -0.04)

	// Platforms.
	scenesUtil.AddPlatformEntity(scene, 160, 121, float64(scene.Width), 10)                      // Left.
	scenesUtil.AddPlatformEntity(scene, float64(scene.Width)+340, 121, float64(scene.Width), 10) // Right.

	// Ladder.
	scenesUtil.AddLadderEntity(scene, 1875, 326, 20, 420)

	// Chat.
	content := []components.SlidesContent{
		{
			Text:         "Jeez, I'm so tired...I've been running for hours. Can't believe its really not out here.",
			PortraitName: "player",
		},
		{
			Text:         "*stomach growls*",
			PortraitName: "player",
		},
		{
			Text:         "And thats a problem too...I did hear the country club has a new chef...",
			PortraitName: "player",
		},
		{
			Text:         "Maybe I can sneak in and get a bite to eat. Guess It's time to grab the whip.",
			PortraitName: "player",
		},
	}

	scenesUtil.AddChatEntity(
		scene,
		"introChat",
		LEVEL_ONE_SCENE_ASSET_PATH,
		15,
		content,
	)

	// Tutorial indicators.

	// Movement.
	scenesUtil.AddOnCollisionIndicatorEntity(
		scene,
		85,
		LEVEL_ONE_SCENE_HEIGHT-40,
		10,
		15,
		UIConstants.IndicatorPlayerOffsets[UIConstants.CurrentLayout][UIConstants.INDICATOR_MOVEMENT].X,
		UIConstants.IndicatorPlayerOffsets[UIConstants.CurrentLayout][UIConstants.INDICATOR_MOVEMENT].Y,
		true,
		UIConstants.INDICATOR_MOVEMENT,
	)

	// Ladder.
	scenesUtil.AddOnCollisionIndicatorEntity(
		scene,
		1870,
		LEVEL_ONE_SCENE_HEIGHT-120,
		10,
		10,
		UIConstants.IndicatorPlayerOffsets[UIConstants.CurrentLayout][UIConstants.INDICATOR_LADDER].X,
		UIConstants.IndicatorPlayerOffsets[UIConstants.CurrentLayout][UIConstants.INDICATOR_LADDER].Y,
		true,
		UIConstants.INDICATOR_LADDER,
	)

	// Jump.
	scenesUtil.AddOnCollisionIndicatorEntity(
		scene,
		1330,
		90,
		50,
		10,
		UIConstants.IndicatorPlayerOffsets[UIConstants.CurrentLayout][UIConstants.INDICATOR_JUMP].X,
		UIConstants.IndicatorPlayerOffsets[UIConstants.CurrentLayout][UIConstants.INDICATOR_JUMP].Y,
		true,
		UIConstants.INDICATOR_JUMP,
	)

	// Door.
	scenesUtil.AddSceneTransitionEntity(
		scene,
		23,
		90,
		20,
		40,
		UIConstants.IndicatorPlayerOffsets[UIConstants.CurrentLayout][UIConstants.INDICATOR_INTERACT].X,
		UIConstants.IndicatorPlayerOffsets[UIConstants.CurrentLayout][UIConstants.INDICATOR_INTERACT].Y,
		true,
		UIConstants.INDICATOR_INTERACT,
		LevelTwoScene{},
		20,
		70,
		0,
		0,
		true,
	)

	return scene
}
