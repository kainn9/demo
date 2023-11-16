package introScenes

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
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
	scenesUtil.AddCameraEntity(scene, 0, 0, 2)

	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		components.NewParallaxLayerConfig(LEVEL_TWO_SCENE_ASSET_PATH, 0, 0, 0, false),
	})

	scenesUtil.AddFrontLayerEntity(scene, LEVEL_TWO_SCENE_ASSET_PATH)

	// Floor.
	scenesUtil.AddBlockEntity(scene, 319, 366, LEVEL_ONE_SCENE_WIDTH*2, 31, 0)

	// Interactables.

	// Chat.
	content := []components.SlidesContent{
		{
			Text:         "Its a flier for some kind of CBT Therapy.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_NAME,
			FacingRight:  true,
		},
	}

	scenesUtil.AddOnInteractChatEntity(
		scene,
		"introChat",
		content,
		175, 276, 81, 110,
	)

	// Transition Entities Door.
	// Back Outside.
	scenesUtil.AddSceneTransitionEntity(
		scene,
		57,
		275,
		60,
		110,
		LevelOneScene{},
		3622,
		275,
		3302,
		0,
	)

	// Into Room.
	scenesUtil.AddSceneTransitionEntity(
		scene,
		1171,
		275,
		60,
		110,
		LevelThreeScene{},
		94, 313, -160, 90,
	)

	return scene
}
