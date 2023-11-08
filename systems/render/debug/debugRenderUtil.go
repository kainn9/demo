package renderDebugSystems

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kainn9/demo/components"
	tBokiComponents "github.com/kainn9/tteokbokki/components"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
)

func debugDrawCircleBody(screen *ebiten.Image, camera *components.Camera, rb tBokiComponents.RigidBody, color color.RGBA) {
	x := -camera.X + rb.Pos.X
	y := -camera.Y + rb.Pos.Y

	vector.StrokeCircle(screen, float32(x), float32(y), float32(rb.Circle.Radius), 1.0, color, false)

	endpoint := tBokiVec.Vec2{
		X: x + rb.Circle.Radius,
		Y: y,
	}

	endpoint = endpoint.RotateAroundPoint(rb.Rotation, rb.Pos)

	vector.StrokeLine(screen, float32(x), float32(y), float32(endpoint.X), float32(endpoint.Y), 1.0, color, false)
}

func debugDrawPolygonBody(screen *ebiten.Image, camera *components.Camera, rb tBokiComponents.RigidBody, color color.RGBA) {

	length := len(rb.Polygon.WorldVertices)

	for i := 0; i <= length-1; i++ {
		vert := (rb.Polygon.WorldVertices)[i]

		nextVertIdx := (i + 1) % length

		vert2 := (rb.Polygon.WorldVertices)[nextVertIdx]

		x1 := -camera.X + vert.X
		y1 := -camera.Y + vert.Y

		x2 := -camera.X + vert2.X
		y2 := -camera.Y + vert2.Y

		vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 1.0, color, false)
	}

	x := -camera.X + rb.Pos.X
	y := -camera.Y + rb.Pos.Y

	vector.StrokeCircle(screen, float32(x), float32(y), 4, 1.0, color, false)

}

func debugDrawBroadPhaseSkin(screen *ebiten.Image, camera *components.Camera, rb tBokiComponents.RigidBody, color color.RGBA) {
	x := -camera.X + rb.Pos.X
	y := -camera.Y + rb.Pos.Y

	vector.StrokeCircle(screen, float32(x), float32(y), float32(rb.BroadPhaseSkin.Radius), 1.0, color, false)
}
