package renderUISystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	"github.com/kainn9/demo/queries"
	animUtil "github.com/kainn9/demo/systems/render/util/anim"
	textUtil "github.com/kainn9/demo/systems/render/util/text"
)

type TitleRendererSystem struct {
	scene    *coldBrew.Scene
	runSetup bool
}

func (sys TitleRendererSystem) Query() *donburi.Query {
	return queries.ParallaxBackGroundLayerQuery
}

func NewTitleRenderer(scene *coldBrew.Scene) *TitleRendererSystem {
	return &TitleRendererSystem{
		scene:    scene,
		runSetup: true,
	}
}

func (sys *TitleRendererSystem) Draw(screen *ebiten.Image, bgEntity *donburi.Entry) {

	sprite := components.SpriteComponent.Get(bgEntity)

	if sys.runSetup {
		sprite.AnimationConfig = components.NewAnimationConfig(
			clientGlobals.SCREEN_WIDTH,
			clientGlobals.SCREEN_HEIGHT,
			8,
			30,
			true,
		)

		sys.runSetup = false
	}

	frame := animUtil.GetAnimFrame(sys.scene.Manager.TickHandler, sprite)

	opts := &ebiten.DrawImageOptions{}
	screen.DrawImage(frame, opts)

	textUtil.RenderTextDefault("Press Enter to Start", 400, 100, 700, 0, 60, &sys.scene.World, sys.scene.Manager.TickHandler, screen)
}
