package simSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/demo/components"

	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiMath "github.com/kainn9/tteokbokki/math"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PlayerMovementHandlerSystem struct {
	maxVelX  float64
	xVelUnit float64
}

func NewPlayerMovementHandler() *PlayerMovementHandlerSystem {
	return &PlayerMovementHandlerSystem{
		maxVelX:  180,
		xVelUnit: 18,
	}
}

func (*PlayerMovementHandlerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.InputBufferComponent),
			filter.Contains(components.RigidBodyComponent),
		),
	)
}

func (sys *PlayerMovementHandlerSystem) Run(dt float64, playerEntity *donburi.Entry) {
	inputBuffer := components.InputBufferComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	sys.gravityHandler(playerBody)
	sys.horizontalMovementHandler(inputBuffer, playerBody)
	sys.clampVelocityX(playerBody)

	sys.IntegrateMovementForcesAndClearActiveInput(inputBuffer, playerBody, dt)
}

func (sys *PlayerMovementHandlerSystem) gravityHandler(playerBody *tBokiComponents.RigidBody) {
	weightForce, _ := tBokiPhysics.ForceFactory.NewWeightForce(playerBody.GetMass())
	weightForce.Y = weightForce.Y / 1 // TEMP!
	tBokiPhysics.Transformer.AddForce(playerBody, weightForce)
}

func (sys *PlayerMovementHandlerSystem) horizontalMovementHandler(inputBuffer *components.InputBuffer, playerBody *tBokiComponents.RigidBody) {

	switch inputBuffer.ActiveInput {

	case ebiten.KeyRight:
		sys.handleKeyRight(playerBody)
	case ebiten.KeyLeft:
		sys.handleKeyLeft(playerBody)
	case -1:
		sys.haltPlayerMovement(playerBody)
	}

}

func (sys *PlayerMovementHandlerSystem) handleKeyRight(playerBody *tBokiComponents.RigidBody) {
	if playerBody.Vel.X < 0 {
		sys.haltPlayerMovement(playerBody)
	}

	tBokiPhysics.Transformer.ApplyImpulseLinear(playerBody, tBokiVec.Vec2{X: sys.xVelUnit, Y: 0})
}

func (sys *PlayerMovementHandlerSystem) handleKeyLeft(playerBody *tBokiComponents.RigidBody) {
	if playerBody.Vel.X > 0 {
		sys.haltPlayerMovement(playerBody)
	}

	tBokiPhysics.Transformer.ApplyImpulseLinear(playerBody, tBokiVec.Vec2{X: -sys.xVelUnit, Y: 0})
}

func (sys *PlayerMovementHandlerSystem) haltPlayerMovement(playerBody *tBokiComponents.RigidBody) {
	playerBody.Vel.X = 0
}

func (sys *PlayerMovementHandlerSystem) clampVelocityX(playerBody *tBokiComponents.RigidBody) {
	playerBody.Vel.X = tBokiMath.Clamp(playerBody.Vel.X, -sys.maxVelX, sys.maxVelX)
}

func (sys *PlayerMovementHandlerSystem) IntegrateMovementForcesAndClearActiveInput(inputBuffer *components.InputBuffer, playerBody *tBokiComponents.RigidBody, dt float64) {
	tBokiPhysics.Transformer.Integrate(playerBody, dt)

	if playerBody.Polygon != nil {
		playerBody.UpdateVertices()
	}

	inputBuffer.ActiveInput = -1
}
