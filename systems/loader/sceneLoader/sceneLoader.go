package loaderSceneLoaderSystems

import (
	"fmt"
	"log"
	"os"

	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"google.golang.org/protobuf/proto"
)

const (
	BLOCK_BIN_FILE_NAME = "block.bin"
)

type SceneLoaderSystem struct{}

func NewSceneLoader(scene *coldBrew.Scene) *SceneLoaderSystem {
	return &SceneLoaderSystem{}
}

func (sys SceneLoaderSystem) Query() *donburi.Query {
	return donburi.NewQuery(
		filter.Contains(components.SceneDataComponent),
	)
}

func (sys SceneLoaderSystem) Load(sceneDataEntity *donburi.Entry) {
	log.Println("Loading scene...")

	sceneData := components.SceneDataComponent.Get(sceneDataEntity)

	sys.LoadBlocksFromBin(sceneData.AssetPath, sceneData)
}

func (sys SceneLoaderSystem) LoadBlocksFromBin(assetPath string, sceneData *components.SceneData) {
	filePath := clientGlobals.ASSET_ROOT_PATH + clientGlobals.SCENE_ASSETS_SUB_PATH + assetPath + BLOCK_BIN_FILE_NAME

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// File doesn't exist, create a new blank bin file
		err := os.WriteFile(filePath, nil, 0644)
		if err != nil {
			panic(err)
		}

		fmt.Println("New blank file created:", filePath)
	} else {
		// File exists, load binary data from the file
		data, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		// Deserialize binary data
		if err := proto.Unmarshal(data, sceneData.TerrainBlocks); err != nil {
			panic(err)
		}

		// Now newScene contains the deserialized data
		fmt.Printf("Deserialized data: %+v\n", sceneData)
	}
}
