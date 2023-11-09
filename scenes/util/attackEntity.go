package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddAttackEntity(scene *coldBrew.Scene, AS components.AttackState) {

	playerEntity := scene.AddEntity(
		components.AttackBoxesComponent,
		components.AttackStateComponent,
	)

	emptyBoxes := []*tBokiComponents.RigidBody{
		tBokiComponents.NewRigidBodyBox(0, 0, 0, 0, 0, false),
	}

	components.AttackBoxesComponent.SetValue(playerEntity, emptyBoxes)
	components.AttackStateComponent.SetValue(playerEntity, AS)

}
