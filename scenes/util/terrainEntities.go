package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddFloorEntity(scene *coldBrew.Scene, x, y, w, h, rotation float64) {

	floorEntity := scene.AddEntity(
		components.RigidBodyComponent,
		tags.FloorTag,
	)

	floorBody := tBokiComponents.NewRigidBodyBox(x, y, w, h, 0, true)
	floorBody.Elasticity = 0
	floorBody.Rotation = rotation
	floorBody.UpdateVertices()

	components.RigidBodyComponent.SetValue(
		floorEntity,
		*floorBody,
	)
}

// Note: Platforms do not support rotation.
func AddPlatformEntity(scene *coldBrew.Scene, x, y, w, h float64) {

	platformEntity := scene.AddEntity(
		components.RigidBodyComponent,
		tags.PlatformTag,
	)

	platformBody := tBokiComponents.NewRigidBodyBox(x, y, w, h, 0, true)
	platformBody.Elasticity = 0
	platformBody.UpdateVertices()

	components.RigidBodyComponent.SetValue(
		platformEntity,
		*platformBody,
	)
}

// Note: Ladders do not support rotation.
func AddLadderEntity(scene *coldBrew.Scene, x, y, w, h float64) {

	ladderEntity := scene.AddEntity(
		components.RigidBodyComponent,
		tags.LadderTag,
	)

	ladderBody := tBokiComponents.NewRigidBodyBox(x, y, w, h, 0, true)
	ladderBody.UpdateVertices()

	components.RigidBodyComponent.SetValue(
		ladderEntity,
		*ladderBody,
	)
}
