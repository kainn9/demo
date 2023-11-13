package introScenes

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"
	"github.com/kainn9/demo/queries"
	scenesUtil "github.com/kainn9/demo/scenes/util"
	"github.com/kainn9/demo/systems/systemInitializers"
	"github.com/yohamta/donburi"
)

type LevelThreeScene struct{}

const (
	LEVEL_THREE_SCENE_WIDTH      = 650
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

	// Background.
	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		components.NewParallaxLayerConfig(LEVEL_THREE_SCENE_ASSET_PATH, 0, 0, 0, false),
	})

	// Floor.
	scenesUtil.AddBlockEntity(scene, float64(scene.Width/2), float64(scene.Height-20), float64(scene.Width), 142, 0)

	// Bookshelf.
	scenesUtil.AddBlockEntity(scene, 513, 212, 91, 101, 0)

	// Chat.
	content := []components.SlidesContent{
		{
			Text:         "Ligma.",
			PortraitName: "therapistOne",
		},
		{
			Text:         "Ligma what?",
			PortraitName: "player",
			FacingRight:  true,
		},
		{
			Text:         "Ligma balls.",
			PortraitName: "therapistTwo",
		},
	}

	scenesUtil.AddOnCollideChatEntity(
		scene,
		"introChat",
		content,
		350, 231, 100, 50,
	)

	// Into Hallway.
	scenesUtil.AddSceneTransitionEntity(
		scene,
		56,
		208,
		60,
		110,
		LevelTwoScene{},
		1167,
		231,
		749,
		50,
	)

	// Off scene(gets moved later).
	gravityMod := components.NewPhysicsConfig(0.25)
	scenesUtil.AddNpcEntity(scene, -200, -200, npcGlobals.NPC_NAME_THERAPIST_TWO, gravityMod, false)

	// Attaching unique chat callback.
	systemInitializers.AttachChatCallback(scene, LevelThreeCallbackSystem{})

	return scene
}

type LevelThreeCallbackSystem struct{}

func (LevelThreeCallbackSystem) ChatName() string {
	return "introChat"
}
func (LevelThreeCallbackSystem) SlideIndex() int {
	return 1
}

func (LevelThreeCallbackSystem) Callback(scene *coldBrew.Scene) {
	query := queries.NpcQuery

	query.Each(scene.World, func(entity *donburi.Entry) {

		npcBody := components.RigidBodyComponent.Get(entity)
		npcBody.Vel.Y = 0
		npcBody.Pos.X = 480
		npcBody.Pos.Y = -50
	})
}
