package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
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

func AddCameraEntity(scene *coldBrew.Scene, x, y float64) {

	cameraEntity := scene.AddEntity(
		components.CameraComponent,
	)

	components.CameraComponent.SetValue(
		cameraEntity,
		*components.NewCamera(x, y, constants.SCREEN_WIDTH, constants.SCREEN_HEIGHT),
	)

}

func AddChatEntity(scene *coldBrew.Scene, slidesCount int, introChat, sceneAssetPath string, portraitNames []string) {

	chatEntity := scene.AddEntity(
		components.ChatStateComponent,
		assetComponents.SpritesSliceComponent,
		assetComponents.SpritesMapComponent,
	)

	chatState := components.NewChatState(introChat, sceneAssetPath)
	chatState.Active = true
	chatState.PopUpMode = true
	chatState.PortraitNames = portraitNames

	slideSprites := make([]*assetComponents.Sprite, slidesCount)
	portraits := make(map[string]*assetComponents.Sprite, 0)

	for i := 0; i < slidesCount; i++ {
		slideSprites[i] = assetComponents.NewSprite(0, 0)
		portraits[portraitNames[i]] = assetComponents.NewSprite(0, 0)
	}

	assetComponents.SpritesMapComponent.SetValue(
		chatEntity,
		portraits,
	)

	assetComponents.SpritesSliceComponent.SetValue(
		chatEntity,
		slideSprites,
	)

	components.ChatStateComponent.SetValue(
		chatEntity,
		*chatState,
	)

}
