package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"

	"github.com/kainn9/demo/queries"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
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

func (sys ParallaxBackgroundRendererSystem) Draw(screen *ebiten.Image, _ *donburi.Entry) {

	world := sys.scene.World
	cameraEntity := systemsUtil.GetCameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)

	// Since this is the first Render System, we clear the camera here.
	// This may get moved to a separate system in the future.
	cameraUtil.Clear(camera)

	queries.ParallaxBackGroundLayerQuery.Each(world, func(bgEntity *donburi.Entry) {
		sys.renderPLaxBGImage(bgEntity, camera)
	})

}

func (sys ParallaxBackgroundRendererSystem) renderPLaxBGImage(bgEntity *donburi.Entry, camera *components.Camera) {
	pLaxLayerConfig := components.ParallaxLayerConfigComponent.Get(bgEntity)
	sprite := components.SpriteComponent.Get(bgEntity)

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

}
