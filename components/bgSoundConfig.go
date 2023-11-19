package components

import "github.com/yohamta/donburi"

type BgSoundConfig struct {
	SceneAssetsPath string // "<sceneCollection>/<sceneName>/" e.g., "intro/levelOne/"
}

var BgSoundConfigComponent = donburi.NewComponentType[BgSoundConfig]()

func NewBgSoundConfig(sceneAssetsPath string) *BgSoundConfig {
	return &BgSoundConfig{
		SceneAssetsPath: sceneAssetsPath,
	}
}
