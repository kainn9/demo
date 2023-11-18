package components

import "github.com/yohamta/donburi"

var AttackHitboxConfigComponent = donburi.NewComponentType[AttackHitboxConfig]()

type AttackHitboxConfig struct {
	Hitboxes []HitBoxDataCluster
}

type HitboxData struct {
	Width, Height, Rotation, OffsetX, OffsetY float64
}

type HitBoxDataCluster []HitboxData

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
