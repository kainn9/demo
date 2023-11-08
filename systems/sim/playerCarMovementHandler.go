package simSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/demo/components"
	inputConstants "github.com/kainn9/demo/constants/input"
	"github.com/kainn9/demo/queries"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
)

type PlayerCarMovementHandlerSystem struct {
}

func NewPlayerCarMovementHandler() *PlayerCarMovementHandlerSystem {
	return &PlayerCarMovementHandlerSystem{}
}

func (sys PlayerCarMovementHandlerSystem) Query() *donburi.Query {
	return queries.PlayerCarQuery
}

// Todo: BREAK THIS UP AND ADD FRICTION!
func (sys PlayerCarMovementHandlerSystem) Run(dt float64, playerCarEntity *donburi.Entry) {
	body := components.RigidBodyComponent.Get(playerCarEntity)

	_, _, jump, _, _, _, _ := inputConstants.ALL_BINDS()

	if ebiten.IsKeyPressed(jump) {
		tBokiPhysics.Transformer.ApplyImpulseLinear(body, tBokiVec.Vec2{X: 10, Y: 0})
	}

	if body.Vel.X > 400 {
		body.Vel.X = 400
	}

	tBokiPhysics.Transformer.Integrate(body, dt)

	if body.Polygon != nil {
		body.UpdateVertices()
	}

}
