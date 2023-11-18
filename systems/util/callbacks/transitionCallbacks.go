package callbacksUtil

import (
	"github.com/kainn9/coldBrew"
	simSystems "github.com/kainn9/demo/systems/sim"
)

func AttachTransitionCallback(scene *coldBrew.Scene, callback simSystems.SceneTransitionPermissionCallback) {

	for _, sys := range scene.Systems {
		switch tSys := sys.(type) {

		case *simSystems.SceneTransitionHandlerSystem:

			tSys.PermissionCallbacks = append(tSys.PermissionCallbacks, callback)
		}
	}
}
