package introScenes

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	npcGlobals "github.com/kainn9/demo/globalConfig/npc"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	"github.com/kainn9/demo/queries"
	scenesUtil "github.com/kainn9/demo/scenes/util"
	"github.com/kainn9/demo/systems/systemInitializers"
	callbacksUtil "github.com/kainn9/demo/systems/util/callbacks"
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

var mainChatOneName = "mainChatOne"
var mainChatTwoName = "mainChatTwo"

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

	// Walls.
	scenesUtil.AddWalls(scene, LEVEL_THREE_SCENE_WIDTH, LEVEL_THREE_SCENE_HEIGHT)

	// Floor.
	scenesUtil.AddBlockEntity(scene, float64(scene.Width/2), float64(scene.Height), float64(scene.Width), 20, 0)

	// Bookshelf.
	scenesUtil.AddPlatformEntity(scene, 439, 250, 91, 11)

	// Main chat one.
	mainChatOneContent := []components.SlidesContent{
		{
			Text:         "Yekshemesh!",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "Hey!",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_ONE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_ONE,
		},
		{
			Text:         "How'd you get in here?",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_ONE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_ONE,
		},
		{
			Text:         "Walked me way in.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "What about the receptionist?",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_ONE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_ONE,
		},
		{
			Text:         "Gonezo. Dipped out.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "...probably having one of her episodes.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_ONE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_ONE,
		},
		{
			Text:         "Does anyone get better 'round here?",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "I think you've made serious progress since joining the Zap Clinic.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_ONE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_ONE,
		},
		// Todo: literally should be emphasized.
		{
			Text:         "I literally just MMA'd an entire street.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "Oh.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_ONE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_ONE,
		},
		{
			Text:         "Yeah. The Bad still boomerangs back.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "What about that routine you do?",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_ONE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_ONE,
		},
		{
			Text:         "The ice baths and sauna? The meditation? The fifty mile-runs and broccolini?",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "I do it all. With a podcast in one ear and an audiobook in the other.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "Me demons still leak out.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "Hoho! Ahh. Mmm. Yeh.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_TWO,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_TWO,
		},
		{
			Text:         "How?!",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "Look man. Tight pants and a sweat aren't going to cut it.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_TWO,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_TWO,
		},
		{
			Text:         "Have some salami or a pickle.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_TWO,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_TWO,
		},
		{
			Text:         "This place isn't normal.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "Let's get you in the holy pharma family. Start you off on a few different brothers.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_TWO,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_TWO,
		},
		{
			Text:         "A few zaps, a few pills.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_TWO,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_TWO,
		},
		{
			Text:         "Heh. I have a pill right here, actually. Nicely bundled in this prosciutto.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_TWO,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_TWO,
		},
		{
			Text:         "C'mere!",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_TWO,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_TWO,
		},
		{
			Text:         "I guess...",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
	}

	scenesUtil.AddOnCollideChatEntity(
		scene,
		mainChatOneName,
		mainChatOneContent,
		330, 315, 100, 50,
	)

	// Main chat two.
	mainChatTwoContent := []components.SlidesContent{
		{

			Text:         "You're a disaster.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{

			Text:         "Mr. Goodrich.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "What the-- Why are you upside down?!",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "Good for blood flow, and I have ED.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "Listen.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "You're all binge purge, half screams half silence. A rampage into a soothing shrimp scampi, so to-speak.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "You need to confront your trauma.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "That body pile outside?",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "It's trauma lash. You lead a life of opposites.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "That charity of yours, the one for Indian cats with arthritis... Trauma lash.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "Your weekend garbage gig, the hospital volunteering... Trauma lash.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "You hurt or heal.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "It's tearing me apart.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "Yes.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "When you do everything you can for anyone at anytime.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "Then turn around to beat up six strangers...",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "This is a 50/50 soul.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "How can I fix it?",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "Go face everything bad that's ever happened to you.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "Where do I start?",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "Life's main events. The answer will be there.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "Your parent's divorce.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "The bully beat-ups and creepy babysitters.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "All the Annabel's and Brenda's that ever giggled at you.",
			PortraitName: npcGlobals.NPC_PORTRAIT_INDEX_THERAPIST_THREE,
			CharName:     npcGlobals.NPC_PORTRAIT_NAME_THERAPIST_THREE,
		},
		{
			Text:         "OK. Hundo. I have to go.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
		{
			Text:         "I feel the Bad coming back.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_GOOD_NAME,
			FacingRight:  true,
		},
	}

	scenesUtil.AddOnCollideChatEntity(
		scene,
		mainChatTwoName,
		mainChatTwoContent,
		442, 275, 93, 144,
	)

	callbacksUtil.AttachSitCallbackToChat(scene, mainChatOneName, len(mainChatOneContent))

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
	callbacksUtil.AttachChatCallback(scene, IntroChatSpawnTherapistTwoCallBack{})

	return scene
}

// Unique:
type IntroChatSpawnTherapistTwoCallBack struct{}

func (IntroChatSpawnTherapistTwoCallBack) ChatName() string {
	return mainChatOneName
}
func (IntroChatSpawnTherapistTwoCallBack) SlideIndex() int {
	return 14
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
