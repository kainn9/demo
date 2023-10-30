package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

type CameraRendererRendererSystem struct {
	scene *coldBrew.Scene
}

func NewCameraRendererRenderer(scene *coldBrew.Scene) *CameraRendererRendererSystem {
	return &CameraRendererRendererSystem{
		scene: scene,
	}
}

func (sys *CameraRendererRendererSystem) Draw(screen *ebiten.Image, _ *donburi.Entry) {
	world := sys.scene.World
	cameraEntity := systemsUtil.GetCameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)
	cameraUtil.Render(camera, screen)
}
