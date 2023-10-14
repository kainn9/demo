package renderSystems

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	cameraUtil "github.com/kainn9/demo/systems/render/util"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

type PlayerRendererSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerRenderer(scene *coldBrew.Scene) *PlayerRendererSystem {
	return &PlayerRendererSystem{
		scene: scene,
	}
}

func (sys *PlayerRendererSystem) Draw(screen *ebiten.Image, entry *donburi.Entry) {
	world := sys.scene.World

	x, y := systemsUtil.GetPlayerPos(world)
	camera := systemsUtil.GetCamera(world)

	camImg := ebiten.NewImage(15, 15)
	camImg.Fill(color.RGBA{R: 255})

	opts := &ebiten.DrawImageOptions{}
	cameraUtil.Translate(camera, opts, x, y)
	cameraUtil.AddImage(camera, camImg, opts)

	cameraUtil.Render(camera, screen)

}
