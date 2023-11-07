package scenesUtil

import (
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerCarConstants "github.com/kainn9/demo/constants/playerCar"
	"github.com/kainn9/demo/tags"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

func AddCarEntity(scene *coldBrew.Scene, x, y float64) tBokiComponents.RigidBody {

	// Entity Initialization.
	playerEntity := scene.AddEntity(
		components.RigidBodyComponent,
		components.SpritesMapComponent,
		tags.PlayerCarTag,
	)

	// RigidBody.
	playerBody := *tBokiComponents.NewRigidBodyBox(x, y, playerCarConstants.PLAYER_CAR_WIDTH, playerCarConstants.PLAYER_CAR_HEIGHT, 1, false)

	components.RigidBodyComponent.SetValue(playerEntity, playerBody)

	// Sprites/Animations.
	playerSprites := make(map[string]*components.Sprite, 0)

	// All for now.
	playerSprites[playerCarConstants.PLAYER_CAR_ANIM_STATE_ALL] = components.NewSprite(
		playerCarConstants.PLAYER_CAR_SPRITE_OFFSET_X,
		playerCarConstants.PLAYER_CAR_SPRITE_OFFSET_Y,
	)

	components.SpritesMapComponent.SetValue(playerEntity, playerSprites)

	return playerBody
}
