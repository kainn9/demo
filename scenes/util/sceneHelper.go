package scenesUtil

import (
	"log"

	"github.com/kainn9/coldBrew"
	assetComponents "github.com/kainn9/demo/components/assets"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
)

func InitFirstScene(
	manager *coldBrew.Manager,
	firstScene coldBrew.SceneFace,
	playerX, playerY float64,
) {

	err := manager.LoadScene(firstScene)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	currentScene := manager.ActiveScene()
	AddPlayerEntity(currentScene, playerX, playerY)
	AddUISpritesSingletonEntity(currentScene)

}

func ChangeScene(
	manager *coldBrew.Manager,
	newScene coldBrew.SceneFace,
	playerX, playerY float64,
) {

	// Get old scene.
	prevScene := manager.ActiveScene()
	prevPlayerEntity := systemsUtil.GetPlayerEntity(prevScene.World)
	prevUISingletonEntity := systemsUtil.GetUISingletonEntity(prevScene.World)

	// Get new scene.
	err := manager.LoadScene(newScene)
	if err != nil {
		log.Println("Yo:", err)
		return
	}

	currentScene := manager.ActiveScene()

	// Transfer necessary player components(state) to new scene.
	transferPlayer(prevPlayerEntity, prevScene, currentScene, playerX, playerY)
	transferUISingleton(prevUISingletonEntity, prevScene, currentScene)

}

func transferPlayer(
	prevPlayerEntity *donburi.Entry,
	prevScene, newScene *coldBrew.Scene,
	playerX, playerY float64,
) {

	AddPlayerEntity(newScene, playerX, playerY)
	newPlayerEntity := systemsUtil.GetPlayerEntity(newScene.World)

	// Transfer sprites to avoid asset reloading.
	oldPlayerSpriteMap := assetComponents.PlayerSpritesAnimMapComponent.Get(prevPlayerEntity)

	newPlayerSpriteMap := assetComponents.PlayerSpritesAnimMapComponent.Get(newPlayerEntity)

	for key, value := range *oldPlayerSpriteMap {
		(*newPlayerSpriteMap)[key] = value
	}

	// Remove player from old scene.
	prevScene.World.Remove(prevPlayerEntity.Entity())

}

func transferUISingleton(
	prevUISingletonEntity *donburi.Entry,
	prevScene, newScene *coldBrew.Scene,
) {

	prevSpritesMap := assetComponents.SpritesMapComponent.Get(prevUISingletonEntity)

	AddUISpritesSingletonEntity(newScene)
	newUISingletonEntity := systemsUtil.GetUISingletonEntity(newScene.World)

	newSpritesMap := assetComponents.SpritesMapComponent.Get(newUISingletonEntity)

	// Transfer sprites to avoid asset reloading.
	for key, value := range *prevSpritesMap {
		(*newSpritesMap)[key] = value
	}

	// Remove UISingleton from old scene.
	prevScene.World.Remove(prevUISingletonEntity.Entity())

}
