package loaderUtil

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
)

func LoadImage(path string, sprite *components.Sprite) {
	// Todo: change for build?
	// if BuildTime == "true" {

	// 	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	// 	if err != nil {
	// 		log.Fatalf("Asset Error: %v\n", err)
	// 	}

	// 	path = dir + path
	// 	path = strings.ReplaceAll(path, "./", "/")
	// }

	concatPath := clientGlobals.ASSET_ROOT_PATH + path
	img, _, err := ebitenutil.NewImageFromFile(concatPath + clientGlobals.IMAGE_EXTENSION)
	if err != nil {
		log.Fatal(err)
	}

	sprite.Image = img
	sprite.AssetData.Loaded = true
}
