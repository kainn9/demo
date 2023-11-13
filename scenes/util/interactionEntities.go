package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	UIGlobals "github.com/kainn9/demo/globalConfig/UI"
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
	x, y, width, height float64,
	targetScene coldBrew.SceneFace,
	spawnX, spawnY, camX, camY float64,
) {

	transitionEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.IndicatorStateAndConfigComponent,
		components.SceneTransitionStateAndConfigComponent,
	)

	body := tBokiComponents.NewRigidBodyBox(x, y, width, height, 0, false)
	components.RigidBodyComponent.SetValue(transitionEntity, *body)

	offX := UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_INTERACT].X
	offY := UIGlobals.IndicatorPlayerOffsets[UIGlobals.CurrentLayout][UIGlobals.INDICATOR_INTERACT].Y

	indicatorState := components.NewIndicatorStateAndConfig(
		offX, offY,
		false,
		true,
		UIGlobals.INDICATOR_INTERACT,
	)

	components.IndicatorStateAndConfigComponent.SetValue(transitionEntity, *indicatorState)
	transitionState := components.NewSceneTransitionStateAndConfig(spawnX, spawnY, camX, camY, targetScene)
	components.SceneTransitionStateAndConfigComponent.SetValue(transitionEntity, *transitionState)
}
