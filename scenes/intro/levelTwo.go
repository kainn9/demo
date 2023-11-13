package introScenes

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	scenesUtil "github.com/kainn9/demo/scenes/util"
	"github.com/kainn9/demo/systems/systemInitializers"
)

type LevelTwoScene struct{}

const (
	LEVEL_TWO_SCENE_WIDTH      = 1229
	LEVEL_TWO_SCENE_HEIGHT     = 360
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
	systemInitializers.InitStandardSystems(scene, true)

	// Entities ----------------------------------------------------------------------------------
	scenesUtil.AddCameraEntity(scene, 0, 0)

	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		components.NewParallaxLayerConfig(LEVEL_TWO_SCENE_ASSET_PATH, 0, 0, 0, false),
	})

	// Floor.
	scenesUtil.AddFloorEntity(scene, float64(scene.Width/2), float64(scene.Height-20), float64(scene.Width), 142, 0)

	// Interactables.

	// Chat.
	content := []components.SlidesContent{
		{
			Text:         "Lorum yolo bolo polo",
			PortraitName: "player",
		},
		{
			Text:         "Ipsum wipsum bipsom",
			PortraitName: "bigBoi",
		},
	}

	scenesUtil.AddOnInteractChatEntity(
		scene,
		"introChat",
		content,
		172, 231, 60, 50,
	)

	// Transition Entities Door.
	// Back Outside.
	scenesUtil.AddSceneTransitionEntity(
		scene,
		57,
		209,
		60,
		110,
		LevelOneScene{},
		3654,
		275,
		3654-float64(clientGlobals.SCREEN_WIDTH/2),
		0,
	)

	// Into Room.
	scenesUtil.AddSceneTransitionEntity(
		scene,
		1171,
		209,
		60,
		110,
		LevelThreeScene{},
		66,
		231,
		0,
		0,
	)

	return scene
}
