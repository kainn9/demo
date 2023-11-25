package sharedCombatGlobals

type AttackHitboxesData struct {
	Hitboxes []AttackHitboxesFrameData
}

type AttackHitboxesFrameData []AttackHitboxData

type AttackHitboxData struct {
	Width, Height, Rotation, OffsetX, OffsetY float64
}

type AttackData struct {
	TotalTickLength int
	TicksPerFrame   int
}

const (
	IS_HIT_DURATION_IN_TICKS = 30
)

var EMPTY_HITBOXES_FRAME = AttackHitboxesFrameData{}

func NewAttackHitboxData(width, height, rotation, offsetX, offsetY float64) AttackHitboxData {
	return AttackHitboxData{
		Width:    width,
		Height:   height,
		OffsetX:  offsetX,
		OffsetY:  offsetY,
		Rotation: rotation,
	}
}

func NewAttackHitboxesFrameData(attackHitboxData ...AttackHitboxData) AttackHitboxesFrameData {
	return attackHitboxData
}

func NewAttackHitboxesData(attackHitboxesFrameData ...AttackHitboxesFrameData) AttackHitboxesData {
	return AttackHitboxesData{
		Hitboxes: attackHitboxesFrameData,
	}
}
