package callbacksUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	sharedStateGlobals "github.com/kainn9/demo/globalConfig/sharedState"
	clientUISystems "github.com/kainn9/demo/systems/client/UI"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

func AttachChatCallback(scene *coldBrew.Scene, callback clientUISystems.ChatCallBack) {

	for _, sys := range scene.Systems {
		switch cSys := sys.(type) {

		case *clientUISystems.ChatHandlerSystem:

			cSys.CallBackSystems = append(cSys.CallBackSystems, callback)
		}
	}
}

// Common:
type SitCallBackStart struct {
	name  string
	index int
}
type SitCallBackEnd struct {
	name  string
	index int
}

func (cb SitCallBackStart) ChatName() string {
	return cb.name
}

func (cb SitCallBackStart) SlideIndex() int {
	return cb.index
}

func (cb SitCallBackStart) Callback(scene *coldBrew.Scene) {
	playerEntity := systemsUtil.PlayerEntity(scene.World)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerState.Animation = playerGlobals.PLAYER_CHAR_STATE_SIT
}

func (cb SitCallBackEnd) ChatName() string {
	return cb.name
}

func (cb SitCallBackEnd) SlideIndex() int {
	return cb.index
}

func (cb SitCallBackEnd) Callback(scene *coldBrew.Scene) {
	playerEntity := systemsUtil.PlayerEntity(scene.World)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerState.Animation = sharedStateGlobals.CHAR_STATE_IDLE
}

func AttachSitCallbackToChat(scene *coldBrew.Scene, chatName string, slideCount int) {
	AttachChatCallback(scene, SitCallBackStart{chatName, 0})
	AttachChatCallback(scene, SitCallBackEnd{chatName, slideCount})
}
