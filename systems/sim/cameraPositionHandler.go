package simSystems

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	cameraSimUtil "github.com/kainn9/demo/systems/sim/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
	cameraSharedUtil "github.com/kainn9/demo/systems/util/camera"
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

func (sys *CameraPositionHandlerSystem) Run(dt float64, _ *donburi.Entry) {

	world := sys.scene.World
	mapWidth, mapHeight := sys.scene.Width, sys.scene.Height
	cameraEntity := systemsUtil.CameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)
	playerEntity := systemsUtil.PlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	xBoundaryLeft := (float64(clientGlobals.SCREEN_WIDTH) / 2) / camera.Zoom
	xBoundaryRight := float64(mapWidth) - xBoundaryLeft

	xOffsetLeft := cameraSharedUtil.ZoomSpacingX(camera.Zoom) / float64(camera.Zoom)
	xOffsetCenter := playerBody.Pos.X - float64(clientGlobals.SCREEN_WIDTH/2)
	xOffsetRight := xBoundaryRight - float64(clientGlobals.SCREEN_WIDTH)/2

	yBoundaryTop := (float64(clientGlobals.SCREEN_HEIGHT) / 2) / camera.Zoom
	yBoundaryBottom := float64(mapHeight) - yBoundaryTop

	yOffsetTop := cameraSharedUtil.ZoomSpacingY(camera.Zoom) / float64(camera.Zoom)
	yOffsetCenter := playerBody.Pos.Y - float64(clientGlobals.SCREEN_HEIGHT/2)
	yOffsetBottom := yBoundaryBottom - float64(clientGlobals.SCREEN_HEIGHT)/2

	playerInsideXBoundsLeft := playerBody.Pos.X < xBoundaryLeft

	playerInsideXBoundsRight := playerBody.Pos.X > xBoundaryRight

	playerInsideYBoundsTop := playerBody.Pos.Y < yBoundaryTop

	playerInsideYBoundsBottom := playerBody.Pos.Y > yBoundaryBottom

	epsilon := 15.0

	if math.Abs(xOffsetCenter-camera.X) < epsilon {
		return
	}

	// Handle X-axis
	switch {
	case playerInsideXBoundsLeft:
		cameraSimUtil.SetPositionLerp(camera, xOffsetLeft, camera.Y)
	case playerInsideXBoundsRight:
		cameraSimUtil.SetPositionLerp(camera, xOffsetRight, camera.Y)
	default:
		cameraSimUtil.SetPositionLerp(camera, xOffsetCenter, camera.Y)
	}

	if math.Abs(yOffsetCenter-camera.Y) < epsilon {
		return
	}

	// Handle Y-axis
	switch {
	case playerInsideYBoundsTop:
		cameraSimUtil.SetPositionLerp(camera, camera.X, yOffsetTop)
	case playerInsideYBoundsBottom:
		cameraSimUtil.SetPositionLerp(camera, camera.X, yOffsetBottom)
	default:
		cameraSimUtil.SetPositionLerp(camera, camera.X, yOffsetCenter)
	}

	if ebiten.IsKeyPressed(ebiten.KeyF) {
		camera.Zoom += 0.1
	}

	if ebiten.IsKeyPressed(ebiten.KeyG) {
		camera.Zoom -= 0.1
	}
}
