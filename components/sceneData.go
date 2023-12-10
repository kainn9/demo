package components

import (
	sceneProto "github.com/kainn9/demo/proto"
	"github.com/yohamta/donburi"
)

var SceneDataComponent = donburi.NewComponentType[SceneData]()

type SceneData struct {
	AssetPath     string
	TerrainBlocks *sceneProto.TerrainBlocks
}

func NewSceneData(assetPath string) *SceneData {
	return &SceneData{
		AssetPath:     assetPath,
		TerrainBlocks: &sceneProto.TerrainBlocks{},
	}
}
