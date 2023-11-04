package components

import (
	"github.com/kainn9/coldBrew"
	"github.com/yohamta/donburi"
)

type SceneTransitionState struct {
	// Not used yet, but intend to use for a potential fade out effect or
	// to wait for transition animations to finish in the future...
	Active                bool
	ActivatedTick         int
	TransitionTimeInTicks int
}

type SceneTransitionConfig struct {
	TargetScene    coldBrew.SceneFace
	SpawnX, SpawnY float64
	CamX, CamY     float64
	ClickBased     bool
}

type SceneTransitionStateAndConfig struct {
	State  *SceneTransitionState
	Config *SceneTransitionConfig
}

var SceneTransitionStateAndConfigComponent = donburi.NewComponentType[SceneTransitionStateAndConfig]()

func NewSceneTransitionStateAndConfig(spawnX, spawnY, camX, camY float64, targetScene coldBrew.SceneFace, clickBased bool) *SceneTransitionStateAndConfig {
	return &SceneTransitionStateAndConfig{
		State: &SceneTransitionState{},
		Config: &SceneTransitionConfig{
			SpawnX:      spawnX,
			SpawnY:      spawnY,
			CamX:        camX,
			CamY:        camY,
			TargetScene: targetScene,
			ClickBased:  clickBased,
		},
	}
}
