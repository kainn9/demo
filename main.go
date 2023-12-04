package main

import (
	"fmt"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	UIScenes "github.com/kainn9/demo/scenes/UI"
	introScenes "github.com/kainn9/demo/scenes/intro"
	scenesUtil "github.com/kainn9/demo/scenes/util"
	loaderUtil "github.com/kainn9/demo/systems/loader/util"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

type game struct {
	manager *coldBrew.Manager
}

func main() {
	game := NewGame()

	ebiten.SetWindowTitle("Demo!")
	windowResizingMode := ebiten.WindowResizingModeEnabled
	ebiten.SetWindowResizingMode(windowResizingMode)

	ebiten.SetTPS(60)
	ebiten.SetVsyncEnabled(false) // Experimenting.

	ebiten.RunGame(game)
}

func NewGame() *game {

	// This is pretty troll, but no big deal for now.
	loaderSprite := components.NewSprite(0, 0)
	loaderSprite.Image = ebiten.NewImage(clientGlobals.SCREEN_WIDTH, clientGlobals.SCREEN_HEIGHT)
	loaderUtil.LoadImage(clientGlobals.UI_ASSETS_SUB_PATH+"loader", loaderSprite)

	manager := coldBrew.NewManager(clientGlobals.SCENE_CACHE_LIMIT, clientGlobals.MAX_TICKS, loaderSprite.Image)

	firstScene := UIScenes.TitleScene{}

	scenesUtil.InitFirstScene(manager, firstScene, 0, 0)

	g := &game{
		manager: manager,
	}

	return g

}

func (g *game) Update() error {

	toggleDebugMode()

	// Temp hack/test. ------------------------
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		scenesUtil.ChangeScene(g.manager, introScenes.LevelOneScene{}, 147, 275, 0, 0)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		scenesUtil.ChangeScene(g.manager, introScenes.LevelTwoScene{}, 96, 313, -160, 90)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		scenesUtil.ChangeScene(g.manager, introScenes.LevelThreeScene{}, 94, 313, -160, 90)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key4) {
		scenesUtil.ChangeScene(g.manager, UIScenes.TitleScene{}, -500, -500, 0, 0)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key5) {
		world := g.manager.ActiveScene().World
		playerEntity := systemsUtil.PlayerEntity(world)
		playerBody := components.RigidBodyComponent.Get(playerEntity)

		cameraEntity := systemsUtil.CameraEntity(world)
		camera := components.CameraComponent.Get(cameraEntity)

		log.Println("cameraPos", camera.X, camera.Y)
		log.Println("playerPos", playerBody.Pos.X, playerBody.Pos.Y)
	}

	// end of hack/test. ----------------------

	activeScene := g.manager.ActiveScene()

	go activeScene.Load()

	activeScene.Sync()

	deltaTime := (0.017)
	activeScene.Sim(deltaTime)

	g.manager.TickHandler.IncrementTick()

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {

	activeScene := g.manager.ActiveScene()
	activeScene.Draw(screen)
	renderDebugInfo(screen)

}

func (g *game) Layout(w, h int) (int, int) {
	return clientGlobals.SCREEN_WIDTH, clientGlobals.SCREEN_HEIGHT
}

func renderDebugInfo(screen *ebiten.Image) {

	if clientGlobals.DEBUG_MODE == false {
		return
	}

	tps := math.Round(ebiten.ActualTPS())
	fps := math.Round(ebiten.ActualFPS())

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %v", tps))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nFPS: %v\nDebug Mode Currently On.\nPress 0 to toggle on/off.", fps))

}

func toggleDebugMode() {
	if inpututil.IsKeyJustPressed(ebiten.Key0) {
		log.Println("Toggling debug mode.")
		clientGlobals.DEBUG_MODE = !clientGlobals.DEBUG_MODE
	}
}
