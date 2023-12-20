package clientSceneBuilderSystems

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	"github.com/kainn9/demo/queries"
	scenesUtil "github.com/kainn9/demo/scenes/util"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	"github.com/yohamta/donburi"
)

type SceneBuilderSystem struct {
	scene       *coldBrew.Scene
	eventsStack []sceneBuilderEvent

	dragStart tBokiVec.Vec2
	dragEnd   tBokiVec.Vec2

	activeBlock       *donburi.Entry
	activeBlockCached tBokiComponents.RigidBody
}

type eventType int

const (
	blockCreationEventType eventType = iota
	blockMovementEventType eventType = iota
)

type sceneBuilderEvent interface {
	eventType() eventType
}

type blockCreationEvent struct {
	entity *donburi.Entry
}

type blockMovementEvent struct {
	entity               *donburi.Entry
	originalX, originalY float64
}

func (blockCreationEvent) eventType() eventType {
	return blockCreationEventType
}

func (blockMovementEvent) eventType() eventType {
	return blockMovementEventType
}

func NewSceneBuilder(scene *coldBrew.Scene) *SceneBuilderSystem {
	return &SceneBuilderSystem{
		scene: scene,
	}
}

func (sys *SceneBuilderSystem) Sync(_ *donburi.Entry) {
	world := sys.scene.World

	cameraEntity := systemsUtil.CameraEntity(world)
	camera := components.CameraComponent.Get(cameraEntity)

	localX, localY := ebiten.CursorPosition()

	worldX := camera.X + float64(localX)
	worldY := camera.Y + float64(localY)

	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		sys.teleportPlayer(tBokiVec.Vec2{X: worldX, Y: worldY})
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		sys.createBlock(worldX, worldY)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
		sys.undoEvent()
	}

	sys.trackDrag(worldX, worldY)

}

func (sys *SceneBuilderSystem) addEventToStack(event sceneBuilderEvent) {
	sys.eventsStack = append(sys.eventsStack, event)
}

func (sys *SceneBuilderSystem) popEventFromStack() sceneBuilderEvent {
	if len(sys.eventsStack) == 0 {
		return nil
	}

	event := sys.eventsStack[len(sys.eventsStack)-1]
	sys.eventsStack = sys.eventsStack[:len(sys.eventsStack)-1]
	return event
}

func (sys *SceneBuilderSystem) undoEvent() {
	event := sys.popEventFromStack()
	if event == nil {
		log.Println("No events to undo.")
		return
	}

	switch event.eventType() {
	case blockCreationEventType:
		sys.undoBlockCreation(event.(blockCreationEvent))

	case blockMovementEventType:
		sys.undoBlockMovement(event.(blockMovementEvent))
	}
}

func (sys *SceneBuilderSystem) undoBlockCreation(event blockCreationEvent) {
	sys.scene.World.Remove(event.entity.Entity())
}

func (sys *SceneBuilderSystem) undoBlockMovement(event blockMovementEvent) {
	if !sys.scene.World.Valid(event.entity.Entity()) {
		return
	}

	block := components.RigidBodyComponent.Get(event.entity)

	block.Pos = tBokiVec.Vec2{X: event.originalX, Y: event.originalY}

}

func (sys *SceneBuilderSystem) trackDrag(x, y float64) {
	world := sys.scene.World

	// Track drag start.
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		sys.dragStart = tBokiVec.Vec2{X: x, Y: y}

		queries.BlockQuery.Each(world, func(blockEntity *donburi.Entry) {

			block := components.RigidBodyComponent.Get(blockEntity)
			blockCLicked, _ := sys.anyClicks(*block, sys.dragStart)

			if blockCLicked {
				sys.activeBlock = blockEntity
				sys.activeBlockCached = *block

				// Copy polygon, since it's a pointer.
				polyCopy := *block.Polygon
				sys.activeBlockCached.Polygon = &polyCopy

				// Copy vertices, since slices are references.
				vertCount := len(sys.activeBlockCached.Polygon.WorldVertices)

				worldVertCopied := make([]tBokiVec.Vec2, vertCount)
				localVertCopied := make([]tBokiVec.Vec2, vertCount)

				copy(worldVertCopied, sys.activeBlockCached.Polygon.WorldVertices)
				copy(localVertCopied, sys.activeBlockCached.Polygon.LocalVertices)

				sys.activeBlockCached.Polygon.WorldVertices = worldVertCopied
				sys.activeBlockCached.Polygon.LocalVertices = localVertCopied

			}

		})

	}

	// Track drag endpoint.
	sys.dragEnd = tBokiVec.Vec2{X: x, Y: y}

	// Track drag end.
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) && sys.activeBlock != nil {
		cachedBlock := sys.activeBlockCached

		sys.addEventToStack(blockMovementEvent{
			entity:    sys.activeBlock,
			originalX: cachedBlock.Pos.X,
			originalY: cachedBlock.Pos.Y,
		})

		sys.activeBlock = nil
		sys.activeBlockCached = tBokiComponents.RigidBody{}

	}

	sys.applyDrag(sys.activeBlock)

}

func (sys *SceneBuilderSystem) anyClicks(block tBokiComponents.RigidBody, click tBokiVec.Vec2) (bool, tBokiVec.Vec2) {
	centerBlockCLicked := sys.isClicked(block.Pos, sys.dragStart)
	topLeftVertexClicked := sys.isClicked(block.Polygon.WorldVertices[0], sys.dragStart)
	bottomLeftVertexClicked := sys.isClicked(block.Polygon.WorldVertices[3], sys.dragStart)
	topRightVertexClicked := sys.isClicked(block.Polygon.WorldVertices[1], sys.dragStart)
	bottomRightVertexClicked := sys.isClicked(block.Polygon.WorldVertices[2], sys.dragStart)

	if centerBlockCLicked {
		return true, block.Pos
	}

	if topLeftVertexClicked {
		return true, block.Polygon.WorldVertices[0]
	}

	if bottomLeftVertexClicked {
		return true, block.Polygon.WorldVertices[3]
	}

	if topRightVertexClicked {
		return true, block.Polygon.WorldVertices[1]
	}

	if bottomRightVertexClicked {
		return true, block.Polygon.WorldVertices[2]
	}

	return false, tBokiVec.Vec2{}

}

func (sys *SceneBuilderSystem) isClicked(center, click tBokiVec.Vec2) bool {
	dist := center.Sub(click)
	return math.Abs(dist.X) < clientGlobals.SCENE_BUILDER_CLICKER_RADIUS && math.Abs(dist.Y) < clientGlobals.SCENE_BUILDER_CLICKER_RADIUS
}

func (sys *SceneBuilderSystem) applyDrag(blockEntity *donburi.Entry) {
	world := sys.scene.World

	if blockEntity == nil || !world.Valid(blockEntity.Entity()) {
		return
	}

	block := components.RigidBodyComponent.Get(blockEntity)

	isClicked, location := sys.anyClicks(sys.activeBlockCached, sys.dragStart)

	if !isClicked {
		return
	}

	isMovementDrag := location.Compare(sys.activeBlockCached.Pos)
	isTopLeftVertexDrag := location.Compare(sys.activeBlockCached.Polygon.WorldVertices[0])
	isTopRightVertexDrag := location.Compare(sys.activeBlockCached.Polygon.WorldVertices[1])
	isBottomLeftVertexDrag := location.Compare(sys.activeBlockCached.Polygon.WorldVertices[3])
	isBottomRightVertexDrag := location.Compare(sys.activeBlockCached.Polygon.WorldVertices[2])

	if isMovementDrag {
		sys.moveBlock(block)
	}

	if isTopLeftVertexDrag {
		sys.resizeBlock(block, 0, 3)
	}

	if isTopRightVertexDrag {
		sys.resizeBlock(block, 1, 2)
	}

	if isBottomLeftVertexDrag {
		sys.resizeBlock(block, 3, 0)

	}

	if isBottomRightVertexDrag {
		sys.resizeBlock(block, 2, 1)
	}

}

func (sys *SceneBuilderSystem) createBlock(x, y float64) {
	blockEntity := scenesUtil.AddBlockEntity(sys.scene, x, y, 100, 100, 0)

	sys.addEventToStack(blockCreationEvent{
		entity: blockEntity,
	})
}

func (sys *SceneBuilderSystem) moveBlock(block *tBokiComponents.RigidBody) {
	drag := sys.dragEnd.Sub(sys.dragStart)
	block.Pos = sys.activeBlockCached.Pos.Add(drag)
}

func (sys *SceneBuilderSystem) resizeBlock(block *tBokiComponents.RigidBody, targetVertIdxOne, targetVertIdxTwo int) {
	drag := sys.dragEnd.Sub(sys.dragStart)

	block.Polygon.LocalVertices[targetVertIdxOne] = sys.activeBlockCached.Polygon.LocalVertices[targetVertIdxOne].Add(drag)
	block.Polygon.LocalVertices[targetVertIdxTwo].X = block.Polygon.LocalVertices[targetVertIdxOne].X

	block.Polygon.Box = nil                                  // No longer a box.
	block.BroadPhaseSkin = block.Polygon.NewBroadPhaseSkin() // Recalculate BroadPhase skin.
	block.UpdateVertices()
}

func (sys *SceneBuilderSystem) teleportPlayer(pos tBokiVec.Vec2) {

	playerEntity := systemsUtil.PlayerEntity(sys.scene.World)
	player := components.RigidBodyComponent.Get(playerEntity)

	player.Pos = pos
	player.Vel = tBokiVec.Vec2{}

}
