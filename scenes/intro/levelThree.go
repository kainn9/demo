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
	systemInitializers.InitStandardSystems(scene, true)

	// Entities ----------------------------------------------------------------------------------
	scenesUtil.AddCameraEntity(scene, 0, 0)

	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		components.NewParallaxLayerConfig(LEVEL_THREE_SCENE_ASSET_PATH, 0, 0, 0, false),
	})

	scenesUtil.AddFloorEntity(scene, float64(scene.Width/2), float64(scene.Height-20), float64(scene.Width), 142, 0)

	return scene
}
