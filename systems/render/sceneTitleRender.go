package renderSystems

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	textUtil "github.com/kainn9/demo/systems/render/util/text"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

type sceneTitleRendererSystem struct {
	scene                    *coldBrew.Scene
	title                    string
	enabled, displayingTitle bool
	startTick                int
}

func NewSceneTitleRenderer(scene *coldBrew.Scene, title string) *sceneTitleRendererSystem {
	return &sceneTitleRendererSystem{
		scene:   scene,
		title:   title,
		enabled: true,
	}
}

func (sys *sceneTitleRendererSystem) Draw(screen *ebiten.Image, _ *donburi.Entry) {
	th := sys.scene.Manager.TickHandler

	if th.CurrentTick() == sys.scene.LastActiveTick {
		sys.enabled = true
	}

	if !sys.enabled {
		return
	}

	isChatActive, _ := systemsUtil.IsChatActive(sys.scene.World)
	if isChatActive && !sys.displayingTitle {
		sys.startTick = th.CurrentTick()
		return
	} else {

		floatSceneLastActiveTick := float64(sys.scene.LastActiveTick)
		floatSysStartTick := float64(sys.startTick)

		sys.startTick = int(math.Max(floatSceneLastActiveTick, floatSysStartTick))
	}

	timeSinceSceneBecameActive := th.TicksSinceNTicks(sys.startTick)

	fadeTimer := 240
	buffer := 60

	if timeSinceSceneBecameActive < buffer {
		return
	}

	if timeSinceSceneBecameActive > fadeTimer {
		sys.enabled = false
		sys.displayingTitle = false
		return
	}

	sys.displayingTitle = true

	text := sys.title
	width := textUtil.GetDefaultFontWidth(text, 420, 3)

	sx := float64((clientGlobals.SCREEN_WIDTH / 2) - (width / 2))
	sy := float64(60)

	textUtil.RenderTextDefault(text, sx, sy, 420, sys.startTick+buffer, 15, 3, &sys.scene.World, sys.scene.Manager.TickHandler, screen)
}
