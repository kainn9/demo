package renderPlayerSystems

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	playerGlobals "github.com/kainn9/demo/globalConfig/player"
	sharedAnimationGlobals "github.com/kainn9/demo/globalConfig/sharedAnimation"
	"github.com/kainn9/demo/queries"
	animUtil "github.com/kainn9/demo/systems/render/util/anim"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

type PlayerRendererSystem struct {
	scene *coldBrew.Scene
}

func NewPlayerRenderer(scene *coldBrew.Scene) *PlayerRendererSystem {
	return &PlayerRendererSystem{
		scene: scene,
	}
}

func (sys PlayerRendererSystem) Query() *donburi.Query {
	return queries.PlayerQuery
}

func (sys PlayerRendererSystem) Draw(screen *ebiten.Image, playerEntity *donburi.Entry) {

	// Get relevant entities and components.
	world := sys.scene.World

	sprites := components.SpritesAnimMapComponent.Get(playerEntity)
	playerBody := components.RigidBodyComponent.Get(playerEntity)
	playerState := components.PlayerStateComponent.Get(playerEntity)

	cameraEntity := systemsUtil.GetCameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)

	sys.updateAnimationState(playerState, sprites)

	currentSpriteSheet := sys.currentSpriteSheet(playerState, sprites)

	opts := sys.configureDrawOptions(playerState, playerBody, currentSpriteSheet, camera)

	sys.renderPlayerSpriteOnCamera(currentSpriteSheet, camera, opts)

}
func (sys PlayerRendererSystem) updateAnimationState(
	playerState *components.PlayerState,
	sprites *map[components.CharState]*components.Sprite,
) {
	// Clearing old animation data if animation state has changed.
	prevAnimationState := playerState.Animation
	currentAnimationState := sys.determinePlayerAnimationState(playerState)

	if prevAnimationState != currentAnimationState {
		animUtil.ResetAnimationConfig((*sprites)[prevAnimationState])
	}

	// Setting current animation state(selecting matching sprite/sheet).
	playerState.Animation = currentAnimationState

}

func (sys PlayerRendererSystem) currentSpriteSheet(

	playerState *components.PlayerState,
	sprites *map[components.CharState]*components.Sprite,

) *components.Sprite {

	currentSpriteSheet := (*sprites)[playerState.Animation]

	// If the current sprite sheet is nil, default to idle.
	if currentSpriteSheet == nil {
		log.Println("Player sprite sheet is nil, defaulting to idle.")

		currentSpriteSheet = (*sprites)[sharedAnimationGlobals.CHAR_STATE_IDLE]
	}

	return currentSpriteSheet
}

func (sys PlayerRendererSystem) configureDrawOptions(

	playerState *components.PlayerState,
	playerBody *tBokiComponents.RigidBody,
	currentSpriteSheet *components.Sprite,
	camera *components.Camera,

) *ebiten.DrawImageOptions {

	opts := &ebiten.DrawImageOptions{}

	// Scaling player sprite to face correct direction.
	opts.GeoM.Scale(playerState.Direction(), 1)

	flashWhite := sys.scene.Manager.TickHandler.CurrentTick()%20 > 15
	if playerState.Combat.IsHit && flashWhite {
		red := color.RGBA{255, 0, 0, 255}
		opts.ColorScale.ScaleWithColor(red)
	} else {
		opts.ColorScale.Reset()
	}

	// Translating player sprite to correct position on camera/screen.
	// Also handles offsetting of sprite onto the players Rigid body.
	xOff := currentSpriteSheet.OffSetX * playerState.Direction()
	yOff := currentSpriteSheet.OffSetY
	xPos := playerBody.Pos.X + xOff
	yPos := playerBody.Pos.Y + yOff

	cameraUtil.Translate(camera, opts, xPos, yPos)

	return opts
}

func (sys PlayerRendererSystem) renderPlayerSpriteOnCamera(currentSpriteSheet *components.Sprite, camera *components.Camera, opts *ebiten.DrawImageOptions) {
	// Selecting correct sprite frame to render.
	spriteAtFrameIndex := animUtil.GetAnimFrame(sys.scene.Manager, currentSpriteSheet)
	// Adding sprite frame to camera.
	cameraUtil.AddImage(camera, spriteAtFrameIndex, opts)

}

func (sys PlayerRendererSystem) determinePlayerAnimationState(playerState *components.PlayerState) components.CharState {

	if playerState.Combat.Defeated {
		return sharedAnimationGlobals.CHAR_STATE_DEFEATED
	}

	if playerState.Combat.IsHit {
		return sharedAnimationGlobals.CHAR_STATE_HURT
	}
	// Todo: Handler func for specific attacks once theres more than one.
	if playerState.Combat.Attacking {
		return sharedAnimationGlobals.CHAR_STATE_ATTACK_PRIMARY
	}

	if playerState.Collision.Climbing && (playerState.Transform.Up || playerState.Transform.Down) {
		return playerGlobals.PLAYER_CHAR_STATE_CLIMB_LADDER_ACTIVE
	}

	if playerState.Collision.Climbing {
		return playerGlobals.PLAYER_CHAR_STATE_CLIMB_LADDER_IDLE
	}

	if playerState.Transform.Jumping {
		return sharedAnimationGlobals.CHAR_STATE_JUMP
	}

	if !playerState.Transform.Jumping && !playerState.Collision.OnGround {

		return sharedAnimationGlobals.CHAR_STATE_FALL
	}

	if playerState.Collision.OnGround && playerState.Transform.BasicHorizontalMovement {
		return sharedAnimationGlobals.CHAR_STATE_RUN
	}

	return sharedAnimationGlobals.CHAR_STATE_IDLE
}
