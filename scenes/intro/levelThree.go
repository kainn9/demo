package introScenes

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
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
	systemInitializers.InitStandardSystems(scene, "Therapists Office.", true)

	// Entities ----------------------------------------------------------------------------------
	scenesUtil.AddCameraEntity(scene, 0, 0, 2)

	// Background.
	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		components.NewParallaxLayerConfig(LEVEL_THREE_SCENE_ASSET_PATH, 0, 0, 0, false),
	})

	// Floor.
	scenesUtil.AddBlockEntity(scene, float64(scene.Width/2), float64(scene.Height), float64(scene.Width), 20, 0)

	// Bookshelf.
	scenesUtil.AddPlatformEntity(scene, 439, 250, 91, 11)

	// Chat.
	content := []components.SlidesContent{
		{
			Text:         "Lorem ipsum lala something something something...",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_ONE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_ONE,
		},
		{
			Text:         "Lorem ipsum lala something something something?",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			FacingRight:  true,
			CharName:     playerGlobals.PLAYER_NAME,
		},
		{
			Text:         "Lorem ipsum lala something something something!",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_TWO,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_TWO,
		},
		{
			Text:         "Blah blah blah blah.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_TWO,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_TWO,
		},
		{
			Text:         "Blah blah blah blah.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_TWO,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_TWO,
		},
	}

	chatName := "introChat"
	scenesUtil.AddOnCollideChatEntity(
		scene,
		chatName,
		content,
		330, 315, 100, 50,
	)

	systemInitializers.AttachSitCallbackToChat(scene, chatName, len(content))

	// Into Hallway.
	scenesUtil.AddSceneTransitionEntity(
		scene,
		56,
		295,
		60,
		110,
		LevelTwoScene{},
		1131, 313, 749, 90,
	)

	// Off scene(gets moved later).
	gravityMod := components.NewPhysicsConfig(0.25)
	scenesUtil.AddNpcEntity(scene, -200, -200, npcGlobals.NPC_NAME_THERAPIST_TWO, gravityMod, false, false)

	// Attaching unique chat callback.
	systemInitializers.AttachChatCallback(scene, IntroChatSpawnTherapistTwoCallBack{})

	return scene
}

// Unique:
type IntroChatSpawnTherapistTwoCallBack struct{}

func (IntroChatSpawnTherapistTwoCallBack) ChatName() string {
	return "introChat"
}
func (IntroChatSpawnTherapistTwoCallBack) SlideIndex() int {
	return 1
}

func (IntroChatSpawnTherapistTwoCallBack) Callback(scene *coldBrew.Scene) {
	query := queries.NpcQuery

	query.Each(scene.World, func(entity *donburi.Entry) {

		npcBody := components.RigidBodyComponent.Get(entity)
		npcBody.Vel.Y = 0
		npcBody.Pos.X = 410
		npcBody.Pos.Y = -50
	})
}
