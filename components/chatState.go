package components

import "github.com/yohamta/donburi"

type ChatStateAndConfig struct {
	// State.
	Active, PopUpMode, PopDownMode bool
	CurrentSlideIndex              int

	// Config.
	SlidesContent                   []SlidesContent
	ChatName                        string
	TextAnimStartTick, TicksPerWord int
}

type SlidesContent struct {
	Text, PortraitName string
}

var ChatStateComponent = donburi.NewComponentType[ChatStateAndConfig]()

func NewChatStateAndConfig(chatName, sceneAssetsPath string, ticksPerWord int, content []SlidesContent) *ChatStateAndConfig {
	return &ChatStateAndConfig{
		ChatName:      chatName,
		SlidesContent: content,

		TicksPerWord: ticksPerWord,
	}
}
