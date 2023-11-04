package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddOnCollisionIndicatorEntity(
	scene *coldBrew.Scene,
	x, y, width, height, playerOffsetX, playerOffsetY float64,
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

func AddSceneTransitionEntity(
	scene *coldBrew.Scene,
	x, y, width, height, playerOffsetX, playerOffsetY float64,
	onPlayer bool,
	indicatorType components.IndicatorType,
	targetScene coldBrew.SceneFace,
	spawnX, spawnY, camX, camY float64,
	clickBased bool,
) {

	transitionEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.IndicatorStateAndConfigComponent,
		components.SceneTransitionStateAndConfigComponent,
	)

	body := tBokiComponents.NewRigidBodyBox(x, y, width, height, 0, false)
	components.RigidBodyComponent.SetValue(transitionEntity, *body)

	indicatorState := components.NewIndicatorStateAndConfig(
		playerOffsetX, playerOffsetY,
		false,
		onPlayer,
		indicatorType,
	)

	components.IndicatorStateAndConfigComponent.SetValue(transitionEntity, *indicatorState)
	transitionState := components.NewSceneTransitionStateAndConfig(spawnX, spawnY, camX, camY, targetScene, clickBased)
	components.SceneTransitionStateAndConfigComponent.SetValue(transitionEntity, *transitionState)
}
