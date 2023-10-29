package intro

import (
	"github.com/kainn9/coldBrew"
	assetComponents "github.com/kainn9/demo/components/assets"
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
	scenesUtil.InitStandardSystems(scene)

	// Entities ----------------------------------------------------------------------------------

	scenesUtil.AddParallaxBackgroundEntity(scene, []*assetComponents.ParallaxLayerConfig{
		// Sky.
		assetComponents.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 0, 22, 22, false),

		// City Far Shadow.
		assetComponents.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 1, 18, 18, false),

		// City Lights.
		assetComponents.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 2, 8, 8, false),

		// Mountains
		assetComponents.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 3, 0, 0, false),

		// Green Trees.
		assetComponents.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 4, 10, 10, false),

		// Statues.
		assetComponents.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 5, 5, 12, false),

		// Red Shrubs.
		assetComponents.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 6, 4, 8, false),

		// Gears Close.
		assetComponents.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 7, 2, 0, false),
		// Main layer.
		assetComponents.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 8, 0, 0, false),
	})

	scenesUtil.AddPlayerEntity(scene, 100, 600)
	scenesUtil.AddCameraEntity(scene, 0, 262)

	scenesUtil.AddFloorEntity(scene, float64(scene.Width/2), float64(scene.Height-40), float64(scene.Width), 20, -0.04)
	scenesUtil.AddChatEntity(
		scene,
		3,
		"introChat",
		LEVEL_ONE_SCENE_ASSET_PATH,
		[]string{
			"player",
			"player",
			"player",
		},
	)

	return scene
}
