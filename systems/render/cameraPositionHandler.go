package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/constants"
	cameraUtil "github.com/kainn9/demo/systems/render/util"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
)

type CameraPositionHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewCameraPositionHandler(scene *coldBrew.Scene) *CameraPositionHandlerSystem {
	return &CameraPositionHandlerSystem{
		scene: scene,
	}
}

func (sys *CameraPositionHandlerSystem) Draw(screen *ebiten.Image, entry *donburi.Entry) {

	world := sys.scene.World

	mapWidth := sys.scene.Width

	camera := systemsUtil.GetCamera(world)

	x, y := systemsUtil.GetPlayerPos(world)

	cameraUtil.Clear(camera)

	halfScreenInt := constants.SCREEN_WIDTH / 2
	halfScreen := float64(halfScreenInt)

	xBoundaryLeft := halfScreen

	xBoundaryRight := float64(mapWidth - halfScreenInt)

	playerInsideXBoundsLeft := x < xBoundaryLeft

	playerInsideXBoundsRight := x > xBoundaryRight

	if playerInsideXBoundsLeft {
		cameraUtil.SetPosition(camera, 0, 0)
		return
	}

	if playerInsideXBoundsRight {
		cameraUtil.SetPosition(camera, float64(mapWidth-constants.SCREEN_WIDTH), y)
		return
	}

	cameraUtil.SetPosition(camera, x-halfScreen, y)

}
