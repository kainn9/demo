package assetComponents

import "github.com/yohamta/donburi"

var ParallaxLayerConfigComponent = donburi.NewComponentType[ParallaxLayerConfig]()

type ParallaxLayerConfig struct {
	SceneAssetsPath string // "<sceneCollection>/<sceneName>/" e.g., "intro/levelOne/"

	// Slide speed for the layer.
	// 0 means no sliding.
	CoefficientX, CoefficientY float64

	ZIndex int

	// For things that we want to always be visible
	// on the camera...like a moon. The coefficients
	// will still dictate sliding, but within the
	// bounds of the camera(ish).
	AlwaysVisible bool
}

func NewParallaxLayerConfig(sceneAssetsPath string, zIndex int, coefficientX, coefficientY float64, alwaysVisible bool) *ParallaxLayerConfig {

	return &ParallaxLayerConfig{
		SceneAssetsPath: sceneAssetsPath,
		CoefficientX:    coefficientX,
		CoefficientY:    coefficientY,
		ZIndex:          zIndex,
		AlwaysVisible:   alwaysVisible,
	}
}
