package intro

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
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
	scenesUtil.AddCameraEntity(scene, 0, 0)

	// Chat(Temp).
	content := []components.SlidesContent{
		{
			Text:         "Oh man, I hate this place...This jump really is the worst.",
			PortraitName: "player",
		},
		{
			Text:         "Good thing there is no fall damage in this universe.",
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

	scenesUtil.AddParallaxBackgroundEntity(scene, []*assetComponents.ParallaxLayerConfig{
		assetComponents.NewParallaxLayerConfig(LEVEL_TWO_SCENE_ASSET_PATH, 0, 8, 0, false),
		assetComponents.NewParallaxLayerConfig(LEVEL_TWO_SCENE_ASSET_PATH, 1, 0, 0, false),
	})
	scenesUtil.AddFrontLayerEntity(scene, LEVEL_TWO_SCENE_ASSET_PATH)

	// Platform.
	scenesUtil.AddPlatformEntity(scene, float64(scene.Width/2), float64(scene.Height-528), float64(scene.Width), 15)

	// Floor.
	scenesUtil.AddFloorEntity(scene, float64(scene.Width/2), float64(scene.Height-20), float64(scene.Width), 20, 0)

	return scene
}
