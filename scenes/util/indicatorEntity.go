package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	UIConstants "github.com/kainn9/demo/constants/UI"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddOnCollisionIndicatorEntity(
	scene *coldBrew.Scene,
	x, y, height, width float64,
	onPlayer bool,
	indicatorType UIConstants.IndicatorType,
) {

	indicatorEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.IndicatorStateAndConfigComponent,
	)

	body := tBokiComponents.NewRigidBodyBox(x, y, width, height, 0, false)
	components.RigidBodyComponent.SetValue(indicatorEntity, *body)

	state := components.NewIndicatorStateAndConfig(
		false,
		onPlayer,
		indicatorType,
	)
	components.IndicatorStateAndConfigComponent.SetValue(indicatorEntity, *state)

}
