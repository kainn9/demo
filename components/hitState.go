package components

import "github.com/yohamta/donburi"

var HitStateComponent = donburi.NewComponentType[HitState]()

// Do I store the NPC BODY?....Yes.
// Also were gonna refactor tbOKI to use RB pointers too!
// I've changed my mind!!!!!!!

type HitState struct {
	AttackName CharState
	Hits       map[int]bool
	EndTick    int

	Initiator, Target *donburi.Entry
}

func NewHitState(name CharState, endTick int, hits map[int]bool, initiator, target *donburi.Entry) *HitState {
	return &HitState{
		AttackName: name,
		Initiator:  initiator,
		Target:     target,
		EndTick:    endTick,
		Hits:       hits,
	}
}

func (hs *HitState) HitCachingDisabled() bool {
	return hs.Hits == nil
}

func (hs *HitState) HitCachingEnabled() bool {
	return !hs.HitCachingDisabled()
}
