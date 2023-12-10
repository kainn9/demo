package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
)

func AddSceneData(scene *coldBrew.Scene, assetPath string) {
	sceneDataEntity := scene.AddEntity(
		components.SceneDataComponent,
	)

	sceneData := components.NewSceneData(assetPath)

	components.SceneDataComponent.SetValue(
		sceneDataEntity,
		*sceneData,
	)

}
