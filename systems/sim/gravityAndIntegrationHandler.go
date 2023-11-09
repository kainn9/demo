package simSystems

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiPhysics "github.com/kainn9/tteokbokki/physics"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type GravityAndIntegrationHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewGravityAndIntegrationHandler(scene *coldBrew.Scene) *GravityAndIntegrationHandlerSystem {
	return &GravityAndIntegrationHandlerSystem{
		scene: scene,
	}
}

func (sys GravityAndIntegrationHandlerSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.RigidBodyComponent),
	)
}

func (sys GravityAndIntegrationHandlerSystem) Run(dt float64, bodyEntity *donburi.Entry) {
	body := components.RigidBodyComponent.Get(bodyEntity)

	if bodyEntity.HasComponent(components.PlayerStateComponent) {
		state := components.PlayerStateComponent.Get(bodyEntity)
		sys.gravityHandler(body, state)
	} else {
		sys.gravityHandler(body, nil)
	}

	sys.integrateForces(body, dt)
}

// Probably worth moving this to a more general gravity system,
// once we have more than one entity that needs gravity.
func (sys GravityAndIntegrationHandlerSystem) gravityHandler(body *tBokiComponents.RigidBody, playerState *components.PlayerState) {

	if playerState != nil && playerState.Collision.Climbing {
		return
	}

	weightForce, _ := tBokiPhysics.ForceFactory.NewWeightForce(body.GetMass())
	tBokiPhysics.Transformer.AddForce(body, weightForce)

}

func (sys GravityAndIntegrationHandlerSystem) integrateForces(body *tBokiComponents.RigidBody, dt float64) {

	tBokiPhysics.Transformer.Integrate(body, dt)

	if body.Polygon != nil {
		body.UpdateVertices()
	}

}
