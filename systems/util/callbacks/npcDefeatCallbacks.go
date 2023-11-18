package callbacksUtil

import (
	"github.com/kainn9/coldBrew"
	simNpcSystems "github.com/kainn9/demo/systems/sim/npc"
)

func AttachNpcDefeatCallback(scene *coldBrew.Scene, callback simNpcSystems.DefeatCallback) {

	for _, sys := range scene.Systems {
		switch dSys := sys.(type) {

		case *simNpcSystems.NpcDefeatedHandlerSystem:
			dSys.DefeatCallbacks = append(dSys.DefeatCallbacks, callback)
		}
	}
}
