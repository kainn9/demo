package fontConstants

// Font Constants.
const (
	FONT_LOWER_CASE_SPRITE_NAME = "lower"
	FONT_UPPER_CASE_SPRITE_NAME = "upper"
	FONT_NUMBERS_SPRITE_NAME    = "numbers"
	FONT_SPECIAL_SPRITE_NAME    = "special"
)

var lowerCaseSlice = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u",
	"v", "w", "x", "y", "z",
}

var upperCaseSlice = []string{
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
	"V", "W", "X", "Y", "Z",
}

var specialCharSlice = []string{
	"!", "?", ".", "%", ";", ":", "$", "#", "'", "\"",
	"/", "\\", "(", ")", "&", "*", "+", ",", "-",
	"<", ">", "=", "@", "[", "]", "^", "_", "`",
	"{", "}",
}

var LowerCaseCharacterMap = func() map[string]int {
	charMap := make(map[string]int)
	for index, value := range lowerCaseSlice {
		charMap[value] = index + 1 // Add one so its not 0 Based.
	}
	return charMap
}()

var UpperCaseCharacterMap = func() map[string]int {
	charMap := make(map[string]int)

	for index, value := range upperCaseSlice {
		charMap[value] = index + 1 // Add one so its not 0 Based.
	}
	return charMap
}()

var SpecialCharacterMap = func() map[string]int {
	charMap := make(map[string]int)

	for index, value := range specialCharSlice {
		charMap[value] = index + 1 // Add one so its not 0 Based.
	}
	return charMap
}()

var AllCharacterMap = func() map[string]int {
	charMap := make(map[string]int)

	for key, value := range LowerCaseCharacterMap {
		charMap[key] = value
	}

	for key, value := range UpperCaseCharacterMap {
		charMap[key] = value
	}

	for key, value := range SpecialCharacterMap {
		charMap[key] = value
	}

	return charMap
}()

type FontAdjustment struct {
	X float64
	Y float64
}

// Default Font constants.
const (
	FONT_DEFAULT_NAME   = "defaultFont"
	FONT_DEFAULT_WIDTH  = 7
	FONT_DEFAULT_HEIGHT = 15
)

var FONT_DEFAULT_ADJUSTMENT_MAP = map[string]FontAdjustment{
	"i": {X: -2, Y: 0},
	"I": {X: -2, Y: 0},
	"1": {X: -2, Y: 0},
	"'": {X: -2, Y: 0},
	"g": {X: 1, Y: 1},
	"l": {X: -3, Y: 0},
}
