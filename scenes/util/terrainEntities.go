package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	"github.com/yohamta/donburi"
)

func AddBlockEntity(scene *coldBrew.Scene, x, y, w, h, rotation float64) *donburi.Entry {

	blockEntity := scene.AddEntity(
		components.RigidBodyComponent,
		tags.BlockTag,
	)

	mass := 0.0

	blockBody := tBokiComponents.NewRigidBodyBox(x, y, w, h, mass, true)
	blockBody.Elasticity = 0
	blockBody.Rotation = rotation
	blockBody.Friction = 0

	blockBody.UpdateVertices()

	components.RigidBodyComponent.SetValue(
		blockEntity,
		*blockBody,
	)

	return blockEntity

}

func AddBlockEntityPoly(scene *coldBrew.Scene, vertices []tBokiVec.Vec2) *donburi.Entry {

	blockEntity := scene.AddEntity(
		components.RigidBodyComponent,
		tags.BlockTag,
	)

	mass := 0.0
	angular := false

	blockBody := tBokiComponents.NewRigidBodyPolygonWorld(mass, vertices, angular)
	blockBody.Elasticity = 0
	blockBody.Friction = 0

	blockBody.UpdateVertices()

	components.RigidBodyComponent.SetValue(
		blockEntity,
		*blockBody,
	)

	return blockEntity

}

// Note: Platforms do not support rotation yet.
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

func AddWalls(scene *coldBrew.Scene, sceneWidth, sceneHeight float64) {
	AddBlockEntity(scene, 0, sceneHeight/2, 10, sceneHeight, 0)
	AddBlockEntity(scene, sceneWidth, 0+sceneHeight/2, 10, sceneHeight, 0)
}
