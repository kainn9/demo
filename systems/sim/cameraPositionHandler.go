package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientConstants "github.com/kainn9/demo/constants/client"
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
	scene *coldBrew.Scene
}

func NewCameraPositionHandler(scene *coldBrew.Scene) *CameraPositionHandlerSystem {
	return &CameraPositionHandlerSystem{
		scene: scene,
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

	// We stop moving the camera, if the player is within the bounds of the screen.
	// E.g., bottom left, top left, bottom right, top right.
	halfScreenWidthInt := clientConstants.SCREEN_WIDTH / 2
	halfScreenHeightInt := clientConstants.SCREEN_HEIGHT / 2 // Half of the screen height

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
		cameraUtil.SetPosition(camera, 0, camera.Y, true)
	} else if playerInsideXBoundsRight {
		cameraUtil.SetPosition(camera, float64(mapWidth-clientConstants.SCREEN_WIDTH), camera.Y, true)
	} else {
		cameraUtil.SetPosition(camera, playerBody.Pos.X-halfScreenWidth, camera.Y, true)
	}

	// Handle Y-axis
	if playerInsideYBoundsTop {
		cameraUtil.SetPosition(camera, camera.X, 0, true)
	} else if playerInsideYBoundsBottom {
		cameraUtil.SetPosition(camera, camera.X, float64(mapHeight-clientConstants.SCREEN_HEIGHT), true)
	} else {
		cameraUtil.SetPosition(camera, camera.X, playerBody.Pos.Y-halfScreenHeight, true)
	}
}
