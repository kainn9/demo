package loaderUtil

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/kainn9/demo/components"
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
)

func LoadSound(path string, sound *components.Sound) []byte {
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

	songBytes, err := os.ReadFile(concatPath + clientGlobals.SOUND_EXTENSION)
	if err != nil {
		log.Fatalf("Error decoding Song Bytes: %v\n", err)
	}

	s, err := mp3.DecodeWithSampleRate(clientGlobals.SOUND_SAMPLE_RATE, bytes.NewReader(songBytes))
	if err != nil {
		log.Fatalf("Error decoding Song Bytes: %v\n", err)
	}

	b, _ := io.ReadAll(s)
	if err != nil {
		log.Fatal(err)
	}

	sound.Bytes = b
	sound.Loaded = true

	return songBytes
}
