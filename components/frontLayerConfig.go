package components

import "github.com/yohamta/donburi"

type FrontLayerConfig struct {
	SceneAssetPath string
}

var FrontLayerComponent = donburi.NewComponentType[FrontLayerConfig]()

func NewFrontLayerConfig(sceneAssetPath string) *FrontLayerConfig {
	return &FrontLayerConfig{
		SceneAssetPath: sceneAssetPath,
	}
}
