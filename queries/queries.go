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

var Block = donburi.NewQuery(
	filter.And(
		filter.Contains(tags.StandardCollisionTag),
		filter.Not(filter.Contains(tags.PlayerTag)),
		filter.Contains(components.RigidBodyComponent),
	),
)

var ChatPopUpEntityQuery = donburi.NewQuery(
	filter.Contains(tags.ChatPopUpSpriteTag),
)

var ChatPopDownEntityQuery = donburi.NewQuery(
	filter.Contains(tags.ChatPopDownSpriteTag),
)
