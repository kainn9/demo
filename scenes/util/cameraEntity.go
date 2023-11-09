package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
)

func AddCameraEntity(scene *coldBrew.Scene, x, y float64) {

	cameraEntity := scene.AddEntity(
		components.CameraComponent,
	)

	components.CameraComponent.SetValue(
		cameraEntity,
		*components.NewCamera(x, y, clientGlobals.SCREEN_WIDTH, clientGlobals.SCREEN_HEIGHT),
	)

}
