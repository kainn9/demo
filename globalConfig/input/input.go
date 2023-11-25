package inputGlobals

import "github.com/hajimehoshi/ebiten/v2"

// Special Inputs.
const (
	NO_INPUT               = -1
	COMBO_DOWN_SPACE       = -2
	RELEASED_HORIZONTAL    = -3
	RELEASED_VERTICAL_UP   = -4
	RELEASED_VERTICAL_DOWN = -5
)

// Key Binds.
type KeyBind = string

const (
	KEY_BIND_LEFT           KeyBind = "left"
	KEY_BIND_RIGHT          KeyBind = "right"
	KEY_BIND_JUMP           KeyBind = "jump"
	KEY_BIND_DODGE          KeyBind = "dodge"
	KEY_BIND_DOWN           KeyBind = "down"
	KEY_BIND_UP             KeyBind = "up"
	KEY_BIND_INTERACT       KeyBind = "interact"
	KEY_BIND_PRIMARY_ATTACK KeyBind = "atkPrimary"
)

var KEY_BINDS = map[KeyBind]ebiten.Key{
	KEY_BIND_LEFT:           ebiten.KeyA,
	KEY_BIND_RIGHT:          ebiten.KeyD,
	KEY_BIND_JUMP:           ebiten.KeySpace,
	KEY_BIND_UP:             ebiten.KeyW,
	KEY_BIND_DOWN:           ebiten.KeyS,
	KEY_BIND_INTERACT:       ebiten.KeyEnter,
	KEY_BIND_PRIMARY_ATTACK: ebiten.KeyK,
	KEY_BIND_DODGE:          ebiten.KeyShift,
}

func KeyInteract() ebiten.Key {
	return KEY_BINDS[KEY_BIND_INTERACT]
}

func KeyLeft() ebiten.Key {
	return KEY_BINDS[KEY_BIND_LEFT]
}

func KeyRight() ebiten.Key {
	return KEY_BINDS[KEY_BIND_RIGHT]
}

func KeyJump() ebiten.Key {
	return KEY_BINDS[KEY_BIND_JUMP]
}

func KeyUp() ebiten.Key {
	return KEY_BINDS[KEY_BIND_UP]
}

func KeyDown() ebiten.Key {
	return KEY_BINDS[KEY_BIND_DOWN]
}

func KeyDodge() ebiten.Key {
	return KEY_BINDS[KEY_BIND_DODGE]
}

func KeyPrimaryAttack() ebiten.Key {
	return KEY_BINDS[KEY_BIND_PRIMARY_ATTACK]
}

func KeyReleasedHorizontal() ebiten.Key {
	return ebiten.Key(RELEASED_HORIZONTAL)
}

func KeyReleasedVerticalUp() ebiten.Key {
	return ebiten.Key(RELEASED_VERTICAL_UP)
}

func KeyReleasedVerticalDown() ebiten.Key {
	return ebiten.Key(RELEASED_VERTICAL_DOWN)
}

func KeyComboDownSpace() ebiten.Key {
	return ebiten.Key(COMBO_DOWN_SPACE)
}

func KeyNoInput() ebiten.Key {
	return ebiten.Key(NO_INPUT)
}
