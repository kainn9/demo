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

	sys.createBlocks(collisionMatrix, collisionData.GridSize)

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

func (sys SceneLoaderSystem) createBlocks(collisionMatrix [][]int, gridSize int) {
	log.Println("Creating blocks...")

	// Create a block for each 1 in the collision matrix
	for i := 0; i < len(collisionMatrix); i++ {
		for j := 0; j < len(collisionMatrix[i]); j++ {

			current := collisionMatrix[i][j]

			if current == 0 {
				continue
			}

			leftOfCurrentIsBlank := j == 0 || collisionMatrix[i][j-1] == 0
			topOfCurrentIsBlank := i == 0 || collisionMatrix[i-1][j] == 0

			currentIsValidStartPoint := leftOfCurrentIsBlank && topOfCurrentIsBlank

			if currentIsValidStartPoint {
				vertices := sys.findVertices(i, j, gridSize, collisionMatrix)
				scenesUtil.AddBlockEntityPoly(sys.scene, vertices)
			}
		}
	}
}

func (sys SceneLoaderSystem) findVertices(row, col, gridSize int, collisionMatrix [][]int) []tBokiVec.Vec2 {
	vertices := []tBokiVec.Vec2{}

	topLeftVert := tBokiVec.Vec2{
		X: float64(col * gridSize),
		Y: float64(row * gridSize),
	}

	topRightVert, rCol := sys.rightVert(row, col, gridSize, collisionMatrix)
	bottomRightVert := sys.bottomVert(false, row, rCol, gridSize, collisionMatrix)
	bottomLeftVert := sys.bottomVert(true, row, col, gridSize, collisionMatrix)

	vertices = append(vertices, topLeftVert)
	vertices = append(vertices, topRightVert)
	vertices = append(vertices, bottomRightVert)
	vertices = append(vertices, bottomLeftVert)

	log.Println("Top right vert: ", topRightVert, "|", "Bottom right vert: ", bottomRightVert)

	return vertices
}

func (sys SceneLoaderSystem) bottomVert(left bool, i, j, gridSize int, collisionMatrix [][]int) (cord tBokiVec.Vec2) {
	for k := i; k < len(collisionMatrix); k++ {
		colFinished := k == len(collisionMatrix)-1
		currentIsABlock := collisionMatrix[k][j] == BLOCK_INT

		x := j + 1

		if left {
			x = j
		}

		if !currentIsABlock || colFinished {
			return tBokiVec.Vec2{
				X: float64(x * gridSize),
				Y: float64(k * gridSize),
			}
		}
	}

	log.Println("Error finding bottom vert")
	return tBokiVec.Vec2{}
}

func (sys SceneLoaderSystem) rightVert(i, j, gridSize int, collisionMatrix [][]int) (cord tBokiVec.Vec2, col int) {
	for k := j; k < len(collisionMatrix[i]); k++ {

		rowFinished := k == len(collisionMatrix[i])-1
		currentIsNotABlock := collisionMatrix[i][k] != BLOCK_INT

		if currentIsNotABlock || rowFinished {
			return tBokiVec.Vec2{
				X: float64(k * gridSize),
				Y: float64(i * gridSize),
			}, k - 1
		}
	}
	log.Println("Error finding right vert")
	return tBokiVec.Vec2{}, 0
}
