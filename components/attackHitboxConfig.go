package components

import "github.com/yohamta/donburi"

type HitboxData struct {
	Width, Height, Rotation, OffsetX, OffsetY float64
}

type AttackHitboxConfig struct {
	Hitboxes []HitBoxDataCluster
}

type HitBoxDataCluster []HitboxData

var AttackHitboxConfigComponent = donburi.NewComponentType[AttackHitboxConfig]()

func NewHitboxData(width, height, rotation, offsetX, offsetY float64) HitboxData {
	return HitboxData{
		Width:    width,
		Height:   height,
		OffsetX:  offsetX,
		OffsetY:  offsetY,
		Rotation: rotation,
	}
}

func NewAttackHitboxConfig(cluster ...HitBoxDataCluster) *AttackHitboxConfig {

	return &AttackHitboxConfig{
		Hitboxes: cluster,
	}
}