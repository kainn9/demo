package loaderSceneLoaderSystems

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/kainn9/coldBrew"
	scenesUtil "github.com/kainn9/demo/scenes/util"
	tBokiVec "github.com/kainn9/tteokbokki/math/vec"
	"github.com/yohamta/donburi"
)

type SceneLoaderSystem struct {
	scene *coldBrew.Scene
}

type layerInstanceData struct {
	Identifier string `json:"__identifier"`
	GridSize   int    `json:"__gridSize"`
	IntGridCsv []int  `json:"intGridCsv"`
	Width      int    `json:"__cWid"`
	Height     int    `json:"__cHei"`
}

type levelData struct {
	LayerInstances []layerInstanceData `json:"layerInstances"`
}

const (
	BLOCK_INT = 1
)

func NewSceneLoader(scene *coldBrew.Scene) *SceneLoaderSystem {

	return &SceneLoaderSystem{
		scene: scene,
	}
}

func (sys SceneLoaderSystem) Load(sceneDataEntity *donburi.Entry) {

	collisionData, err := sys.GetLevelDataFromJson()
	if err != nil {
		log.Printf("Error loading scene: %v", err)
		return
	}

	collisionMatrix := sys.collisionMatrix(collisionData)
	log.Println(collisionMatrix)

	sys.create(collisionMatrix, collisionData.GridSize)

}

func (SceneLoaderSystem) GetLevelDataFromJson() (layerInstanceData, error) {
	log.Println("Loading scene from JSON...")
	path := "./LDTK/world/Cemetery.ldtkl"

	// Open the JSON file
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error opening JSON file: %v", err)
		return layerInstanceData{}, err
	}
	defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error reading JSON file: %v", err)
		return layerInstanceData{}, err
	}

	// Unmarshal JSON data into a levelData struct
	var data levelData
	if err := json.Unmarshal(content, &data); err != nil {
		log.Printf("Error unmarshalling JSON data: %v", err)
		return layerInstanceData{}, err
	}

	return data.LayerInstances[0], nil
}

func (SceneLoaderSystem) collisionMatrix(layerData layerInstanceData) [][]int {
	log.Println("Loading collision matrix...")

	width := layerData.Width
	height := layerData.Height

	// Create a 2D array of ints
	collisionMatrix := make([][]int, height)
	for i := range collisionMatrix {
		collisionMatrix[i] = make([]int, width)
	}

	// Fill the 2D array with the collision data
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			collisionMatrix[i][j] = layerData.IntGridCsv[i*width+j]
		}
	}

	return collisionMatrix
}

func (sys SceneLoaderSystem) create(collisionMatrix [][]int, gridSize int) {
	log.Println("Creating blocks...")

	// Create a block for each 1 in the collision matrix
	for i := 0; i < len(collisionMatrix); i++ {
		for j := 0; j < len(collisionMatrix[i]); j++ {
			if collisionMatrix[i][j] == 1 {

				// Find the maximum rectangle size for the current block
				maxX, maxY := findMaxRectangle(collisionMatrix, i, j, gridSize)

				// Create vertices for the merged rectangle
				vert1 := tBokiVec.Vec2{X: float64(j * gridSize), Y: float64(i * gridSize)}
				vert2 := tBokiVec.Vec2{X: float64((j+maxX)*gridSize - 1), Y: float64(i * gridSize)}
				vert3 := tBokiVec.Vec2{X: float64((j+maxX)*gridSize - 1), Y: float64((i+maxY)*gridSize - 1)}
				vert4 := tBokiVec.Vec2{X: float64(j * gridSize), Y: float64((i+maxY)*gridSize - 1)}

				// Create a slice of vertices for the merged rectangle
				vertices := []tBokiVec.Vec2{vert1, vert2, vert3, vert4}

				// Add rigid body to the scene using scenesUtil.AddBlockEntityPoly
				scenesUtil.AddBlockEntityPoly(sys.scene, vertices)

				// Mark the merged blocks as processed
				markAsProcessed(collisionMatrix, i, j, maxX, maxY)

			}
		}
	}
}

// Check if there are obstacles within the rectangle
func hasObstacles(collisionMatrix [][]int, row, col, maxX, maxY int) bool {
	for i := row; i < row+maxY; i++ {
		for j := col; j < col+maxX; j++ {
			if i < len(collisionMatrix) && j < len(collisionMatrix[i]) && collisionMatrix[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

// Find the maximum rectangle size for the current block

func findMaxRectangle(collisionMatrix [][]int, row, col, gridSize int) (maxX, maxY int) {
	maxX, maxY = 0, 0

	// Calculate maxX based on the width of the rectangle
	for j := col; j < len(collisionMatrix[row]) && collisionMatrix[row][j] == 1; j++ {
		maxX++
	}

	// Calculate maxY based on the height of the rectangle
	for i := row; i < len(collisionMatrix) && collisionMatrix[i][col] == 1; i++ {
		maxY++
	}

	// Check if there are obstacles within the rectangle, if so, reduce the rectangle size vertically.
	currX, currY := maxX, maxY

	for hasObstacles(collisionMatrix, row, col, currX, currY) && currY > 1 {
		currY--
	}

	// If vertical reduction is successful, return the new rectangle size.
	if currY > 1 {
		return currX, currY
	}

	// Otherwise, reduce the rectangle size horizontally.

	currX, currY = maxX, maxY

	for hasObstacles(collisionMatrix, row, col, currX, currY) && currX > 1 {
		currX--
	}

	// If horizontal reduction is successful, return the new rectangle size.
	if currX > 1 {
		return currX, currY
	}

	// Otherwise, fatal error.
	log.Fatalf("Error finding max rectangle size for block at (%d, %d)", row, col)

	return 1, 1
}

// Mark the merged blocks as processed
func markAsProcessed(collisionMatrix [][]int, row, col, maxX, maxY int) {
	for i := row; i < row+maxY; i++ {
		for j := col; j < col+maxX; j++ {
			if i < len(collisionMatrix) && j < len(collisionMatrix[i]) {
				collisionMatrix[i][j] = 0
			}
		}
	}
}

// if maxX > 1 || maxY > 1 {
// 	// Create a smaller rectangle to handle the tile without dropping it
// 	vert1 := tBokiVec.Vec2{X: float64(j * gridSize), Y: float64(i * gridSize)}
// 	vert2 := tBokiVec.Vec2{X: float64((j+1)*gridSize - 1), Y: float64(i * gridSize)}
// 	vert3 := tBokiVec.Vec2{X: float64((j+1)*gridSize - 1), Y: float64((i+1)*gridSize - 1)}
// 	vert4 := tBokiVec.Vec2{X: float64(j * gridSize), Y: float64((i+1)*gridSize - 1)}

// 	// Create a slice of vertices for the smaller rectangle
// 	vertices := []tBokiVec.Vec2{vert1, vert2, vert3, vert4}

// 	// Add rigid body to the scene using scenesUtil.AddBlockEntityPoly
// 	scenesUtil.AddBlockEntityPoly(sys.scene, vertices)

// 	// Mark the individual block as processed (not the entire rectangle)
// 	markAsProcessed(collisionMatrix, i, j, 1, 1)
// }
