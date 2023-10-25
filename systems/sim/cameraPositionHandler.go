package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/constants"
	cameraUtil "github.com/kainn9/demo/systems/render/util"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
)

// Note: While the system is camera based, and the camera is an image, we do not consider it a
// a render system because it does not render anything. It only updates the camera position(aka preps the
// data/options that will be used during the render phase). Setting the camera position during the draw
// phase, leads to a lot of jittering due to the draw phase being tied to the frame rate(I think). Its much
// more stable to handle this in the sim phase.

type CameraPositionHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewCameraPositionHandler(scene *coldBrew.Scene) *CameraPositionHandlerSystem {
	return &CameraPositionHandlerSystem{
		scene: scene,
	}
}

func (sys *CameraPositionHandlerSystem) Run(dt float64, _ *donburi.Entry) {
	world := sys.scene.World
	mapWidth := sys.scene.Width
	mapHeight := sys.scene.Height // Add map height
	camera := systemsUtil.GetCamera(world)
	playerPos := systemsUtil.GetPlayerPos(world)

	halfScreenWidthInt := constants.SCREEN_WIDTH / 2
	halfScreenHeightInt := constants.SCREEN_HEIGHT / 2 // Half of the screen height

	halfScreenWidth := float64(halfScreenWidthInt)
	halfScreenHeight := float64(halfScreenHeightInt)

	xBoundaryLeft := halfScreenWidth
	xBoundaryRight := float64(mapWidth - halfScreenWidthInt)

	yBoundaryTop := halfScreenHeight
	yBoundaryBottom := float64(mapHeight - halfScreenHeightInt) // Calculate bottom Y boundary

	playerInsideXBoundsLeft := playerPos.X < xBoundaryLeft
	playerInsideXBoundsRight := playerPos.X > xBoundaryRight
	playerInsideYBoundsTop := playerPos.Y < yBoundaryTop
	playerInsideYBoundsBottom := playerPos.Y > yBoundaryBottom

	// Handle X-axis
	if playerInsideXBoundsLeft {
		cameraUtil.SetPosition(camera, 0, camera.Y, true)
	} else if playerInsideXBoundsRight {
		cameraUtil.SetPosition(camera, float64(mapWidth-constants.SCREEN_WIDTH), camera.Y, true)
	} else {
		cameraUtil.SetPosition(camera, playerPos.X-halfScreenWidth, camera.Y, true)
	}

	// Handle Y-axis
	if playerInsideYBoundsTop {
		cameraUtil.SetPosition(camera, camera.X, 0, true)
	} else if playerInsideYBoundsBottom {
		cameraUtil.SetPosition(camera, camera.X, float64(mapHeight-constants.SCREEN_HEIGHT), true)
	} else {
		cameraUtil.SetPosition(camera, camera.X, playerPos.Y-halfScreenHeight, true)
	}
}
