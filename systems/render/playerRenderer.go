package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/queries"
	cameraUtil "github.com/kainn9/demo/systems/render/util"
	systemsUtil "github.com/kainn9/demo/systems/util"
)

type PlayerRendererSystem struct {
	scene *coldBrew.Scene
}

func (sys *PlayerRendererSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func NewPlayerRenderer(scene *coldBrew.Scene) *PlayerRendererSystem {
	return &PlayerRendererSystem{
		scene: scene,
	}
}

func (sys *PlayerRendererSystem) Draw(screen *ebiten.Image, playerEntity *donburi.Entry) {
	world := sys.scene.World
	sprites := assetComponents.SpriteComponents.Get(playerEntity)

	playerBody := components.RigidBodyComponent.Get(playerEntity)

	sprite := (*sprites)[0]

	camera := systemsUtil.GetCamera(world)

	opts := &ebiten.DrawImageOptions{}

	// Todo: Find a better way to handle offset.
	xPos := playerBody.Pos.X - playerBody.Polygon.Box.Width
	yPos := playerBody.Pos.Y - 48

	cameraUtil.Translate(camera, opts, xPos, yPos)

	cameraUtil.AddImage(camera, sprite.Image, opts)

	cameraUtil.Render(camera, screen)

}
