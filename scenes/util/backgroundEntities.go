package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
)

func AddParallaxBackgroundEntity(scene *coldBrew.Scene, layers []*components.ParallaxLayerConfig) {

	for _, layer := range layers {

		bgLayerEntity := scene.AddEntity(components.ParallaxLayerConfigComponent, components.SpriteComponent)

		components.ParallaxLayerConfigComponent.SetValue(
			bgLayerEntity,
			*layer,
		)

		components.SpriteComponent.SetValue(
			bgLayerEntity,
			*components.NewSprite(0, 0),
		)

	}

}

func AddFrontLayerEntity(scene *coldBrew.Scene, sceneAssetPath string) {
	frontLayerBackgroundEntity := scene.AddEntity(components.FrontLayerComponent, components.SpriteComponent)

	components.FrontLayerComponent.SetValue(
		frontLayerBackgroundEntity,
		*components.NewFrontLayerConfig(sceneAssetPath),
	)

	components.SpriteComponent.SetValue(
		frontLayerBackgroundEntity,
		*components.NewSprite(0, 0),
	)
}

func AddBgSoundEntity(scene *coldBrew.Scene, sceneAssetPath string) {
	bgSoundEntity := scene.AddEntity(components.BgSoundConfigComponent, components.SoundComponent)

	components.BgSoundConfigComponent.SetValue(
		bgSoundEntity,
		*components.NewBgSoundConfig(sceneAssetPath),
	)

	components.SoundComponent.SetValue(
		bgSoundEntity,
		*components.NewSound(-1, 1),
	)
}
