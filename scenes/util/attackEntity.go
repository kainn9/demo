package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddAttackEntity(scene *coldBrew.Scene, attackData components.AttackData) {

	attackEntity := scene.AddEntity(
		components.AttackDataComponent,
		components.AttackHitboxesComponent,
	)

	emptyBoxes := []*tBokiComponents.RigidBody{}

	components.AttackHitboxesComponent.SetValue(attackEntity, emptyBoxes)
	components.AttackDataComponent.SetValue(attackEntity, attackData)
}
