package introScenes

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	inventoryUtil "github.com/kainn9/demo/systems/sim/util/inventory"
	"github.com/kainn9/demo/systems/systemInitializers"
	systemsUtil "github.com/kainn9/demo/systems/util"
	callbacksUtil "github.com/kainn9/demo/systems/util/callbacks"
	"github.com/yohamta/donburi"

	UIGlobals "github.com/kainn9/demo/globalConfig/UI"
	inventoryGlobals "github.com/kainn9/demo/globalConfig/inventory"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	scenesUtil "github.com/kainn9/demo/scenes/util"
)

type LevelOneScene struct{}

const (
	LEVEL_ONE_SCENE_WIDTH      = 4030
	LEVEL_ONE_SCENE_HEIGHT     = 360
	LEVEL_ONE_SCENE_NAME       = "levelOne"
	LEVEL_ONE_SCENE_SECTION    = "intro"
	LEVEL_ONE_SCENE_ASSET_PATH = LEVEL_ONE_SCENE_SECTION + "/" + LEVEL_ONE_SCENE_NAME + "/"
)

var (
	levelOneFirstMobDefeated         = false
	levelOneFirstMobDefeatedChatName = "firstThugDefeat"
)

func (LevelOneScene) Index() string {
	return LEVEL_ONE_SCENE_NAME
}

func (LevelOneScene) New(m *coldBrew.Manager) *coldBrew.Scene {

	scene := coldBrew.NewScene(m, LEVEL_ONE_SCENE_WIDTH, LEVEL_ONE_SCENE_HEIGHT)

	// Systems ----------------------------------------------------------------------------------
	systemInitializers.InitStandardSystems(scene, "The Outskirts.", false)

	// Entities ----------------------------------------------------------------------------------
	scenesUtil.AddCameraEntity(scene, 0, 0, 1)

	// Todo: consider saving duration in a global config thingy
	scenesUtil.AddBgSoundEntity(scene, LEVEL_ONE_SCENE_ASSET_PATH, 1800)

	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		// Sky.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 0, 28, 0, false),

		// Moon
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 1, 0, 0, true),

		// Clouds Back.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 2, 18, 0, false),

		// Clouds Front.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 3, 6, 0, false),

		// City.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 4, 4, 0, false),

		// Front.
		components.NewParallaxLayerConfig(LEVEL_ONE_SCENE_ASSET_PATH, 5, 0, 0, false),
	})

	// Walls.
	scenesUtil.AddWalls(scene, LEVEL_ONE_SCENE_WIDTH, LEVEL_ONE_SCENE_HEIGHT)

	// Floors left to right.
	scenesUtil.AddBlockEntity(scene, 255, float64(scene.Height), 514, 95, 0)
	scenesUtil.AddBlockEntity(scene, 937, float64(scene.Height), 515, 95, 0)
	scenesUtil.AddBlockEntity(scene, 3445, float64(scene.Height), 1180, 95, 0)

	scenesUtil.AddBlockEntity(scene, 994, 290, 283, 130, 0)

	// Ladder.
	scenesUtil.AddLadderEntity(scene, 1130, 180, 40, 360)

	// Platforms left to right.
	scenesUtil.AddPlatformEntity(scene, 1491, 175, 322, 10)
	scenesUtil.AddPlatformEntity(scene, 2360, 191, 962, 10)

	// Tutorial indicators.
	// Movement.
	scenesUtil.AddOnCollisionIndicatorEntity(
		scene,
		147,
		275,
		10,
		15,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_MOVEMENT].X,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_MOVEMENT].Y,
		true,
		UIGlobals.INDICATOR_MOVEMENT,
	)

	// // Jump.
	scenesUtil.AddOnCollisionIndicatorEntity(
		scene,
		462,
		275,
		50,
		10,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_JUMP].X,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_JUMP].Y,
		true,
		UIGlobals.INDICATOR_JUMP,
	)

	// // Ladder.
	scenesUtil.AddOnCollisionIndicatorEntity(
		scene,
		1120,
		187,
		10,
		10,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_LADDER].X,
		UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_LADDER].Y,
		true,
		UIGlobals.INDICATOR_LADDER,
	)

	// Auto Chat.

	introLevelOneGetClinicContent := []components.SlidesContent{
		{
			Text:         "Get to the clinic. I need to figure this out.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_BAD_NAME,
			FacingRight:  true,
		},
	}

	scenesUtil.AddBasicChatEntity(scene, "introLevelOneGetClinic", introLevelOneGetClinicContent, true)

	// Thugs.
	groupX := 3052.0
	thugs := []*donburi.Entry{
		scenesUtil.NpcEntityFactory.AddNpcThug(scene, 1071, 188, 861, 1125),
		scenesUtil.NpcEntityFactory.AddNpcThug(scene, 1636, 45, 1350, 1637),
		scenesUtil.NpcEntityFactory.AddNpcThug(scene, 2095, 149, 1897, 2835),
		scenesUtil.NpcEntityFactory.AddNpcThug(scene, 2606, 159, 1897, 2835),

		scenesUtil.NpcEntityFactory.AddNpcThug(scene, groupX, 275, 2865, 4017),
		scenesUtil.NpcEntityFactory.AddNpcThug(scene, groupX+100, 275, 2865, 4017),
		scenesUtil.NpcEntityFactory.AddNpcThug(scene, groupX+200, 275, 2865, 4017),
		scenesUtil.NpcEntityFactory.AddNpcThug(scene, groupX+300, 275, 2865, 4017),
	}

	for _, npc := range thugs {
		callbacksUtil.AttachNpcDefeatCallback(scene, LevelOneThugsUniqueDropCallback{npc: npc})
	}

	firstThugDefeatChat := []components.SlidesContent{
		{
			Text:         "HA! Deadzo Mcgee.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_BAD_NAME,
			FacingRight:  true,
		},
	}

	scenesUtil.AddBasicChatEntity(scene, levelOneFirstMobDefeatedChatName, firstThugDefeatChat, false)

	// Transition Entity Door.
	zapClinicRestrictionChatContent := []components.SlidesContent{
		{
			Text:         "I can't until I hurt.",
			PortraitName: playerGlobals.PLAYER_PORTRAIT_INDEX,
			CharName:     playerGlobals.PLAYER_BAD_NAME,
			FacingRight:  true,
		},
	}

	zapClinicRestrictionChatEntity := scenesUtil.AddBasicChatEntity(scene, "zapClinicRestrictionChat", zapClinicRestrictionChatContent, false)

	scenesUtil.AddSceneTransitionEntity(
		scene,
		3661,
		250,
		60,
		110,
		LevelTwoScene{},
		96, 313, -160, 90,
	)

	callbacksUtil.AttachTransitionCallback(
		scene,
		ZapClinicDoorRequirements{
			chatEntity: zapClinicRestrictionChatEntity,
		},
	)

	return scene
}

// Unique Drop Callbacks.
type LevelOneThugsUniqueDropCallback struct {
	npc *donburi.Entry
}

func (cb LevelOneThugsUniqueDropCallback) Npc() *donburi.Entry {
	return cb.npc
}

func (cb LevelOneThugsUniqueDropCallback) OnDefeat(scene *coldBrew.Scene, npcEntity *donburi.Entry) {
	world := scene.World
	playerEntity := systemsUtil.PlayerEntity(world)

	inventory := components.InventoryComponent.Get(playerEntity)
	itemToAdd := components.NewInventoryItem(inventoryGlobals.ITEM_NAME_ZAP_CLINIC_UNLOCK, 1)
	inventoryUtil.AddItemToInventory(inventory, itemToAdd)

	queries.ChatQuery.Each(world, func(chatEntity *donburi.Entry) {
		stateAndConfig := components.ChatStateAndConfigComponent.Get(chatEntity)
		config := stateAndConfig.Config

		chatName := config.ChatName

		if levelOneFirstMobDefeatedChatName == chatName && !levelOneFirstMobDefeated {
			levelOneFirstMobDefeated = true
			stateAndConfig.Enable()
		}

	})
}

// Transition Permission Callbacks.
type ZapClinicDoorRequirements struct {
	chatEntity *donburi.Entry
}

func (ZapClinicDoorRequirements) AllowedToTransition(scene *coldBrew.Scene) bool {

	playerEntity := systemsUtil.PlayerEntity(scene.World)
	inventory := components.InventoryComponent.Get(playerEntity)

	doorKey := inventoryUtil.GetItemFromInventory(*inventory, inventoryGlobals.ITEM_NAME_ZAP_CLINIC_UNLOCK)

	return doorKey != nil
}

func (cb ZapClinicDoorRequirements) ChatEntity() *donburi.Entry {
	return cb.chatEntity
}

func (req ZapClinicDoorRequirements) Index() string {
	return LevelTwoScene{}.Index()
}
