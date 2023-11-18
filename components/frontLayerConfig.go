package components

import "github.com/yohamta/donburi"

var FrontLayerComponent = donburi.NewComponentType[FrontLayerConfig]()

type FrontLayerConfig struct {
	SceneAssetPath string
}

func NewFrontLayerConfig(sceneAssetPath string) *FrontLayerConfig {
	return &FrontLayerConfig{
		SceneAssetPath: sceneAssetPath,
	}
}
