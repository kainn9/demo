package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"

	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiMath "github.com/kainn9/tteokbokki/math"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type PlayerMovementHandlerSystem struct {
	scene *coldBrew.Scene

	// May eventually lift these out into a more shared scope,
	// but for now it seems this is the only system that needs them.
	minVelocity float64
	maxVelY     float64
	maxVelX     float64
	xVelUnit    float64
	yVelUnit    float64
	jumpDelay   int
}

func NewPlayerMovementHandler(scene *coldBrew.Scene) *PlayerMovementHandlerSystem {
	sys := &PlayerMovementHandlerSystem{
		scene: scene,
	}

	sys.minVelocity = 3.0 // The minimum velocity to consider the player moving.
	sys.maxVelY = 750.0   // Max speed up or down.
	sys.maxVelX = 180.0   // Max speed left or right.
	sys.xVelUnit = 18.0   // Left or right.
	sys.yVelUnit = -275.0 // Jump.
	sys.jumpDelay = 15    // The amount of ticks to wait before jumping.

	return sys
}

func (PlayerMovementHandlerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.PlayerStateComponent),
			filter.Contains(components.RigidBodyComponent),
		),
	)
}

func (sys PlayerMovementHandlerSystem) Run(dt float64, playerEntity *donburi.Entry) {
	manager := sys.scene.Manager
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	sys.clampToMinVelocity(playerBody)
	sys.gravityHandler(playerBody, playerState)
	sys.horizontalMovementHandler(playerState, playerBody)
	sys.jumpHandler(playerBody, playerState, manager)
	sys.clampToMaxVelocity(playerBody)
	sys.integrateMovementForces(playerBody, dt)

}

func (sys PlayerMovementHandlerSystem) clampToMinVelocity(playerBody *tBokiComponents.RigidBody) {
	epsilon := sys.minVelocity

	if playerBody.Vel.X > -epsilon && playerBody.Vel.X < epsilon {
		playerBody.Vel.X = 0
	}

	if playerBody.Vel.Y > -epsilon && playerBody.Vel.Y < epsilon {
		playerBody.Vel.Y = 0
	}
}

// Probably worth moving this to a more general gravity system,
// once we have more than one entity that needs gravity.
func (sys PlayerMovementHandlerSystem) gravityHandler(playerBody *tBokiComponents.RigidBody, playerState *components.PlayerState) {

	if playerState.Collision.Climbing {
		return
	}

	weightForce, _ := tBokiPhysics.ForceFactory.NewWeightForce(playerBody.GetMass())
	tBokiPhysics.Transformer.AddForce(playerBody, weightForce)

}

func (sys PlayerMovementHandlerSystem) horizontalMovementHandler(playerState *components.PlayerState, playerBody *tBokiComponents.RigidBody) {

	if playerState.Combat.Attacking {
		sys.haltPlayerMovement(playerBody, playerState)
		return
	}

	if playerState.Transform.BasicHorizontalMovement {
		sys.handlePlayerBasicHorizontalMovement(playerBody, playerState)
	}

	if playerState.Transform.BasicHorizontalMovement == false {
		sys.haltPlayerMovement(playerBody, playerState)
	}

}

func (sys PlayerMovementHandlerSystem) handlePlayerBasicHorizontalMovement(
	playerBody *tBokiComponents.RigidBody,
	playerState *components.PlayerState,
) {
	direction := playerState.Direction()

	// If the player is switching directions, halt velocity first(for more responsive movement).
	// Since direction is either 1(right) or -1(left), we can multiply it by players velocity to
	// check if the direction matches the velocity/movement.
	// Its based on the fact that:
	//------------------------------------
	// negative * negative = positive,
	// positive * positive = positive,
	// negative * positive = negative.
	// -----------------------------------
	// So if the result is negative, they are moving in a opposite direction
	// then the velocity, and we should halt the velocity before applying the opposite.
	if playerBody.Vel.X*direction < 0 {
		sys.haltPlayerMovement(playerBody, playerState)
	}

	tBokiPhysics.Transformer.ApplyImpulseLinear(playerBody, tBokiVec.Vec2{X: sys.xVelUnit * direction, Y: 0})
}

func (sys PlayerMovementHandlerSystem) haltPlayerMovement(playerBody *tBokiComponents.RigidBody, playerState *components.PlayerState) {
	playerBody.Vel.X = 0
	playerState.Transform.BasicHorizontalMovement = false
}

func (sys PlayerMovementHandlerSystem) jumpHandler(playerBody *tBokiComponents.RigidBody, playerState *components.PlayerState, m *coldBrew.Manager) {

	// Exit early if player winding up to jump, but the windup has not finished.
	tickHandler := m.TickHandler

	playerWindupNotFinished := tickHandler.TicksSinceNTicks(playerState.Transform.JumpWindupStart) < sys.jumpDelay

	if playerWindupNotFinished {
		return
	}

	// If player is preparing to jump and the jump
	// windup has finished(guarded above),
	// apply the jump impulse.
	playerPreparingToJump := playerState.Transform.JumpWindupStart != 0
	if playerPreparingToJump {
		playerState.Transform.Jumping = true
		playerState.Transform.JumpWindupStart = 0
		tBokiPhysics.Transformer.ApplyImpulseLinear(playerBody, tBokiVec.Vec2{X: 0, Y: sys.yVelUnit})
	}

	if playerBody.Vel.Y >= 0 {
		playerState.Transform.Jumping = false
	}

}

func (sys PlayerMovementHandlerSystem) clampToMaxVelocity(playerBody *tBokiComponents.RigidBody) {
	playerBody.Vel.X = tBokiMath.Clamp(playerBody.Vel.X, -sys.maxVelX, sys.maxVelX)
	playerBody.Vel.Y = tBokiMath.Clamp(playerBody.Vel.Y, -sys.maxVelY, sys.maxVelY)

}

func (sys PlayerMovementHandlerSystem) integrateMovementForces(playerBody *tBokiComponents.RigidBody, dt float64) {

	tBokiPhysics.Transformer.Integrate(playerBody, dt)

	if playerBody.Polygon != nil {
		playerBody.UpdateVertices()
	}

}
