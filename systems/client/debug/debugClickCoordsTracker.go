package debugClientSystems

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	systemsUtil "github.com/kainn9/demo/systems/util"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type DebugClickCoordsTrackerSystem struct {
	scene         *coldBrew.Scene
	coordIndex    int
	coords        []tBokiVec.Vec2
	shapeEntities []*donburi.Entry
}

func NewDebugClickCoordsTracker(scene *coldBrew.Scene) *DebugClickCoordsTrackerSystem {
	return &DebugClickCoordsTrackerSystem{
		scene:      scene,
		coords:     make([]tBokiVec.Vec2, 4),
		coordIndex: 0,
	}
}

func (DebugClickCoordsTrackerSystem) InputsQuery() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.InputsComponent),
	)
}

func (sys *DebugClickCoordsTrackerSystem) Sync(_ *donburi.Entry) {

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) {
		world := sys.scene.World
		cameraEntity := systemsUtil.GetCameraEntity(world)
		camera := components.CameraComponent.Get(cameraEntity)

		localX, localY := ebiten.CursorPosition()

		worldX := camera.X + float64(localX)
		worldY := camera.Y + float64(localY)

		coord := tBokiVec.Vec2{X: worldX, Y: worldY}
		sys.coords[sys.coordIndex%4] = coord

		if sys.coordIndex%4 == 0 {
			sys.clearShapes()
		}
		sys.createVertBody(worldX, worldY)

		if sys.coordIndex%4 == 3 {
			sys.createPolyBody()
			sys.printCoords()
		}

		sys.coordIndex++
	}

}

func (sys *DebugClickCoordsTrackerSystem) printCoords() {

	polyEntity := sys.shapeEntities[len(sys.shapeEntities)-1]
	polyBody := components.RigidBodyComponent.Get(polyEntity)

	log.Println("Debug Box Shape Center Coord:")
	log.Println(polyBody.Pos.X, polyBody.Pos.Y)

	log.Println("Debug Box Shape Local Coords:")
	for _, v := range polyBody.Polygon.LocalVertices {
		log.Println(v)
	}

	log.Println("Debug Box Shape World Coords:")
	for _, v := range polyBody.Polygon.WorldVertices {
		log.Println(v)
	}

}

func (sys *DebugClickCoordsTrackerSystem) clearShapes() {
	for _, shapeEntity := range sys.shapeEntities {
		sys.scene.World.Remove(shapeEntity.Entity())
	}
	sys.shapeEntities = make([]*donburi.Entry, 0)
}

func (sys *DebugClickCoordsTrackerSystem) createVertBody(worldX, worldY float64) {
	vertBody := tBokiComponents.NewRigidBodyCircle(worldX, worldY, 3, 0, false)
	newVertShape := sys.scene.AddEntity(components.RigidBodyComponent)
	sys.shapeEntities = append(sys.shapeEntities, newVertShape)
	components.RigidBodyComponent.SetValue(newVertShape, *vertBody)
}

func (sys *DebugClickCoordsTrackerSystem) createPolyBody() {

	polyBody := tBokiComponents.NewRigidBodyPolygonWorld(0, sys.coords, false)

	newPolyShapeEntity := sys.scene.AddEntity(components.RigidBodyComponent)
	sys.shapeEntities = append(sys.shapeEntities, newPolyShapeEntity)

	components.RigidBodyComponent.SetValue(newPolyShapeEntity, *polyBody)

}
