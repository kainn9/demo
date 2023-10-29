package components

import "github.com/yohamta/donburi"

type ChatState struct {
	Active, PopUpMode, PopDownMode bool
	CurrentSlideIndex              int
	PortraitNames                  []string
	ChatName, SceneAssetsPath      string
}

var ChatStateComponent = donburi.NewComponentType[ChatState]()

func NewChatState(chatName, sceneAssetsPath string) *ChatState {
	return &ChatState{
		ChatName:        chatName,
		SceneAssetsPath: sceneAssetsPath,
	}
}
