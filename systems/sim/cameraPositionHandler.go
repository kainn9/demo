package simSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
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

	tx := (1 - camera.Zoom) * float64(clientGlobals.SCREEN_WIDTH) / 2

	xOffsetLeft := tx / float64(camera.Zoom)
	xOffsetCenter := playerBody.Pos.X - float64(clientGlobals.SCREEN_WIDTH/2)
	xOffsetRight := xBoundaryRight - float64(clientGlobals.SCREEN_WIDTH)/2

	yBoundaryTop := (float64(clientGlobals.SCREEN_HEIGHT) / 2) / camera.Zoom
	yBoundaryBottom := float64(mapHeight) - yBoundaryTop

	ty := (1 - camera.Zoom) * float64(clientGlobals.SCREEN_HEIGHT) / 2

	yOffsetTop := ty / float64(camera.Zoom)
	yOffsetCenter := playerBody.Pos.Y - float64(clientGlobals.SCREEN_HEIGHT/2)
	yOffsetBottom := yBoundaryBottom - float64(clientGlobals.SCREEN_HEIGHT)/2

	playerInsideXBoundsLeft := playerBody.Pos.X < xBoundaryLeft

	playerInsideXBoundsRight := playerBody.Pos.X > xBoundaryRight

	playerInsideYBoundsTop := playerBody.Pos.Y < yBoundaryTop

	playerInsideYBoundsBottom := playerBody.Pos.Y > yBoundaryBottom

	// Handle X-axis
	switch {
	case playerInsideXBoundsLeft:
		cameraUtil.SetPosition(camera, xOffsetLeft, camera.Y, useSmoothCam)
	case playerInsideXBoundsRight:
		cameraUtil.SetPosition(camera, xOffsetRight, camera.Y, useSmoothCam)
	default:
		cameraUtil.SetPosition(camera, xOffsetCenter, camera.Y, useSmoothCam)
	}

	// Handle Y-axis
	switch {
	case playerInsideYBoundsTop:
		cameraUtil.SetPosition(camera, camera.X, yOffsetTop, useSmoothCam)
	case playerInsideYBoundsBottom:
		cameraUtil.SetPosition(camera, camera.X, yOffsetBottom, useSmoothCam)
	default:
		cameraUtil.SetPosition(camera, camera.X, yOffsetCenter, useSmoothCam)
	}

	if ebiten.IsKeyPressed(ebiten.KeyF) {
		camera.Zoom += 0.1
	}

	if ebiten.IsKeyPressed(ebiten.KeyG) {
		camera.Zoom -= 0.1
	}
}
