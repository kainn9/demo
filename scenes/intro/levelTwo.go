package intro

import (
	"github.com/kainn9/coldBrew"
	assetComponents "github.com/kainn9/demo/components/assets"
	scenesUtil "github.com/kainn9/demo/scenes/util"
)

type LevelTwoScene struct{}

const (
	LEVEL_TWO_SCENE_WIDTH      = 830
	LEVEL_TWO_SCENE_HEIGHT     = 620
	LEVEL_TWO_SCENE_NAME       = "levelTwo"
	LEVEL_TWO_SCENE_SECTION    = "intro"
	LEVEL_TWO_SCENE_ASSET_PATH = LEVEL_TWO_SCENE_SECTION + "/" + LEVEL_TWO_SCENE_NAME + "/"
)

func (LevelTwoScene) Index() string {
	return LEVEL_TWO_SCENE_NAME
}

func (LevelTwoScene) New(m *coldBrew.Manager) *coldBrew.Scene {

	scene := coldBrew.NewScene(m, LEVEL_TWO_SCENE_WIDTH, LEVEL_TWO_SCENE_HEIGHT)

	// Systems ----------------------------------------------------------------------------------
	scenesUtil.InitStandardSystems(scene)

	// Entities ----------------------------------------------------------------------------------
	scenesUtil.AddParallaxBackgroundEntity(scene, []*assetComponents.ParallaxLayerConfig{
		// Background.
		assetComponents.NewParallaxLayerConfig(LEVEL_TWO_SCENE_ASSET_PATH, 0, 8, 0, false),

		// Front.
		assetComponents.NewParallaxLayerConfig(LEVEL_TWO_SCENE_ASSET_PATH, 1, 0, 0, false),
	})

	scenesUtil.AddPlayerEntity(scene, 100, 0)
	scenesUtil.AddCameraEntity(scene, 0, 0)

	// Top(Todo: replace with platform).
	scenesUtil.AddFloorEntity(scene, float64(scene.Width/2), float64(scene.Height-528), float64(scene.Width), 15, 0)

	// Bottom.
	scenesUtil.AddFloorEntity(scene, float64(scene.Width/2), float64(scene.Height-20), float64(scene.Width), 20, 0)

	return scene
}
