package simPlayerSystems

import (
	"log"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	"github.com/kainn9/demo/queries"
	combatUtil "github.com/kainn9/demo/systems/sim/combatUtil"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	"github.com/yohamta/donburi"
)

type PlayerMeleeAttackHandlerSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerMeleeAttackHandler(scene *coldBrew.Scene) *PlayerMeleeAttackHandlerSystem {
	return &PlayerMeleeAttackHandlerSystem{
		scene: scene,
	}
}

func (sys PlayerMeleeAttackHandlerSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys PlayerMeleeAttackHandlerSystem) Run(dt float64, playerEntity *donburi.Entry) {

	world := sys.scene.World
	tickHandler := sys.scene.Manager.TickHandler

	playerState := components.PlayerStateComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)

	if playerState.Combat.CurrentAttack == "" {
		return
	}

	attackData, ok := playerGlobals.AttackDataMap[playerState.Combat.CurrentAttack]
	if !ok {
		log.Println("key:", playerState.Combat.CurrentAttack)
		panic("attackData is nil!")
	}

	ticksSinceAttackStarted := tickHandler.TicksSinceNTicks(playerState.Combat.AttackStartTick)
	attackIsFinished := ticksSinceAttackStarted >= attackData.TotalTickLength

	if attackIsFinished || playerState.Combat.IsHit || playerState.Transform.Dodging {
		combatUtil.RemoveAttackEntity(world, systemsUtil.ID(playerEntity))
		playerState.Combat.ClearAttackState()
		return
	}

	currentAttackFrame := ticksSinceAttackStarted / attackData.TicksPerFrame

	sys.processAttack(
		playerState.Combat.CurrentAttack,
		currentAttackFrame,
		playerBody,
		playerState,
		playerEntity,
	)
}

func (sys PlayerMeleeAttackHandlerSystem) processAttack(
	attackName components.CharState,
	frame int,
	playerBody *tBokiComponents.RigidBody,
	playerState *components.PlayerState,
	playerEntity *donburi.Entry,
) {

	attackHitboxesData, ok := playerGlobals.AttackHitboxesData[attackName]
	if !ok {
		log.Println("key:", attackName)
		panic("attackHitboxesData is nil!")
	}

	if frame >= len(attackHitboxesData.Hitboxes) {
		return
	}

	currentFrameHitboxesData := attackHitboxesData.Hitboxes[frame]

	currentFrameHitboxes := make([]*tBokiComponents.RigidBody, 0)

	for _, attackBoxData := range currentFrameHitboxesData {

		xPos := playerBody.Pos.X + (attackBoxData.OffsetX * playerState.Direction())
		yPos := playerBody.Pos.Y + attackBoxData.OffsetY

		isAngular := attackBoxData.Rotation != 0

		box := tBokiComponents.NewRigidBodyBox(xPos, yPos, attackBoxData.Width, attackBoxData.Height, 0, isAngular)
		box.Rotation = (attackBoxData.Rotation * playerState.Direction())
		box.UpdateVertices()

		currentFrameHitboxes = append(currentFrameHitboxes, box)
	}

	queries.AttackQuery.Each(sys.scene.World, func(attackEntity *donburi.Entry) {
		attackState := components.AttackDataComponent.Get(attackEntity)

		if attackState.Initiator == playerEntity {
			components.AttackHitboxesComponent.SetValue(attackEntity, currentFrameHitboxes)
		}
	})

}
