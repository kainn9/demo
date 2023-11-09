package renderDebugSystems

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

type DebugRigidBodyRendererSystem struct {
	scene *coldBrew.Scene
}

func NewDebugRigidBodyRenderer(scene *coldBrew.Scene) *DebugRigidBodyRendererSystem {
	return &DebugRigidBodyRendererSystem{
		scene: scene,
	}
}

func (DebugRigidBodyRendererSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Or(
			filter.Contains(components.RigidBodyComponent),
			filter.Contains(components.AttackBoxesComponent),
		),
	)
}

func (sys DebugRigidBodyRendererSystem) Draw(screen *ebiten.Image, entity *donburi.Entry) {
	if clientGlobals.DEBUG_MODE == false {
		return
	}

	bodies := make([]*tBokiComponents.RigidBody, 0)

	if entity.HasComponent(components.AttackBoxesComponent) {
		attackBoxes := components.AttackBoxesComponent.Get(entity)

		for _, box := range *attackBoxes {
			bodies = append(bodies, box)
		}

	} else {
		bodies = append(bodies, components.RigidBodyComponent.Get(entity))
	}

	for _, body := range bodies {
		cameraEntity := systemsUtil.GetCameraEntity(sys.scene.World)
		camera := components.CameraComponent.Get(cameraEntity)

		red := color.RGBA{R: 255, G: 0, B: 0, A: 255}
		blue := color.RGBA{R: 0, G: 0, B: 255, A: 255}

		if body.Circle != nil {
			debugDrawCircleBody(screen, camera, *body, red)
		}

		if body.Polygon != nil {
			debugDrawPolygonBody(screen, camera, *body, red)
		}

		debugDrawBroadPhaseSkin(screen, camera, *body, blue)
	}

}
