package textUtil

import (
	"image"
	"math"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kainn9/coldBrew"
	"github.com/kainn9/demo/components"
	assetComponents "github.com/kainn9/demo/components/assets"
	fontConstants "github.com/kainn9/demo/constants/font"
	cameraUtil "github.com/kainn9/demo/systems/render/util/camera"
	systemsUtil "github.com/kainn9/demo/systems/util"
	"github.com/yohamta/donburi"
)

// Draws text with the given parameters.
// The text will be rendered with the given font sprites.
// The text will be rendered with the given adjustments.
// The text will be rendered with the given animation parameters
// (startTick, ticksPerWord) — these can be set to 0 to render the
// text immediately/with no animation.
func RenderText(
	text string,
	x, y, maxWidth, charWidth, charHeight, spaceWidth float64,
	startTick, ticksPerWord int,
	camera *components.Camera,
	lower, upper, numbers, special *assetComponents.Sprite,
	adjustmentMap map[string]fontConstants.FontAdjustment,
	manager *coldBrew.Manager,

) {

	if int(maxWidth)%fontConstants.FONT_DEFAULT_WIDTH != 0 {
		panic("maxWidth must be a multiple of charWidth")
	}

	startX := x
	currX := x
	currY := y

	tickHandler := manager.TickHandler

	ticksSinceStart := tickHandler.TicksSinceNTicks(startTick)

	words := strings.Split(text, " ")

	for i, word := range words {

		var animCoefficient float64 = 1
		if ticksSinceStart < ticksPerWord*int(i) {
			targetTickValue := ticksPerWord * i
			coefficient := float64(ticksSinceStart) / float64(targetTickValue)
			coefficient = math.Min(1, math.Max(0, coefficient))
			animCoefficient = coefficient
		}

		word = StripInvalidCharacters(word)
		PrintWord(word, currX, currY, charWidth, charHeight, animCoefficient, lower, upper, numbers, special, adjustmentMap, camera)
		currX, currY = NextWordPosition(currX, currY, charWidth, charHeight, maxWidth, startX, spaceWidth, word, adjustmentMap)

	}
}

func StripInvalidCharacters(word string) string {
	return strings.Map(func(r rune) rune {
		if _, ok := fontConstants.AllCharacterMap[string(r)]; ok {
			return r
		}

		if num, err := strconv.Atoi(string(r)); err == nil && num <= 9 {
			return r
		}
		return -1

	}, word)
}

// Draws a word at the given position.
// Anim coefficient is used for initial render effects
// but can be set to 1 to render the word normally.
func PrintWord(
	word string,
	currentX, currentY, charWidth, charHeight, animCoefficient float64,
	lower, upper, numbers, special *assetComponents.Sprite,
	adjustmentMap map[string]fontConstants.FontAdjustment,
	camera *components.Camera,
) {

	var charSheet *ebiten.Image
	var charSheetIndex int
	prevChar := ""
	adjustmentX := 0.0

	for i, charRune := range word {

		if animCoefficient < 0.7 {
			continue
		}

		char := string(charRune)

		if _, ok := fontConstants.LowerCaseCharacterMap[char]; ok {
			charSheet = lower.Image
			charSheetIndex = fontConstants.LowerCaseCharacterMap[char]
		}

		if _, ok := fontConstants.UpperCaseCharacterMap[char]; ok {
			charSheet = upper.Image
			charSheetIndex = fontConstants.UpperCaseCharacterMap[char]
		}

		if _, ok := fontConstants.SpecialCharacterMap[char]; ok {
			charSheet = special.Image
			charSheetIndex = fontConstants.SpecialCharacterMap[char]
		}

		if num, err := strconv.Atoi(char); err == nil && num <= 9 {
			charSheet = numbers.Image
			charSheetIndex = num + 1
		}

		charWidthInt := int(charWidth)
		sx, sy := (charSheetIndex-1)*charWidthInt, (0)

		rect := image.Rect(sx, sy, sx+charWidthInt, int(charHeight))

		charImage := charSheet.SubImage(rect).(*ebiten.Image)

		adjustmentX += adjustmentMap[prevChar].X
		xPos := currentX + float64(i)*charWidth + adjustmentX
		yPos := currentY + adjustmentMap[char].Y - ((charHeight / 2) * (1 - animCoefficient))

		charOpts := &ebiten.DrawImageOptions{}
		charOpts.GeoM.Translate(xPos, yPos)

		charOpts.ColorScale.ScaleAlpha(float32(animCoefficient))

		cameraUtil.AddImage(camera, charImage, charOpts)

		prevChar = char
	}

}

// Returns the next position of a word, uses maxWidth to determine if a new line is needed.
func NextWordPosition(
	currX, currY, charWidth, charHeight, maxWidth, startX, spaceWidth float64,
	word string,
	adjustmentMap map[string]fontConstants.FontAdjustment,
) (nextX, nextY float64) {

	nextX = currX
	nextY = currY

	for _, charRune := range word {

		char := string(charRune)
		adjustment := adjustmentMap[char]
		nextX += charWidth + adjustment.X

	}

	// New line.
	if nextX >= maxWidth {
		factor := math.Floor(nextX / maxWidth)
		nextY = nextY + (charHeight * factor)
		nextX = startX
	} else {
		// Space.
		nextX += spaceWidth
	}

	return nextX, nextY
}

// Wrapper around RenderText that uses the default font.
func RenderTextDefault(
	text string,
	x, y, maxWidth float64,
	startTick, ticksPerWord int,
	camera *components.Camera,
	world *donburi.World,
	manager *coldBrew.Manager,
) {

	lower, upper, numbers, special := systemsUtil.GetDefaultFontSpriteMap(*world)

	RenderText(
		text,
		x, y, maxWidth, fontConstants.FONT_DEFAULT_WIDTH, fontConstants.FONT_DEFAULT_HEIGHT, fontConstants.FONT_DEFAULT_WIDTH-1,
		startTick, ticksPerWord,
		camera,
		lower, upper, numbers, special,
		fontConstants.FONT_DEFAULT_ADJUSTMENT_MAP,
		manager,
	)
}