package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/queries"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

type FrontLayerRendererSystem struct {
	scene *coldBrew.Scene
}

func NewFrontLayerRenderer(scene *coldBrew.Scene) *FrontLayerRendererSystem {
	return &FrontLayerRendererSystem{
		scene: scene,
	}
}

func (sys FrontLayerRendererSystem) Draw(screen *ebiten.Image, _ *donburi.Entry) {

	world := sys.scene.World
	cameraEntity := systemsUtil.GetCameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)

	queries.FrontLayerQuery.Each(world, func(layerEntity *donburi.Entry) {

		sprite := assetComponents.SpriteComponent.Get(layerEntity)
		drawOptions := &ebiten.DrawImageOptions{}
		cameraUtil.Translate(camera, drawOptions, 0, 0)
		cameraUtil.AddImage(camera, sprite.Image, drawOptions)

	})

}
