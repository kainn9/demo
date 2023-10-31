package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientConstants "github.com/kainn9/demo/constants/client"
)

func AddCameraEntity(scene *coldBrew.Scene, x, y float64) {

	cameraEntity := scene.AddEntity(
		components.CameraComponent,
	)

	components.CameraComponent.SetValue(
		cameraEntity,
		*components.NewCamera(x, y, clientConstants.SCREEN_WIDTH, clientConstants.SCREEN_HEIGHT),
	)

}
