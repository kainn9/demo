package components

import "github.com/yohamta/donburi"

type ChatConfig struct {
	ChatName     string
	TicksPerWord int
}

type ChatState struct {
	Active, PopUpMode, PopDownMode                              bool
	CurrentSlideIndex, TextAnimStartTick, NameTextAnimStartTick int
	SlidesContent                                               []SlidesContent
	HasBeenRead                                                 bool
}

type ChatStateAndConfig struct {
	Config *ChatConfig
	State  *ChatState
}

type SlidesContent struct {
	Text, PortraitName, CharName string
	FacingRight                  bool
}

var ChatStateAndConfigComponent = donburi.NewComponentType[ChatStateAndConfig]()

func NewChatStateAndConfig(chatName string, content []SlidesContent) *ChatStateAndConfig {
	return &ChatStateAndConfig{

		Config: &ChatConfig{
			ChatName:     chatName,
			TicksPerWord: 15,
		},

		State: &ChatState{
			TextAnimStartTick:     -1,
			NameTextAnimStartTick: -1,
			SlidesContent:         content,
		},
	}
}
