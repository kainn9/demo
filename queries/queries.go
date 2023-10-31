package queries

import (
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

var PlayerQuery = donburi.NewQuery(
	filter.And(
		filter.Contains(tags.PlayerTag),
	),
)

var CameraQuery = donburi.NewQuery(
	filter.And(
		filter.Contains(components.CameraComponent),
	),
)

var ParallaxBackGroundLayerQuery = donburi.NewQuery(
	filter.And(
		filter.Contains(assetComponents.SpriteComponent),
		filter.Contains(assetComponents.ParallaxLayerConfigComponent),
	),
)

var FrontLayerQuery = donburi.NewQuery(
	filter.And(
		filter.Contains(assetComponents.SpriteComponent),
		filter.Contains(assetComponents.FrontLayerComponent),
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

var ChatPopUpAnimQuery = donburi.NewQuery(
	filter.Contains(tags.ChatPopUpSpriteTag),
)

var ChatPopDownAnimQuery = donburi.NewQuery(
	filter.Contains(tags.ChatPopDownSpriteTag),
)
