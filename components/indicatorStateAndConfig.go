package components

import (
	"github.com/yohamta/donburi"
)

// Indicator Constants.
type IndicatorType string

type IndicatorState struct {
	Active bool
}

type IndicatorConfig struct {
	X, Y     float64
	Type     IndicatorType
	OnPlayer bool
}

type IndicatorStateAndConfig struct {
	State  *IndicatorState
	Config *IndicatorConfig
}

var IndicatorStateAndConfigComponent = donburi.NewComponentType[IndicatorStateAndConfig]()

// If onPlayer is true, then x,y is relative to the player,
// otherwise, it global/relative to the scene.
func NewIndicatorStateAndConfig(
	x, y float64,
	active, onPlayer bool,
	indicatorType IndicatorType,

) *IndicatorStateAndConfig {

	return &IndicatorStateAndConfig{

		Config: &IndicatorConfig{
			X:        x,
			Y:        y,
			Type:     indicatorType,
			OnPlayer: onPlayer,
		},

		State: &IndicatorState{
			Active: active,
		},
	}
}
