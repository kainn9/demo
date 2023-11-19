package renderDebugSystems

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	"github.com/kainn9/demo/queries"
	animUtil "github.com/kainn9/demo/systems/render/util/anim"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

type HitBoxPreviewerSystem struct {
	scene       *coldBrew.Scene
	jsonLoaded  bool
	previewData hitboxPreviewConfigFromJson
}

type hitboxPreviewConfigFromJson struct {
	Enabled       bool                 `json:"enabled"`
	AnimationName components.CharState `json:"animationName"`
	Frame         int                  `json:"frame"`
	Hitboxes      []hitboxFromJson     `json:"hitboxes"`
}
type hitboxFromJson struct {
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Rotation float64 `json:"rotation"`
	OffsetX  float64 `json:"offsetX"`
	OffsetY  float64 `json:"offsetY"`
}

// Temp hacky system to help place/position hitboxes with visual feedback.
func NewHitBoxPreviewer(scene *coldBrew.Scene) *HitBoxPreviewerSystem {

	return &HitBoxPreviewerSystem{
		scene: scene,
	}
}

func (HitBoxPreviewerSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys *HitBoxPreviewerSystem) Draw(screen *ebiten.Image, playerEntity *donburi.Entry) {

	if !clientGlobals.DEBUG_MODE {
		return
	}

	world := sys.scene.World
	cameraEntity := systemsUtil.GetCameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)

	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	sprites := components.SpritesCharStateMapComponent.Get(playerEntity)

	if !sys.jsonLoaded || inpututil.IsKeyJustPressed(ebiten.Key8) {
		jsonFile, err := os.Open(clientGlobals.DEBUG_HITBOX_PREVIEW_JSON_PATH)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Successfully loaded Hitbox json!")

		defer jsonFile.Close()

		byteValue, _ := io.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &sys.previewData)

		sys.jsonLoaded = true
	}

	if !sys.previewData.Enabled {
		return
	}

	if !ebiten.IsKeyPressed(ebiten.Key9) {
		return
	}

	currentSpriteSheet := (*sprites)[sys.previewData.AnimationName]

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(playerState.Direction(), 1)
	blue := color.RGBA{0, 0, 255, 255}
	opts.ColorScale.ScaleWithColor(blue)

	xOff := currentSpriteSheet.OffSetX * playerState.Direction()
	yOff := currentSpriteSheet.OffSetY

	cameraUtil.Translate(camera, opts, playerBody.Pos.X+xOff, playerBody.Pos.Y+yOff)

	currFrame := animUtil.PlaySpecificFrame(sys.previewData.Frame, currentSpriteSheet)
	cameraUtil.AddImage(camera, currFrame, opts)

	cameraUtil.Render(camera, screen)

	for _, hitboxData := range sys.previewData.Hitboxes {
		xPos := playerBody.Pos.X + (hitboxData.OffsetX * playerState.Direction())
		yPos := playerBody.Pos.Y + hitboxData.OffsetY

		isAngular := hitboxData.Rotation != 0

		hitbox := tBokiComponents.NewRigidBodyBox(xPos, yPos, hitboxData.Width, hitboxData.Height, 0, isAngular)
		hitbox.Rotation = hitboxData.Rotation * playerState.Direction()
		hitbox.UpdateVertices()

		green := color.RGBA{0, 255, 0, 255}
		debugDrawPolygonBody(screen, camera, *hitbox, green)
	}

}
