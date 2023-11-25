package playerCombatUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/kainn9/demo/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

func PlayerIsInvincible(playerState *components.PlayerState) bool {
	return playerState.Combat.Defeated || playerState.Combat.IsInRecoveryIframe || playerState.Transform.Dodging
}

func NpcIsInvincible(npcState *components.NpcState) bool {
	return npcState.Combat.Defeated || !npcState.Combat.Hittable
}

func CreateHitEntity(scene *coldBrew.Scene, name components.CharState, endTick int, hits map[int]bool, initiator, target *donburi.Entry) {
	HitEntity := scene.AddEntity(
		components.HitStateComponent,
	)

	hitState := components.NewHitState(name, endTick, hits, initiator, target)
	components.HitStateComponent.Set(HitEntity, hitState)
}

func RemoveAttackEntity(world donburi.World, targetId int) {

	queries.AttackQuery.Each(world, func(attackEntity *donburi.Entry) {

		attackState := components.AttackDataComponent.Get(attackEntity)

		id := systemsUtil.ID(attackState.Initiator)

		if id != targetId {
			return
		}

		clearAttackFromHits(world, id)
		world.Remove(attackEntity.Entity())

	})

}

func clearAttackFromHits(world donburi.World, id int) {

	query := donburi.NewQuery(
		filter.Or(
			filter.Contains(
				tags.PlayerTag,
			),
			filter.Contains(
				tags.NpcTag,
			),
		),
	)

	query.Each(world, func(entity *donburi.Entry) {

		if entity.HasComponent(components.PlayerStateComponent) {
			playerState := components.PlayerStateComponent.Get(entity)
			delete(playerState.Combat.Hits, id)
			return
		}

		if entity.HasComponent(components.NpcStateComponent) {
			npcState := components.NpcStateComponent.Get(entity)
			delete(npcState.Combat.Hits, id)
			return
		}
	})
}
