package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	clientConstants "github.com/kainn9/demo/constants/client"
	"github.com/kainn9/demo/scenes"
	scenesUtil "github.com/kainn9/demo/scenes/util"
)

type game struct {
	width, height int
	manager       *coldBrew.Manager
}

func main() {
	game := NewGame()
	ebiten.RunGame(game)
}

func NewGame() *game {

	loaderImage := ebiten.NewImage(clientConstants.SCREEN_WIDTH, clientConstants.SCREEN_HEIGHT)
	loaderImage.Fill(color.RGBA{B: 255})

	manager := coldBrew.NewManager(clientConstants.SCENE_CACHE_LIMIT, clientConstants.MAX_TICKS, loaderImage)

	firstScene := scenes.Intro.LevelOneScene
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

	// Temp hack/test. ------------------------
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		scenesUtil.ChangeScene(g.manager, scenes.Intro.LevelOneScene, 100, 600)
	}

	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		scenesUtil.ChangeScene(g.manager, scenes.Intro.LevelTwoScene, 20, 70)
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

}

func (g *game) Layout(w, h int) (int, int) {
	return g.width, g.height
}
