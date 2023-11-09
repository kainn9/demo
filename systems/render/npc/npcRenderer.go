package renderPlayerSystems

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"

	sharedAnimationGlobals "github.com/kainn9/demo/globalConfig/sharedAnimation"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	"github.com/kainn9/demo/queries"
	animUtil "github.com/kainn9/demo/systems/render/util/anim"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
)

type NpcRendererSystem struct {
	scene *coldBrew.Scene
}

func NewNpcRenderer(scene *coldBrew.Scene) *NpcRendererSystem {
	return &NpcRendererSystem{
		scene: scene,
	}
}

func (sys NpcRendererSystem) Query() *donburi.Query {
	return queries.NpcQuery
}

func (sys NpcRendererSystem) Draw(screen *ebiten.Image, npcEntity *donburi.Entry) {

	// Get relevant entities and components.
	world := sys.scene.World

	sprites := components.SpritesAnimMapComponent.Get(npcEntity)
	state := components.NpcStateComponent.Get(npcEntity)
	body := components.RigidBodyComponent.Get(npcEntity)

	cameraEntity := systemsUtil.GetCameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)

	sys.updateAnimationState(state, sprites)

	currentSpriteSheet := sys.currentSpriteSheet(state, sprites)

	opts := sys.configureDrawOptions(state, body, currentSpriteSheet, camera)

	// Selecting correct sprite frame to render.
	spriteAtFrameIndex := animUtil.GetAnimFrame(sys.scene.Manager, currentSpriteSheet)
	// Adding sprite frame to camera.
	cameraUtil.AddImage(camera, spriteAtFrameIndex, opts)
}

func (sys NpcRendererSystem) updateAnimationState(
	npcState *components.NpcState,
	sprites *map[components.CharState]*components.Sprite,
) {
	// Clearing old animation data if animation state has changed.
	prevAnimationState := npcState.Animation
	currentAnimationState := sys.determineNpcAnimationState(npcState)

	if prevAnimationState != currentAnimationState {
		animUtil.ResetAnimationConfig((*sprites)[prevAnimationState])
	}

	// Setting current animation state(selecting matching sprite/sheet).
	npcState.Animation = currentAnimationState

}

func (sys NpcRendererSystem) determineNpcAnimationState(npcState *components.NpcState) components.CharState {

	if npcState.Combat.Defeated {
		return sharedAnimationGlobals.CHAR_STATE_DEFEATED
	}
	if npcState.Combat.IsHit {
		return sharedAnimationGlobals.CHAR_STATE_HURT
	}

	return sharedAnimationGlobals.CHAR_STATE_IDLE
}

func (sys NpcRendererSystem) currentSpriteSheet(

	npcState *components.NpcState,
	sprites *map[components.CharState]*components.Sprite,

) *components.Sprite {

	currentSpriteSheet := (*sprites)[npcState.Animation]

	// If the current sprite sheet is nil, default to idle.
	if currentSpriteSheet == nil {
		log.Println("Player sprite sheet is nil, defaulting to idle.")

		currentSpriteSheet = (*sprites)[sharedAnimationGlobals.CHAR_STATE_IDLE]
	}

	return currentSpriteSheet
}

func (sys NpcRendererSystem) configureDrawOptions(

	npcState *components.NpcState,
	playerBody *tBokiComponents.RigidBody,
	currentSpriteSheet *components.Sprite,
	camera *components.Camera,

) *ebiten.DrawImageOptions {

	opts := &ebiten.DrawImageOptions{}

	// Scaling player sprite to face correct direction.
	opts.GeoM.Scale(npcState.Direction(), 1)

	// Handling hit flashing options.
	flashWhite := sys.scene.Manager.TickHandler.CurrentTick()%20 > 15
	if npcState.Combat.IsHit && flashWhite {
		opts.ColorScale.SetR(255)
		opts.ColorScale.SetG(255)
		opts.ColorScale.SetB(255)
		opts.ColorScale.SetA(255)
	} else {
		opts.ColorScale.Reset()
	}

	// Translating player sprite to correct position on camera/screen.
	// Also handles offsetting of sprite onto the players Rigid body.
	xOff := currentSpriteSheet.OffSetX * npcState.Direction()
	yOff := currentSpriteSheet.OffSetY
	xPos := playerBody.Pos.X + xOff
	yPos := playerBody.Pos.Y + yOff

	cameraUtil.Translate(camera, opts, xPos, yPos)

	return opts
}
