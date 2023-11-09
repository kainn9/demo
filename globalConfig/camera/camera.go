package cameraGlobals

import (
	clientGlobals "github.com/kainn9/demo/globalConfig/client"
)

var (
	CAMERA_MIN_SPEED     = 2.5
	CAMERA_MAX_SPEED     = 3.5
	CAMERA_EPSILON   int = clientGlobals.SCREEN_WIDTH/2 - 270
)
