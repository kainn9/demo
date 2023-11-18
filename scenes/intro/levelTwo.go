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
	systemInitializers.InitStandardSystems(scene, "The Zap Clinic.", true)

	// Entities ----------------------------------------------------------------------------------
	scenesUtil.AddCameraEntity(scene, 0, 0, 2)

	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		components.NewParallaxLayerConfig(LEVEL_TWO_SCENE_ASSET_PATH, 0, 0, 0, false),
	})

	scenesUtil.AddFrontLayerEntity(scene, LEVEL_TWO_SCENE_ASSET_PATH)

	// Floor.
	scenesUtil.AddBlockEntity(scene, 319, 366, LEVEL_ONE_SCENE_WIDTH*2, 31, 0)

	// Interactables.

	// Backpacks.
	levelOneBackpacksChatContent := []components.SlidesContent{
		{
			Text:         "What is this? Homeroom? I remember when she'd show up late. Hair all drippy and shampooy.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
	}

	scenesUtil.AddOnInteractChatEntity(
		scene,
		"levelOneBackpacksChat",
		levelOneBackpacksChatContent,
		318, 287,
		88, 132,
		-10, -60,
	)

	// Landscape.
	levelOneLandscapeChatContent := []components.SlidesContent{
		{
			Text:         "The woods decay, the woods decay and fall.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
	}

	scenesUtil.AddOnInteractChatEntity(
		scene,
		"levelOneLandscapeChat",
		levelOneLandscapeChatContent,
		480, 291,
		55, 135,
		-13, -60,
	)

	// Book.
	levelOneBookContent := []components.SlidesContent{
		{
			Text:         "Two roads diverged in a yellow wood.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
	}

	scenesUtil.AddOnInteractChatEntity(
		scene,
		"levelOneBookChat",
		levelOneBookContent,
		698, 320,
		42, 111,
		-14, -50,
	)

	// Radiator.
	levelOneRadiatorContent := []components.SlidesContent{
		{
			Text:         "Dad's house had these. At mom's we'd use extra blankets.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
	}

	scenesUtil.AddOnInteractChatEntity(
		scene,
		"levelOneRadiatorChat",
		levelOneRadiatorContent,
		813, 330,
		53, 30,
		-12, -40,
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
