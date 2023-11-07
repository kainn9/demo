package introScenes

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	scenesUtil "github.com/kainn9/demo/scenes/util"
	"github.com/kainn9/demo/systems/systemInitializers"
)

type LevelThreeScene struct{}

const (
	LEVEL_THREE_SCENE_WIDTH      = 2185
	LEVEL_THREE_SCENE_HEIGHT     = 360
	LEVEL_THREE_SCENE_NAME       = "levelThree"
	LEVEL_THREE_SCENE_SECTION    = "intro"
	LEVEL_THREE_SCENE_ASSET_PATH = LEVEL_THREE_SCENE_SECTION + "/" + LEVEL_THREE_SCENE_NAME + "/"
)

func (LevelThreeScene) Index() string {
	return LEVEL_THREE_SCENE_NAME
}

func (LevelThreeScene) New(m *coldBrew.Manager) *coldBrew.Scene {

	scene := coldBrew.NewScene(m, LEVEL_THREE_SCENE_WIDTH, LEVEL_THREE_SCENE_HEIGHT)

	// Systems ----------------------------------------------------------------------------------
	systemInitializers.InitDrivingSystems(scene)

	// Entities ----------------------------------------------------------------------------------
	scenesUtil.AddCameraEntity(scene, 0, 0)

	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		// City behind mountains.
		components.NewParallaxLayerConfig(LEVEL_THREE_SCENE_ASSET_PATH, 0, 14, 0, false),

		// Mountains.
		components.NewParallaxLayerConfig(LEVEL_THREE_SCENE_ASSET_PATH, 1, 12, 0, false),

		// Buildings Far.
		components.NewParallaxLayerConfig(LEVEL_THREE_SCENE_ASSET_PATH, 2, 10, 0, false),

		// Buildings Close.
		components.NewParallaxLayerConfig(LEVEL_THREE_SCENE_ASSET_PATH, 3, 8, 0, false),

		// Foreground.
		components.NewParallaxLayerConfig(LEVEL_THREE_SCENE_ASSET_PATH, 4, 4, 0, false),
		// Red Shrubs.
		components.NewParallaxLayerConfig(LEVEL_THREE_SCENE_ASSET_PATH, 5, 2, 0, false),

		// Road.
		components.NewParallaxLayerConfig(LEVEL_THREE_SCENE_ASSET_PATH, 6, 0, 0, false),
	})

	// Front Layer Trees.
	scenesUtil.AddFrontLayerEntity(scene, LEVEL_THREE_SCENE_ASSET_PATH)

	scenesUtil.AddCarEntity(scene, 60, LEVEL_THREE_SCENE_HEIGHT-61)

	return scene
}
