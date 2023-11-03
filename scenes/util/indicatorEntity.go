package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddOnCollisionIndicatorEntity(
	scene *coldBrew.Scene,
	x, y, height, width, playerOffsetX, playerOffsetY float64,
	onPlayer bool,
	indicatorType components.IndicatorType,
) {

	indicatorEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.IndicatorStateAndConfigComponent,
	)

	body := tBokiComponents.NewRigidBodyBox(x, y, width, height, 0, false)
	components.RigidBodyComponent.SetValue(indicatorEntity, *body)

	state := components.NewIndicatorStateAndConfig(
		playerOffsetX, playerOffsetY,
		false,
		onPlayer,
		indicatorType,
	)
	components.IndicatorStateAndConfigComponent.SetValue(indicatorEntity, *state)

}
