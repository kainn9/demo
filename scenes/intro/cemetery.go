package introScenes

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/systems/systemInitializers"

	scenesUtil "github.com/kainn9/demo/scenes/util"
)

type CemeteryScene struct{}

const (
	CEMETERY_SCENE_WIDTH      = 1630
	CEMETERY_SCENE_HEIGHT     = 680
	CEMETERY_SCENE_NAME       = "cemetery"
	CEMETERY_SCENE_SECTION    = "eastCity"
	CEMETERY_SCENE_ASSET_PATH = CEMETERY_SCENE_SECTION + "/" + CEMETERY_SCENE_NAME + "/"
)

func (CemeteryScene) Index() string {
	return CEMETERY_SCENE_NAME
}

func (CemeteryScene) New(m *coldBrew.Manager) *coldBrew.Scene {

	scene := coldBrew.NewScene(m, CEMETERY_SCENE_WIDTH, CEMETERY_SCENE_HEIGHT)

	// Systems ----------------------------------------------------------------------------------
	systemInitializers.InitStandardSystems(scene, "East City - Cemetery.", false)

	// Entities ----------------------------------------------------------------------------------
	scenesUtil.AddSceneData(scene, CEMETERY_SCENE_ASSET_PATH)

	scenesUtil.AddCameraEntity(scene, 0, 0, 1)

	// scenesUtil.AddBgSoundEntity(scene, CEMETERY_SCENE_ASSET_PATH)

	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		// Sky.
		components.NewParallaxLayerConfig(CEMETERY_SCENE_ASSET_PATH, 0, 0, 0, false),

		// City far.
		components.NewParallaxLayerConfig(CEMETERY_SCENE_ASSET_PATH, 1, 8, 0, false),

		// Trees Back.
		components.NewParallaxLayerConfig(CEMETERY_SCENE_ASSET_PATH, 2, 4, 0, false),

		// Main Level.
		components.NewParallaxLayerConfig(CEMETERY_SCENE_ASSET_PATH, 3, 0, 0, false),
	})

	scenesUtil.AddFrontLayerEntity(scene, CEMETERY_SCENE_ASSET_PATH)

	// Walls.
	scenesUtil.AddWalls(scene, CEMETERY_SCENE_WIDTH, CEMETERY_SCENE_HEIGHT)

	return scene
}
