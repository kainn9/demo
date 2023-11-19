package clientUISystems

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	inputGlobals "github.com/kainn9/demo/globalConfig/input"
	scenesUtil "github.com/kainn9/demo/scenes/util"

	"github.com/yohamta/donburi"
)

type TitleSceneHandlerSystem struct {
	scene         *coldBrew.Scene
	nextSceneFace coldBrew.SceneFace
}

func NewTitleSceneHandler(scene *coldBrew.Scene, nextSceneFace coldBrew.SceneFace) *TitleSceneHandlerSystem {
	return &TitleSceneHandlerSystem{
		scene:         scene,
		nextSceneFace: nextSceneFace,
	}
}

func (sys *TitleSceneHandlerSystem) Sync(_ *donburi.Entry) {

	_, _, _, _, _, interact, _ := inputGlobals.ALL_BINDS()

	if inpututil.IsKeyJustPressed(interact) {
		scenesUtil.ChangeScene(sys.scene.Manager, sys.nextSceneFace, 147, 275, 0, 0)
	}

}
