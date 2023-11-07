package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerCarConstants "github.com/kainn9/demo/constants/playerCar"
	"github.com/kainn9/demo/queries"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

type PlayerCarRendererSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerCarRenderer(scene *coldBrew.Scene) *PlayerCarRendererSystem {
	return &PlayerCarRendererSystem{
		scene: scene,
	}
}

func (sys PlayerCarRendererSystem) Query() *donburi.Query {
	return queries.PlayerCarQuery
}

func (sys PlayerCarRendererSystem) Draw(screen *ebiten.Image, playerCarEntity *donburi.Entry) {

	world := sys.scene.World
	sprites := components.SpritesMapComponent.Get(playerCarEntity)
	body := components.RigidBodyComponent.Get(playerCarEntity)

	cameraEntity := systemsUtil.GetCameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)

	opts := &ebiten.DrawImageOptions{}

	allSprite := (*sprites)[playerCarConstants.PLAYER_CAR_ANIM_STATE_ALL]
	xPos := body.Pos.X + allSprite.OffSetX
	yPos := body.Pos.Y + allSprite.OffSetY

	cameraUtil.Translate(camera, opts, xPos, yPos)

	cameraUtil.AddImage(camera, allSprite.Image, opts)

}
