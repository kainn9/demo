package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	animUtil "github.com/kainn9/demo/systems/render/util/anim"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type IndicatorRendererSystem struct {
	scene *coldBrew.Scene
}

func (sys IndicatorRendererSystem) IndicatorQuery() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.IndicatorStateAndConfigComponent),
	)
}

func NewIndicatorRenderer(scene *coldBrew.Scene) *IndicatorRendererSystem {
	return &IndicatorRendererSystem{
		scene: scene,
	}
}

func (sys IndicatorRendererSystem) Draw(screen *ebiten.Image, _ *donburi.Entry) {

	world := sys.scene.World

	if systemsUtil.IsChatActive(sys.scene.World) {
		return
	}

	cameraEntity := systemsUtil.GetCameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	uiSingletonEntity := systemsUtil.GetUISingletonEntity(world)
	UISpritesMap := components.SpritesMapComponent.Get(uiSingletonEntity)

	sys.IndicatorQuery().Each(world, func(indicatorEntity *donburi.Entry) {
		sys.renderIndicator(indicatorEntity, camera, playerBody, UISpritesMap)
	})

}

func (sys IndicatorRendererSystem) renderIndicator(
	indicatorEntity *donburi.Entry,
	camera *components.Camera,
	playerBody *tBokiComponents.RigidBody,
	UISpritesMap *map[string]*components.Sprite,
) {

	indicatorConfigAndState := components.IndicatorStateAndConfigComponent.Get(indicatorEntity)

	if indicatorConfigAndState.State.Active {
		indicatorSprite := (*UISpritesMap)[string(indicatorConfigAndState.Config.Type)]
		indicatorOpts := &ebiten.DrawImageOptions{}

		currentFrame := animUtil.GetAnimFrame(
			sys.scene.Manager,
			indicatorSprite,
		)

		xPos := indicatorConfigAndState.Config.X
		yPos := indicatorConfigAndState.Config.Y

		if indicatorConfigAndState.Config.OnPlayer {
			xPos += playerBody.Pos.X
			yPos += playerBody.Pos.Y
		}

		cameraUtil.Translate(camera, indicatorOpts, xPos, yPos)
		cameraUtil.AddImage(camera, currentFrame, indicatorOpts)

	}
}
