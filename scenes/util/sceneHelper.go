package scenesUtil

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	cameraSimUtil "github.com/kainn9/demo/systems/sim/util/camera"
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
	PlayerEntityFactory.AddPlayerEntity(currentScene, playerX, playerY)

	AddUISpritesSingletonEntity(currentScene)
	AddUISoundsSingletonEntity(currentScene)

}

func ChangeScene(
	manager *coldBrew.Manager,
	newScene coldBrew.SceneFace,
	playerX, playerY, camX, camY float64,
) {

	// Get old scene.
	prevScene := manager.ActiveScene()
	prevPlayerEntity := systemsUtil.GetPlayerEntity(prevScene.World)
	prevUISpriteSingletonEntity := systemsUtil.GetUISpritesSingletonEntity(prevScene.World)
	prevUISoundsSingletonEntity := systemsUtil.GetUISoundsSingletonEntity(prevScene.World)

	// Get new scene.
	err := manager.LoadScene(newScene)
	if err != nil {
		log.Println("Yo:", err)
		return
	}

	currentScene := manager.ActiveScene()

	// Transfer necessary player components(state) to new scene.
	transferPlayer(prevPlayerEntity, prevScene, currentScene, playerX, playerY, camX, camY)
	transferUISpriteSingleton(prevUISpriteSingletonEntity, prevUISoundsSingletonEntity, prevScene, currentScene)
	pauseBgSound(prevScene)
}

func transferPlayer(
	prevPlayerEntity *donburi.Entry,
	prevScene, newScene *coldBrew.Scene,
	playerX, playerY, camX, camY float64,
) {

	PlayerEntityFactory.AddPlayerEntity(newScene, playerX, playerY)
	newPlayerEntity := systemsUtil.GetPlayerEntity(newScene.World)

	// Transfer sprites to avoid asset reloading.
	oldPlayerSpriteMap := components.SpritesCharStateMapComponent.Get(prevPlayerEntity)

	newPlayerSpriteMap := components.SpritesCharStateMapComponent.Get(newPlayerEntity)

	for key, value := range *oldPlayerSpriteMap {
		(*newPlayerSpriteMap)[key] = value
	}

	// Transfer Sound Assets to avoid reloading
	oldPlayerSoundMap := components.SoundCharStateMapComponent.Get(prevPlayerEntity)
	newPlayerSoundMap := components.SoundCharStateMapComponent.Get(newPlayerEntity)

	for key, value := range *oldPlayerSoundMap {
		(*newPlayerSoundMap)[key] = value
	}

	// Preserve player direction.
	prevPlayerState := components.PlayerStateComponent.Get(prevPlayerEntity)
	newPlayerState := components.PlayerStateComponent.Get(newPlayerEntity)

	if prevPlayerState.Direction() == -1 {
		newPlayerState.SetDirectionLeft()
	} else {
		newPlayerState.SetDirectionRight()
	}

	// Preserve inventory.
	prevInventory := components.InventoryComponent.Get(prevPlayerEntity)
	components.InventoryComponent.SetValue(newPlayerEntity, *prevInventory)

	// Remove player from old scene.
	prevScene.World.Remove(prevPlayerEntity.Entity())

	// Set Camera Position.
	currCameraEntity := systemsUtil.GetCameraEntity(newScene.World)
	currCamera := components.CameraComponent.Get(currCameraEntity)

	cameraSimUtil.SetPosition(currCamera, camX, camY, false)
}

func transferUISpriteSingleton(
	prevUISpritesSingletonEntity *donburi.Entry,
	prevUISoundsSingletonEntity *donburi.Entry,
	prevScene, newScene *coldBrew.Scene,
) {

	prevSpritesMap := components.SpritesMapComponent.Get(prevUISpritesSingletonEntity)
	prevSoundsMap := components.SoundsMapComponent.Get(prevUISoundsSingletonEntity)

	AddUISpritesSingletonEntity(newScene)
	newUISingletonEntity := systemsUtil.GetUISpritesSingletonEntity(newScene.World)

	AddUISoundsSingletonEntityWithoutContext(newScene)
	newUISoundsSingletonEntity := systemsUtil.GetUISoundsSingletonEntity(newScene.World)

	newSpritesMap := components.SpritesMapComponent.Get(newUISingletonEntity)
	newSoundsMap := components.SoundsMapComponent.Get(newUISoundsSingletonEntity)

	// Transfer sprites to avoid asset reloading.
	for key, value := range *prevSpritesMap {
		(*newSpritesMap)[key] = value
	}

	// Transfer sounds to avoid asset reloading.
	for key, value := range *prevSoundsMap {
		(*newSoundsMap)[key] = value
	}

	// Transfer AudioContext to avoid asset reloading.
	prevAudioContext := components.AudioContextComponent.Get(prevUISoundsSingletonEntity)
	components.AudioContextComponent.SetValue(newUISoundsSingletonEntity, *prevAudioContext)

	// Remove UISingletons from old scene.
	prevScene.World.Remove(prevUISpritesSingletonEntity.Entity())
	prevScene.World.Remove(prevUISoundsSingletonEntity.Entity())

}

func pauseBgSound(prevScene *coldBrew.Scene) {
	bgSoundEntity, ok := queries.BackgroundSoundQuery.First(prevScene.World)
	if !ok {
		return
	}

	bgSound := components.SoundComponent.Get(bgSoundEntity)

	bgSound.State.Player.Pause()

}
