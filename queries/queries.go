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

var FloorQuery = donburi.NewQuery(
	filter.Contains(tags.FloorTag),
)

var PlatformQuery = donburi.NewQuery(
	filter.Contains(tags.PlatformTag),
)

var LadderQuery = donburi.NewQuery(
	filter.Contains(tags.LadderTag),
)

var UISingletonQuery = donburi.NewQuery(
	filter.Contains(tags.UISingletonTag),
)

var AttackQuery = donburi.NewQuery(
	filter.Contains(components.AttackStateComponent),
)

var IndicatorQuery = donburi.NewQuery(
	filter.And(
		filter.Contains(components.IndicatorStateAndConfigComponent),
		filter.Contains(components.RigidBodyComponent),
	),
)
