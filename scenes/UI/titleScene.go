package UIScenes

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientUISystems "github.com/kainn9/demo/systems/client/UI"
	loaderSystems "github.com/kainn9/demo/systems/loader"
	renderUISystems "github.com/kainn9/demo/systems/render/UI"

	scenesUtil "github.com/kainn9/demo/scenes/util"
)

type TitleScene struct{}

const (
	TITLE_SCENE_WIDTH      = 640
	TITLE_SCENE_HEIGHT     = 360
	TITLE_SCENE_NAME       = "titleScene"
	TITLE_SCENE_SECTION    = "UI"
	TITLE_SCENE_ASSET_PATH = TITLE_SCENE_SECTION + "/" + TITLE_SCENE_NAME + "/"
)

func (TitleScene) Index() string {
	return TITLE_SCENE_NAME
}

func (TitleScene) New(m *coldBrew.Manager) *coldBrew.Scene {

	scene := coldBrew.NewScene(m, TITLE_SCENE_WIDTH, TITLE_SCENE_HEIGHT)

	scene.AddSystem(loaderSystems.NewBackgroundLoader(scene))
	scene.AddSystem(loaderSystems.NewUIGlobalLoader(scene))

	scene.AddSystem(clientUISystems.NewTitleSceneHandler(scene))

	scene.AddSystem(renderUISystems.NewTitleRenderer(scene))
	// Entities ----------------------------------------------------------------------------------

	scenesUtil.AddCameraEntity(scene, 0, 0, 1)

	scenesUtil.AddParallaxBackgroundEntity(scene, []*components.ParallaxLayerConfig{
		components.NewParallaxLayerConfig(TITLE_SCENE_ASSET_PATH, 0, 0, 0, false),
	})

	return scene
}
