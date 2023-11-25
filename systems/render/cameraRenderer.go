package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

type CameraRendererSystem struct {
	scene *coldBrew.Scene
}

func NewCameraRenderer(scene *coldBrew.Scene) *CameraRendererSystem {
	return &CameraRendererSystem{
		scene: scene,
	}
}

func (sys CameraRendererSystem) Draw(screen *ebiten.Image, _ *donburi.Entry) {
	world := sys.scene.World
	cameraEntity := systemsUtil.CameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)

	cameraUtil.Render(camera, screen)
}
