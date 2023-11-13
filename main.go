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

	ebiten.RunGame(game)
}

func NewGame() *game {

	// This is pretty troll, but no big deal for now.
	loaderSprite := components.NewSprite(0, 0)
	loaderSprite.Image = ebiten.NewImage(clientGlobals.SCREEN_WIDTH, clientGlobals.SCREEN_HEIGHT)
	loaderUtil.LoadImage(clientGlobals.UI_ASSETS_SUB_PATH+"loader", loaderSprite)

	manager := coldBrew.NewManager(clientGlobals.SCENE_CACHE_LIMIT, clientGlobals.MAX_TICKS, loaderSprite.Image)

	firstScene := introScenes.LevelOneScene{}
	scenesUtil.InitFirstScene(manager, firstScene, 147, 275)

	g := &game{
		manager: manager,
	}

	ebiten.SetWindowTitle("Demo!")
	windowResizingMode := ebiten.WindowResizingModeEnabled
	ebiten.SetWindowResizingMode(windowResizingMode)
	ebiten.SetVsyncEnabled(true) // Experimental.
	// ebiten.SetFullscreen(true)

	return g

}

func (g *game) Update() error {
	toggleDebugMode()

	// Temp hack/test. ------------------------
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		scenesUtil.ChangeScene(g.manager, introScenes.LevelOneScene{}, 147, 275, 0, 0)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		scenesUtil.ChangeScene(g.manager, introScenes.LevelTwoScene{}, 66, 231, 0, 0)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		scenesUtil.ChangeScene(g.manager, introScenes.LevelThreeScene{}, 66, 231, 0, 0)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key5) {
		world := g.manager.ActiveScene().World
		playerEntity := systemsUtil.GetPlayerEntity(world)
		playerBody := components.RigidBodyComponent.Get(playerEntity)

		cameraEntity := systemsUtil.GetCameraEntity(world)
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

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {

	activeScene := g.manager.ActiveScene()
	activeScene.Draw(screen)
	g.manager.TickHandler.IncrementTick()

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
