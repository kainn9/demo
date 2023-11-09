package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
)

// Note: While the system is camera based, and the camera is an image, we do not consider it a
// a render system because it does not render anything. It only updates the camera position(aka preps the
// data/options that will be used during the render phase). Setting the camera position during the draw
// phase, leads to a lot of jittering due to the draw phase being tied to the frame rate(I think). Its much
// more stable to handle this in the sim phase.

type CameraPositionHandlerSystem struct {
	scene    *coldBrew.Scene
	Zoomed   bool
	ZoomPosY float64
}

func NewCameraPositionHandler(scene *coldBrew.Scene, zoomed bool) *CameraPositionHandlerSystem {
	return &CameraPositionHandlerSystem{
		ZoomPosY: 50,
		scene:    scene,
		Zoomed:   zoomed,
	}
}

func (sys CameraPositionHandlerSystem) Run(dt float64, _ *donburi.Entry) {

	world := sys.scene.World
	mapWidth := sys.scene.Width
	mapHeight := sys.scene.Height

	cameraEntity := systemsUtil.GetCameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)

	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	// Only smooth cam for player.
	useSmoothCam := true

	// We stop moving the camera, if the player is within the bounds of the screen.
	// E.g., bottom left, top left, bottom right, top right.
	halfScreenWidthInt := clientGlobals.SCREEN_WIDTH / 2
	halfScreenHeightInt := clientGlobals.SCREEN_HEIGHT / 2 // Half of the screen height

	halfScreenWidth := float64(halfScreenWidthInt)
	halfScreenHeight := float64(halfScreenHeightInt)

	xBoundaryLeft := halfScreenWidth
	xBoundaryRight := float64(mapWidth - halfScreenWidthInt)

	yBoundaryTop := halfScreenHeight
	yBoundaryBottom := float64(mapHeight - halfScreenHeightInt) // Calculate bottom Y boundary

	playerInsideXBoundsLeft := playerBody.Pos.X < xBoundaryLeft
	playerInsideXBoundsRight := playerBody.Pos.X > xBoundaryRight
	playerInsideYBoundsTop := playerBody.Pos.Y < yBoundaryTop
	playerInsideYBoundsBottom := playerBody.Pos.Y > yBoundaryBottom

	// Handle X-axis
	if playerInsideXBoundsLeft {
		cameraUtil.SetPosition(camera, 0, camera.Y, useSmoothCam)
	} else if playerInsideXBoundsRight {
		cameraUtil.SetPosition(camera, float64(mapWidth-clientGlobals.SCREEN_WIDTH), camera.Y, useSmoothCam)
	} else {
		cameraUtil.SetPosition(camera, playerBody.Pos.X-halfScreenWidth, camera.Y, useSmoothCam)
	}

	// Handle Y-axis
	if playerInsideYBoundsTop {
		cameraUtil.SetPosition(camera, camera.X, 0, useSmoothCam)
	} else if playerInsideYBoundsBottom {
		cameraUtil.SetPosition(camera, camera.X, float64(mapHeight-clientGlobals.SCREEN_HEIGHT), useSmoothCam)
	} else {
		cameraUtil.SetPosition(camera, camera.X, playerBody.Pos.Y-halfScreenHeight, useSmoothCam)
	}

	// Todo: Consider moving this into a separate system.
	// The offset is just a hack until we pick an official
	// and zoom level and handle the assets accordingly.
	if sys.Zoomed {
		clientGlobals.SCREEN_WIDTH = 480
		clientGlobals.SCREEN_HEIGHT = 270
		camera.Height = clientGlobals.SCREEN_WIDTH
		camera.Width = clientGlobals.SCREEN_HEIGHT
		cameraUtil.SetPosition(camera, camera.X, sys.ZoomPosY, false)
	} else {
		clientGlobals.SCREEN_WIDTH = 640
		clientGlobals.SCREEN_HEIGHT = 360
		camera.Height = clientGlobals.SCREEN_WIDTH
		camera.Width = clientGlobals.SCREEN_HEIGHT
	}

}
