package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	clientConstants "github.com/kainn9/demo/constants/client"
	introScenes "github.com/kainn9/demo/scenes/intro"
	scenesUtil "github.com/kainn9/demo/scenes/util"
)

type game struct {
	width, height int
	manager       *coldBrew.Manager
}

func main() {
	game := NewGame()

	ebiten.SetVsyncEnabled(true) // Experimental.

	ebiten.RunGame(game)
}

func NewGame() *game {

	loaderImage := ebiten.NewImage(clientConstants.SCREEN_WIDTH, clientConstants.SCREEN_HEIGHT)
	loaderImage.Fill(color.RGBA{B: 255})

	manager := coldBrew.NewManager(clientConstants.SCENE_CACHE_LIMIT, clientConstants.MAX_TICKS, loaderImage)

	firstScene := introScenes.LevelOneScene{}
	scenesUtil.InitFirstScene(manager, firstScene, 100, 600)

	g := &game{
		width:   clientConstants.SCREEN_WIDTH,
		height:  clientConstants.SCREEN_HEIGHT,
		manager: manager,
	}

	ebiten.SetWindowTitle("Demo!")
	ebiten.SetWindowSize(g.width*2, g.height*2)
	windowResizingMode := ebiten.WindowResizingModeEnabled
	ebiten.SetWindowResizingMode(windowResizingMode)

	return g

}

func (g *game) Update() error {
	toggleDebugMode()

	// Temp hack/test. ------------------------
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		scenesUtil.ChangeScene(g.manager, introScenes.LevelOneScene{}, 100, 600, 0, 0)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		scenesUtil.ChangeScene(g.manager, introScenes.LevelTwoScene{}, 20, 70, 0, 0)
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
	return g.width, g.height
}

func renderDebugInfo(screen *ebiten.Image) {

	if clientConstants.DEBUG_MODE == false {
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
		clientConstants.DEBUG_MODE = !clientConstants.DEBUG_MODE
	}
}
