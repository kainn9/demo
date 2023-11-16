package clientUISystems

import (
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	inputGlobals "github.com/kainn9/demo/globalConfig/input"
	introScenes "github.com/kainn9/demo/scenes/intro"
	scenesUtil "github.com/kainn9/demo/scenes/util"

	"github.com/yohamta/donburi"
)

type TitleSceneHandlerSystem struct {
	scene                   *coldBrew.Scene
	tickLeftKeyLastPressed  int
	tickRightKeyLastPressed int
}

func NewTitleSceneHandler(scene *coldBrew.Scene) *TitleSceneHandlerSystem {
	return &TitleSceneHandlerSystem{
		scene: scene,
	}
}

func (sys *TitleSceneHandlerSystem) Sync(_ *donburi.Entry) {

	_, _, _, _, _, interact, _ := inputGlobals.ALL_BINDS()

	if inpututil.IsKeyJustPressed(interact) {
		scenesUtil.ChangeScene(sys.scene.Manager, introScenes.LevelOneScene{}, 147, 275, 0, 0)
	}

}
