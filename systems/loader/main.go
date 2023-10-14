package loaderSystems

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	assetComponents "github.com/kainn9/demo/components/assets"
	"github.com/kainn9/demo/constants"
)

func LoadImage(path string, sprite *assetComponents.Sprite) {
	// Todo: change for build?
	// if BuildTime == "true" {

	// 	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	// 	if err != nil {
	// 		log.Fatalf("Asset Error: %v\n", err)
	// 	}

	// 	path = dir + path
	// 	path = strings.ReplaceAll(path, "./", "/")
	// }

	concatPath := constants.IMAGE_ASSET_DIR + path
	img, _, err := ebitenutil.NewImageFromFile(concatPath)
	if err != nil {
		log.Fatal(err)
	}

	sprite.Image = img
	sprite.AssetData.Loaded = true
}
