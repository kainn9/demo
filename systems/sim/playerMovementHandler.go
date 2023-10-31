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
}

var (
	minVelocity = 3.0    // The minimum velocity to consider the player moving.
	maxVelY     = 750.0  // Max speed up or down.
	maxVelX     = 180.0  // Max speed left or right.
	xVelUnit    = 18.0   // Left or right.
	yVelUnit    = -275.0 // Jump.
	jumpDelay   = 15     // The amount of ticks to wait before jumping.
)

func NewPlayerMovementHandler(scene *coldBrew.Scene) *PlayerMovementHandlerSystem {
	return &PlayerMovementHandlerSystem{
		scene: scene,
	}
}

func (*PlayerMovementHandlerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.And(
			filter.Contains(components.PlayerStateComponent),
			filter.Contains(components.RigidBodyComponent),
		),
	)
}

func (sys *PlayerMovementHandlerSystem) Run(dt float64, playerEntity *donburi.Entry) {
	manager := sys.scene.Manager
	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	clampToMinVelocity(playerBody)
	gravityHandler(playerBody, playerState)
	horizontalMovementHandler(playerState, playerBody)
	jumpHandler(playerBody, playerState, manager)
	clampToMaxVelocity(playerBody)
	IntegrateMovementForces(playerBody, dt)

}

func clampToMinVelocity(playerBody *tBokiComponents.RigidBody) {
	epsilon := minVelocity

	if playerBody.Vel.X > -epsilon && playerBody.Vel.X < epsilon {
		playerBody.Vel.X = 0
	}

	if playerBody.Vel.Y > -epsilon && playerBody.Vel.Y < epsilon {
		playerBody.Vel.Y = 0
	}
}

// Probably worth moving this to a more general gravity system,
// once we have more than one entity that needs gravity.
func gravityHandler(playerBody *tBokiComponents.RigidBody, playerState *components.PlayerState) {

	if playerState.Climbing {
		return
	}

	weightForce, _ := tBokiPhysics.ForceFactory.NewWeightForce(playerBody.GetMass())
	tBokiPhysics.Transformer.AddForce(playerBody, weightForce)

}

func horizontalMovementHandler(playerState *components.PlayerState, playerBody *tBokiComponents.RigidBody) {

	if playerState.BasicHorizontalMovement {
		handlePlayerBasicHorizontalMovement(playerBody, playerState)
	}

	if playerState.BasicHorizontalMovement == false {
		haltPlayerMovement(playerBody, playerState)
	}

}

func handlePlayerBasicHorizontalMovement(
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
		haltPlayerMovement(playerBody, playerState)
	}

	tBokiPhysics.Transformer.ApplyImpulseLinear(playerBody, tBokiVec.Vec2{X: xVelUnit * direction, Y: 0})
}

func haltPlayerMovement(playerBody *tBokiComponents.RigidBody, playerState *components.PlayerState) {
	playerBody.Vel.X = 0
	playerState.BasicHorizontalMovement = false
}

func jumpHandler(playerBody *tBokiComponents.RigidBody, playerState *components.PlayerState, m *coldBrew.Manager) {

	// Exit early if player winding up to jump, but the windup has not finished.
	tickHandler := m.TickHandler

	playerWindupNotFinished := tickHandler.TicksSinceNTicks(playerState.JumpWindupStart) < jumpDelay

	if playerWindupNotFinished {
		return
	}

	// If player is preparing to jump and the jump
	// windup has finished(guarded above),
	// apply the jump impulse.
	playerPreparingToJump := playerState.JumpWindupStart != 0
	if playerPreparingToJump {
		playerState.Jumping = true
		playerState.JumpWindupStart = 0
		tBokiPhysics.Transformer.ApplyImpulseLinear(playerBody, tBokiVec.Vec2{X: 0, Y: yVelUnit})
	}

	if playerBody.Vel.Y >= 0 {
		playerState.Jumping = false
	}

}

func clampToMaxVelocity(playerBody *tBokiComponents.RigidBody) {
	playerBody.Vel.X = tBokiMath.Clamp(playerBody.Vel.X, -maxVelX, maxVelX)
	playerBody.Vel.Y = tBokiMath.Clamp(playerBody.Vel.Y, -maxVelY, maxVelY)

}

func IntegrateMovementForces(playerBody *tBokiComponents.RigidBody, dt float64) {

	tBokiPhysics.Transformer.Integrate(playerBody, dt)

	if playerBody.Polygon != nil {
		playerBody.UpdateVertices()
	}

}
