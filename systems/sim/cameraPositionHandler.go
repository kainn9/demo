package simSystems

import (
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
	scene    *coldBrew.Scene
	zoomHack float64
}

func NewCameraPositionHandler(scene *coldBrew.Scene) *CameraPositionHandlerSystem {
	return &CameraPositionHandlerSystem{
		scene: scene,
	}
}

func (sys *CameraPositionHandlerSystem) Run(dt float64, _ *donburi.Entry) {
	world := sys.scene.World
	mapWidth, mapHeight := sys.scene.Width, sys.scene.Height
	cameraEntity := systemsUtil.GetCameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)
	playerEntity := systemsUtil.GetPlayerEntity(world)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	useSmoothCam := true

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

	// Handle X-axis
	switch {
	case playerInsideXBoundsLeft:
		cameraSimUtil.SetPosition(camera, xOffsetLeft, camera.Y, useSmoothCam)
	case playerInsideXBoundsRight:
		cameraSimUtil.SetPosition(camera, xOffsetRight, camera.Y, useSmoothCam)
	default:
		cameraSimUtil.SetPosition(camera, xOffsetCenter, camera.Y, useSmoothCam)
	}

	// Handle Y-axis
	switch {
	case playerInsideYBoundsTop:
		cameraSimUtil.SetPosition(camera, camera.X, yOffsetTop, useSmoothCam)
	case playerInsideYBoundsBottom:
		cameraSimUtil.SetPosition(camera, camera.X, yOffsetBottom, useSmoothCam)
	default:
		cameraSimUtil.SetPosition(camera, camera.X, yOffsetCenter, useSmoothCam)
	}

	if ebiten.IsKeyPressed(ebiten.KeyF) {
		camera.Zoom += 0.1
	}

	if ebiten.IsKeyPressed(ebiten.KeyG) {
		camera.Zoom -= 0.1
	}
}
