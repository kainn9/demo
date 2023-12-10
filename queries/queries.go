package queries

import (
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

var PlayerQuery = donburi.NewQuery(
	filter.And(
		filter.Contains(tags.PlayerTag),
	),
)

var PlayerCarQuery = donburi.NewQuery(
	filter.Contains(
		tags.PlayerCarTag,
	),
)

var NpcQuery = donburi.NewQuery(
	filter.And(
		filter.Contains(tags.NpcTag),
	),
)

var CameraQuery = donburi.NewQuery(
	filter.And(
		filter.Contains(components.CameraComponent),
	),
)

var ParallaxBackGroundLayerQuery = donburi.NewQuery(
	filter.And(
		filter.Contains(components.SpriteComponent),
		filter.Contains(components.ParallaxLayerConfigComponent),
	),
)

var FrontLayerQuery = donburi.NewQuery(
	filter.And(
		filter.Contains(components.SpriteComponent),
		filter.Contains(components.FrontLayerComponent),
	),
)

var BackgroundSoundQuery = donburi.NewQuery(
	filter.Contains(components.BgSoundConfigComponent),
)

var BlockQuery = donburi.NewQuery(
	filter.Contains(tags.BlockTag),
)

var PlatformQuery = donburi.NewQuery(
	filter.Contains(tags.PlatformTag),
)

var PlatformAndBlockQuery = donburi.NewQuery(
	filter.Or(
		filter.Contains(tags.BlockTag),
		filter.Contains(tags.PlatformTag),
	),
)

var LadderQuery = donburi.NewQuery(
	filter.Contains(tags.LadderTag),
)

var UISingletonSpritesQuery = donburi.NewQuery(
	filter.Contains(tags.UISingletonSpritesTag),
)

var UISingletonSoundsQuery = donburi.NewQuery(
	filter.Contains(tags.UISingletonSoundsTag),
)

var AttackQuery = donburi.NewQuery(
	filter.Contains(components.AttackDataComponent),
)

var IndicatorQuery = donburi.NewQuery(
	filter.And(
		filter.Contains(components.IndicatorStateAndConfigComponent),
		filter.Contains(components.RigidBodyComponent),
	),
)

var ChatQuery = donburi.NewQuery(
	filter.Contains(components.ChatStateAndConfigComponent),
)

var HitQuery = donburi.NewQuery(
	filter.Contains(components.HitStateComponent),
)
