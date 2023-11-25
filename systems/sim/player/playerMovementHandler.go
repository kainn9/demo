package simPlayerSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"

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
	maxVelY     float64
	maxRunVel   float64
	maxDodgeVel float64
	xVelUnit    float64
	yVelUnit    float64
	indoor      bool
}

func NewPlayerMovementHandler(scene *coldBrew.Scene, indoor bool) *PlayerMovementHandlerSystem {
	sys := &PlayerMovementHandlerSystem{
		scene:  scene,
		indoor: indoor,
	}
	sys.maxVelY = 750.0
	sys.maxRunVel = 170.0
	sys.maxDodgeVel = sys.maxRunVel * 2

	sys.xVelUnit = 18.0   // Left or right.
	sys.yVelUnit = -310.0 // Jump.

	if indoor {
		sys.maxRunVel = 120.0
		sys.xVelUnit = 15.0
		sys.yVelUnit = -240.0

	}

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

	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	if playerState.Combat.IsHit {
		return
	}

	sys.dodgeHandler(playerBody, playerState)

	if playerState.Transform.Dodging {
		return
	}

	sys.horizontalMovementHandler(playerState, playerBody)
	sys.jumpHandler(playerBody, playerState)

}

func (sys PlayerMovementHandlerSystem) horizontalMovementHandler(playerState *components.PlayerState, playerBody *tBokiComponents.RigidBody) {

	if playerState.Combat.Attacking {
		sys.haltPlayerMovement(playerBody, playerState)
		return
	}

	if playerState.Transform.BasicHorizontalMovement {
		sys.handlePlayerBasicHorizontalMovement(playerBody, playerState)
	}

	if !playerState.Transform.BasicHorizontalMovement {
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

	// Clamp the player's horizontal velocity.
	playerBody.Vel.X = tBokiMath.Clamp(playerBody.Vel.X, -sys.maxRunVel, sys.maxRunVel)
}

func (sys PlayerMovementHandlerSystem) haltPlayerMovement(playerBody *tBokiComponents.RigidBody, playerState *components.PlayerState) {
	playerBody.Vel.X = 0
	playerState.Transform.BasicHorizontalMovement = false
}

func (sys PlayerMovementHandlerSystem) jumpHandler(playerBody *tBokiComponents.RigidBody, playerState *components.PlayerState) {

	playerJumpTriggeredAndGrounded := playerState.Transform.JumpTriggered && (playerState.Collision.OnGround || playerState.Collision.Climbing)

	if playerJumpTriggeredAndGrounded && !playerState.Transform.Jumping {
		tBokiPhysics.Transformer.ApplyImpulseLinear(playerBody, tBokiVec.Vec2{X: 0, Y: sys.yVelUnit})
		playerState.Transform.Jumping = true
		playerState.Transform.JumpTriggered = false
	}

	if playerBody.Vel.Y >= 0 {
		playerState.Transform.Jumping = false
	}

	// Clamp the player's vertical velocity.
	playerBody.Vel.Y = tBokiMath.Clamp(playerBody.Vel.Y, -sys.maxVelY, sys.maxVelY)

}

func (sys PlayerMovementHandlerSystem) dodgeHandler(playerBody *tBokiComponents.RigidBody, playerState *components.PlayerState) {

	if sys.indoor {
		return
	}

	th := sys.scene.Manager.TickHandler

	if playerState.Transform.DodgeTriggered && !playerState.Transform.Dodging {
		playerState.Transform.DodgeTriggered = false
		playerState.Transform.Dodging = true
		playerState.Transform.DodgeFinishedTick = th.CurrentTick() + playerGlobals.PLAYER_DODGE_DURATION_TICKS
	}

	if playerState.Transform.Dodging {
		tBokiPhysics.Transformer.ApplyImpulseLinear(playerBody, tBokiVec.Vec2{X: sys.maxDodgeVel * playerState.Direction(), Y: 0})
	}

	// Clamp the player's horizontal velocity.
	playerBody.Vel.X = tBokiMath.Clamp(playerBody.Vel.X, -sys.maxDodgeVel, sys.maxDodgeVel)

	if th.CurrentTick() > playerState.Transform.DodgeFinishedTick && playerState.Transform.Dodging {
		playerState.Transform.Dodging = false
		playerBody.Vel.X = 0
	}
}
