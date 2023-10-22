package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/queries"
	cameraUtil "github.com/kainn9/demo/systems/render/util"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

type ParallaxBackgroundRendererSystem struct {
	scene *coldBrew.Scene
}

func NewParallaxBackgroundRenderer(scene *coldBrew.Scene) *ParallaxBackgroundRendererSystem {
	return &ParallaxBackgroundRendererSystem{
		scene: scene,
	}
}

func (sys *ParallaxBackgroundRendererSystem) Draw(screen *ebiten.Image, _ *donburi.Entry) {

	world := sys.scene.World

	camera := systemsUtil.GetCamera(world)

	queries.ParallaxBackGroundLayerQuery.Each(world, func(entity *donburi.Entry) {
		pLaxLayerConfig := assetComponents.ParallaxLayerConfigComponent.Get(entity)
		sprite := assetComponents.SpriteComponent.Get(entity)

		drawOptions := &ebiten.DrawImageOptions{}

		x := pLaxLayerConfig.CoefficientX

		if x != 0 {
			x = -(camera.X / pLaxLayerConfig.CoefficientX)

		}

		y := pLaxLayerConfig.CoefficientY

		if y != 0 {
			y = -(camera.Y / pLaxLayerConfig.CoefficientY)
		}

		if pLaxLayerConfig.AlwaysVisible {
			drawOptions.GeoM.Translate(camera.X/1.2, y)
			cameraUtil.Translate(camera, drawOptions, 0, y)

		} else {
			cameraUtil.Translate(camera, drawOptions, x, y)
		}

		cameraUtil.AddImage(camera, sprite.Image, drawOptions)

	})

}
