package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/constants"
	"github.com/kainn9/demo/scenes"
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

	loaderImage := ebiten.NewImage(constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT)
	loaderImage.Fill(color.RGBA{B: 255})

	manager := coldBrew.NewManager(constants.SCENE_CACHE_LIMIT, constants.MAX_TICKS, loaderImage)

	firstScene := scenes.Intro.LevelOneScene
	manager.LoadScene(firstScene)

	g := &game{
		width:   constants.SCREEN_WIDTH,
		height:  constants.SCREEN_HEIGHT,
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
		err := g.manager.LoadScene(scenes.Intro.LevelOneScene)
		if err != nil {
			log.Println("Yo!", err)
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		err := g.manager.LoadScene(scenes.Intro.LevelTwoScene)
		if err != nil {
			log.Println("Yo!", err)
		}
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
