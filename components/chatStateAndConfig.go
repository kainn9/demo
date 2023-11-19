package components

import (
	"github.com/yohamta/donburi"
)

var ChatStateAndConfigComponent = donburi.NewComponentType[ChatStateAndConfig]()

type ChatStateAndConfig struct {
	Config *ChatConfig
	State  *ChatState
}

type ChatConfig struct {
	ChatName     string
	TicksPerWord int
}

type ChatState struct {
	Active, PopUpMode, PopDownMode                              bool
	CurrentSlideIndex, TextAnimStartTick, NameTextAnimStartTick int
	SlidesContent                                               []SlidesContent
	HasBeenRead                                                 bool
	JustOpened                                                  bool
}

type SlidesContent struct {
	Text, PortraitName, CharName string
	FacingRight                  bool
}

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

func (stateAndConfig *ChatStateAndConfig) Enable() {
	stateAndConfig.State.Active = true
	stateAndConfig.State.PopUpMode = true
	stateAndConfig.State.JustOpened = true
}
