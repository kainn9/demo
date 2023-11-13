package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddBlockEntity(scene *coldBrew.Scene, x, y, w, h, rotation float64) {

	blockEntity := scene.AddEntity(
		components.RigidBodyComponent,
		tags.BlockTag,
	)

	blockBody := tBokiComponents.NewRigidBodyBox(x, y, w, h, 0, true)
	blockBody.Elasticity = 0
	blockBody.Rotation = rotation
	blockBody.UpdateVertices()

	components.RigidBodyComponent.SetValue(
		blockEntity,
		*blockBody,
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
