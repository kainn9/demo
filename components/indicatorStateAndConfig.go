package components

import (
	UIConstants "github.com/kainn9/demo/constants/UI"
	"github.com/yohamta/donburi"
)

type IndicatorStateAndConfig struct {
	// State:
	Active bool

	// Config:
	Type     UIConstants.IndicatorType
	OnPlayer bool
}

var IndicatorStateAndConfigComponent = donburi.NewComponentType[IndicatorStateAndConfig]()

func NewIndicatorStateAndConfig(

	active, onPlayer bool,
	indicatorType UIConstants.IndicatorType,

) *IndicatorStateAndConfig {

	return &IndicatorStateAndConfig{

		Active:   active,
		Type:     indicatorType,
		OnPlayer: onPlayer,
	}
}
