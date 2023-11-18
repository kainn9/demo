package loaderUtil

import "os"

func LoadMusic(path string) ([]byte, error) {
	// Todo: change for build?
	// if BuildTime == "true" {

	// 	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	// 	if err != nil {
	// 		log.Fatalf("Asset Error: %v\n", err)
	// 	}

	// 	path = dir + path
	// 	path = strings.ReplaceAll(path, "./", "/")
	// }

	songBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return songBytes, nil
}
