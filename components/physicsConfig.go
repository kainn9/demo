package components

import "github.com/yohamta/donburi"

type PhysicsConfig struct {
	GravityCoefficient float64
}

var PhysicsConfigComponent = donburi.NewComponentType[PhysicsConfig]()

func NewPhysicsConfig(gravityCoefficient float64) *PhysicsConfig {
	return &PhysicsConfig{
		GravityCoefficient: gravityCoefficient,
	}
}
