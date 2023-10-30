package loaderSystems

import (
	"log"
	"strconv"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
	"github.com/kainn9/demo/queries"
	"github.com/kainn9/demo/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type ChatLoaderSystem struct {
	scene *coldBrew.Scene
}

func NewChatLoader(scene *coldBrew.Scene) *ChatLoaderSystem {
	return &ChatLoaderSystem{
		scene: scene,
	}
}

func (sys *ChatLoaderSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.ChatStateComponent),
	)
}

func (sys *ChatLoaderSystem) Load(entity *donburi.Entry) {
	loadChatPreReqAssets(sys.scene)

	config := components.ChatStateComponent.Get(entity)
	slideSprites := assetComponents.SpritesSliceComponent.Get(entity)
	loadChatSprites(config, *slideSprites, config.SceneAssetsPath)

	portraitsSpriteMap := assetComponents.SpritesMapComponent.Get(entity)
	loadPortraitSprites(config, config.PortraitNames, *portraitsSpriteMap, config.SceneAssetsPath)

}

func loadChatPreReqAssets(scene *coldBrew.Scene) {

	world := scene.World

	chatAnimationStateQueries := []*donburi.Query{
		queries.ChatPopUpEntityQuery,
		queries.ChatPopDownEntityQuery,
	}

	chatTags := []*donburi.ComponentType[struct{}]{
		tags.ChatPopUpSpriteTag,
		tags.ChatPopDownSpriteTag,
	}

	animationStateNames := []string{
		constants.CHAT_STATE_POP_UP,
		constants.CHAT_STATE_POP_DOWN,
	}

	for i, query := range chatAnimationStateQueries {

		if query.Count(world) > 0 {
			return
		}

		log.Println("Loading Chat PreReq Assets.")

		tag := chatTags[i]

		entity := scene.AddEntity(
			assetComponents.SpriteComponent,
			tag,
		)

		// Flow is a little weird here...
		// First we need to register the sprite,
		// using a value type...
		assetComponents.SpriteComponent.SetValue(
			entity,
			*assetComponents.NewSprite(0, 0),
		)
		// And now we can get the pointer to the sprite, for the good stuff.
		spriteComponent := assetComponents.SpriteComponent.Get(entity)

		// Were gonna load the asset before creating the animation data.
		// So we can use the asset data to help create the animation data.
		path := constants.UI_ASSETS_SUB_PATH + animationStateNames[i]
		LoadImage(path, spriteComponent)

		// Now we can create the animation data.
		frameWidth, frameHeight, frameCount := getAnimData(spriteComponent)
		spriteComponent.AnimationConfig = assetComponents.NewAnimationConfig(frameWidth, frameHeight, frameCount, constants.CHAT_ANIM_POP_UP_SPEED, true)
	}

}

func getAnimData(spriteComponent *assetComponents.Sprite) (frameWidth, frameHeight, frameCount int) {
	totalFrameWidth := spriteComponent.Image.Bounds().Size().X

	frameWidth = constants.CHAT_FRAME_WIDTH
	frameHeight = spriteComponent.Image.Bounds().Size().Y
	frameCount = totalFrameWidth / constants.CHAT_FRAME_WIDTH

	return frameWidth, frameHeight, frameCount
}

func loadChatSprites(config *components.ChatState, spritesSlice []*assetComponents.Sprite, sceneAssetsPath string) {

	for i, slideSprite := range spritesSlice {

		if slideSprite.AssetData.Loaded {
			continue
		}

		path := constants.SCENE_ASSETS_SUB_PATH
		path += sceneAssetsPath
		path += constants.CHAT_SCENE_ASSETS_SUB_PATH
		path += config.ChatName + "/"
		path += strconv.Itoa(i)

		log.Println("Loading Chat Sprite for", path, "at index", i, ".")

		LoadImage(path, slideSprite)

		// Now we can create the animation data from the loaded asset data.
		frameWidth, frameHeight, frameCount := getAnimData(slideSprite)
		slideSprite.AnimationConfig = assetComponents.NewAnimationConfig(frameWidth, frameHeight, frameCount, constants.CHAT_ANIM_TEXT_SPEED, true)
	}
}

// This setup is not very efficient, as it can load the same portrait multiple times.
// But it's fine for now.
func loadPortraitSprites(config *components.ChatState, portraitNames []string, portraitsSpriteMap map[string]*assetComponents.Sprite, sceneAssetsPath string) {
	for _, portraitName := range portraitNames {

		path := constants.CHARACTER_ASSETS_SUB_PATH
		path += portraitName + "/" + "portrait"

		log.Println("Loading Portrait Sprite for", path, ".")
		LoadImage(path, portraitsSpriteMap[portraitName])

	}
}
