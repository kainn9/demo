package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	textUtil "github.com/kainn9/demo/systems/render/util/text"
)

type sceneTitleRendererSystem struct {
	scene     *coldBrew.Scene
	title     string
	firstRun  bool
	fadeTimer int
}

func NewSceneTitleRenderer(scene *coldBrew.Scene, title string) *sceneTitleRendererSystem {
	return &sceneTitleRendererSystem{
		scene: scene,
		title: title,
	}
}

func (sys *sceneTitleRendererSystem) Draw(screen *ebiten.Image, _ *donburi.Entry) {

	timeSinceSceneBecameActive := sys.scene.Manager.TickHandler.TicksSinceNTicks(sys.scene.LastActiveTick)

	fadeTimer := 240
	buffer := 60

	if timeSinceSceneBecameActive < buffer {
		return
	}

	if timeSinceSceneBecameActive > fadeTimer {
		return
	}

	text := sys.title
	width := textUtil.GetDefaultFontWidth(text, 420, 3)

	sx := float64((clientGlobals.SCREEN_WIDTH / 2) - (width / 2))
	sy := float64(60)

	textUtil.RenderTextDefault(text, sx, sy, 420, sys.scene.LastActiveTick+buffer, 30, 3, &sys.scene.World, sys.scene.Manager.TickHandler, screen)
}
