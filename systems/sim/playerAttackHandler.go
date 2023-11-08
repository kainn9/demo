package simSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerConstants "github.com/kainn9/demo/constants/player"
	"github.com/kainn9/demo/queries"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	"github.com/yohamta/donburi"
)

type PlayerAttackHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerAttackHandler(scene *coldBrew.Scene) *PlayerAttackHandlerSystem {
	return &PlayerAttackHandlerSystem{
		scene: scene,
	}
}

func (sys PlayerAttackHandlerSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys PlayerAttackHandlerSystem) Run(dt float64, playerEntity *donburi.Entry) {

	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	tickHandler := sys.scene.Manager.TickHandler

	if playerState.Combat.CurrentAttack == "" {
		return
	}

	currentAttackData := playerConstants.PlayerAttackDataMap[playerState.Combat.CurrentAttack]

	if currentAttackData == nil {
		log.Println(playerState.Combat.CurrentAttack)
		panic("currentAttackData is nil!")
	}

	ticksSinceAttackStart := tickHandler.TicksSinceNTicks(playerState.Combat.AttackStartTick)

	if ticksSinceAttackStart >= currentAttackData.TotalTickLength {
		playerState.Combat.Attacking = false
		playerState.Combat.AttackStartTick = -1
		playerState.Combat.CurrentAttack = ""

		emptyBoxes := []*tBokiComponents.RigidBody{
			tBokiComponents.NewRigidBodyBox(0, 0, 0, 0, 0, false),
		}

		components.PlayerAttackBoxesComponent.SetValue(playerEntity, emptyBoxes)
		return
	}

	hitboxesData := components.AttackHitboxConfigComponent.Get(playerEntity)

	frame := ticksSinceAttackStart / currentAttackData.TicksPerFrame

	if frame > len(hitboxesData.Hitboxes)-1 {
		return
	}

	currBoxesData := hitboxesData.Hitboxes[frame]
	hitboxes := make([]*tBokiComponents.RigidBody, 0)

	for _, boxData := range currBoxesData {
		xPos := playerBody.Pos.X + (boxData.OffsetX * playerState.Direction())
		yPos := playerBody.Pos.Y + boxData.OffsetY

		isAngular := boxData.Rotation != 0

		box := tBokiComponents.NewRigidBodyBox(xPos, yPos, boxData.Width, boxData.Height, 0, isAngular)
		box.Rotation = (boxData.Rotation * playerState.Direction())
		box.UpdateVertices()

		hitboxes = append(hitboxes, box)
	}

	components.PlayerAttackBoxesComponent.SetValue(playerEntity, hitboxes)
}
