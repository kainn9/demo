package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	assetComponents "github.com/kainn9/demo/components/assets"
)

func AddParallaxBackgroundEntity(scene *coldBrew.Scene, layers []*assetComponents.ParallaxLayerConfig) {

	for _, layer := range layers {

		bgLayerEntity := scene.AddEntity(assetComponents.ParallaxLayerConfigComponent, assetComponents.SpriteComponent)

		assetComponents.ParallaxLayerConfigComponent.SetValue(
			bgLayerEntity,
			*layer,
		)

		assetComponents.SpriteComponent.SetValue(
			bgLayerEntity,
			*assetComponents.NewSprite(0, 0),
		)

	}

}

func AddFrontLayerEntity(scene *coldBrew.Scene, sceneAssetPath string) {
	frontLayerBackgroundEntity := scene.AddEntity(assetComponents.FrontLayerComponent, assetComponents.SpriteComponent)

	assetComponents.FrontLayerComponent.SetValue(
		frontLayerBackgroundEntity,
		*assetComponents.NewFrontLayerConfig(sceneAssetPath),
	)

	assetComponents.SpriteComponent.SetValue(
		frontLayerBackgroundEntity,
		*assetComponents.NewSprite(0, 0),
	)
}
