package components

import "github.com/yohamta/donburi"

type RigidBody struct {
	X float64
	Y float64
}

var RigidBodyComponent = donburi.NewComponentType[RigidBody]()

func NewRigidBody(x, y float64) *RigidBody {
	return &RigidBody{
		X: x,
		Y: y,
	}
}
