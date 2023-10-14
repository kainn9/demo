package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
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

	manager := coldBrew.NewManager(constants.SCENE_CACHE_LIMIT, loaderImage)

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

	activeScene := g.manager.GetActiveScene()

	go activeScene.Load()

	activeScene.Sync()

	deltaTime := (0.017)
	activeScene.Sim(deltaTime)

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {

	activeScene := g.manager.GetActiveScene()
	activeScene.Draw(screen)

}

func (g *game) Layout(w, h int) (int, int) {
	return g.width, g.height
}
