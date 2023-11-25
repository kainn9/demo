package components

import "github.com/yohamta/donburi"

type PhysicsModConfig struct {
	GravityCoefficient float64
}

var PhysicsConfigComponent = donburi.NewComponentType[PhysicsModConfig]()

func NewPhysicsModConfig(gravityCoefficient float64) *PhysicsModConfig {
	return &PhysicsModConfig{
		GravityCoefficient: gravityCoefficient,
	}
}
