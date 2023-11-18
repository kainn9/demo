package cameraSharedUtil

import clientGlobals "github.com/kainn9/demo/globalConfig/client"

func ZoomSpacingX(zoom float64) float64 {

	return (1 - zoom) * float64(clientGlobals.SCREEN_WIDTH) / 2
}

func ZoomSpacingY(zoom float64) float64 {
	return (1 - zoom) * float64(clientGlobals.SCREEN_HEIGHT) / 2
}
