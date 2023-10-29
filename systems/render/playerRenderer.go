package renderSystems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
	"github.com/kainn9/demo/queries"
	"github.com/kainn9/demo/systems/render/util/animUtil"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
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

	// Get relevant entities and components.
	world := sys.scene.World

	sprites := assetComponents.SpritesMapComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)
	opts := &ebiten.DrawImageOptions{}

	cameraEntity := systemsUtil.GetCameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)

	// Clearing old animation data if animation state has changed.
	prevAnimationState := playerState.AnimationState
	currentAnimationState := determinePlayerAnimationState(playerState)

	if prevAnimationState != currentAnimationState {
		animUtil.ResetAnimationData((*sprites)[prevAnimationState])
	}

	// Setting current animation state(selecting matching sprite/sheet).
	playerState.AnimationState = currentAnimationState
	currentSpriteSheet := (*sprites)[playerState.AnimationState]

	// If the current sprite sheet is nil, default to idle.
	if currentSpriteSheet == nil {
		currentSpriteSheet = (*sprites)[constants.PLAYER_ANIM_STATE_IDLE]
	}

	// Scaling player sprite to face correct direction.
	opts.GeoM.Scale(playerState.Direction(), 1)

	// Translating player sprite to correct position on camera/screen.
	// Also handles offsetting of sprite onto the players Rigid body.
	xOff := currentSpriteSheet.OffSetX * playerState.Direction()
	yOff := currentSpriteSheet.OffSetY
	xPos := playerBody.Pos.X - xOff
	yPos := playerBody.Pos.Y - yOff

	cameraUtil.Translate(camera, opts, xPos, yPos)

	// Selecting correct sprite frame to render.
	spriteAtFrameIndex := animUtil.GetFrame(sys.scene.Manager, currentSpriteSheet)

	// Adding sprite frame to camera.
	cameraUtil.AddImage(camera, spriteAtFrameIndex, opts)

	// Since this is the last Render System, we render the camera here.
	// This may get moved to a separate system in the future.
	cameraUtil.Render(camera, screen)

}

func determinePlayerAnimationState(playerState *components.PlayerState) string {

	if playerState.Jumping || playerState.JumpWindupStart != 0 {
		return constants.PLAYER_ANIM_STATE_JUMP
	}

	if !playerState.Jumping && !playerState.OnGround {

		return constants.PLAYER_ANIM_STATE_FALL
	}

	if playerState.OnGround && playerState.BasicHorizontalMovement {
		return constants.PLAYER_ANIM_STATE_RUN
	}

	return constants.PLAYER_ANIM_STATE_IDLE
}
