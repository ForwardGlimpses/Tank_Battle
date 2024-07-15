package ebitenextend

import (
	"strings"
	"github.com/hajimehoshi/ebiten/v2"
)

type Key int 


func KeyNameToKeyCode(name string) (Key, bool) {
	switch strings.ToLower(name) {
	case "0":
		return Key(ebiten.Key0), true
	case "1":
		return Key(ebiten.Key1), true
	case "2":
		return Key(ebiten.Key2), true
	case "3":
		return Key(ebiten.Key3), true
	case "4":
		return Key(ebiten.Key4), true
	case "5":
		return Key(ebiten.Key5), true
	case "6":
		return Key(ebiten.Key6), true
	case "7":
		return Key(ebiten.Key7), true
	case "8":
		return Key(ebiten.Key8), true
	case "9":
		return Key(ebiten.Key9), true
	case "a":
		return Key(ebiten.KeyA), true
	case "b":
		return Key(ebiten.KeyB), true
	case "c":
		return Key(ebiten.KeyC), true
	case "d":
		return Key(ebiten.KeyD), true
	case "e":
		return Key(ebiten.KeyE), true
	case "f":
		return Key(ebiten.KeyF), true
	case "g":
		return Key(ebiten.KeyG), true
	case "h":
		return Key(ebiten.KeyH), true
	case "i":
		return Key(ebiten.KeyI), true
	case "j":
		return Key(ebiten.KeyJ), true
	case "k":
		return Key(ebiten.KeyK), true
	case "l":
		return Key(ebiten.KeyL), true
	case "m":
		return Key(ebiten.KeyM), true
	case "n":
		return Key(ebiten.KeyN), true
	case "o":
		return Key(ebiten.KeyO), true
	case "p":
		return Key(ebiten.KeyP), true
	case "q":
		return Key(ebiten.KeyQ), true
	case "r":
		return Key(ebiten.KeyR), true
	case "s":
		return Key(ebiten.KeyS), true
	case "t":
		return Key(ebiten.KeyT), true
	case "u":
		return Key(ebiten.KeyU), true
	case "v":
		return Key(ebiten.KeyV), true
	case "w":
		return Key(ebiten.KeyW), true
	case "x":
		return Key(ebiten.KeyX), true
	case "y":
		return Key(ebiten.KeyY), true
	case "z":
		return Key(ebiten.KeyZ), true
	case "alt":
		return Key(ebiten.KeyAlt), true
	case "altleft":
		return Key(ebiten.KeyAltLeft), true
	case "altright":
		return Key(ebiten.KeyAltRight), true
	case "apostrophe":
		return Key(ebiten.KeyApostrophe), true
	case "arrowdown":
		return Key(ebiten.KeyArrowDown), true
	case "arrowleft":
		return Key(ebiten.KeyArrowLeft), true
	case "arrowright":
		return Key(ebiten.KeyArrowRight), true
	case "arrowup":
		return Key(ebiten.KeyArrowUp), true
	case "backquote":
		return Key(ebiten.KeyBackquote), true
	case "backslash":
		return Key(ebiten.KeyBackslash), true
	case "backspace":
		return Key(ebiten.KeyBackspace), true
	case "bracketleft":
		return Key(ebiten.KeyBracketLeft), true
	case "bracketright":
		return Key(ebiten.KeyBracketRight), true
	case "capslock":
		return Key(ebiten.KeyCapsLock), true
	case "comma":
		return Key(ebiten.KeyComma), true
	case "contextmenu":
		return Key(ebiten.KeyContextMenu), true
	case "control":
		return Key(ebiten.KeyControl), true
	case "controlleft":
		return Key(ebiten.KeyControlLeft), true
	case "controlright":
		return Key(ebiten.KeyControlRight), true
	case "delete":
		return Key(ebiten.KeyDelete), true
	case "digit0":
		return Key(ebiten.KeyDigit0), true
	case "digit1":
		return Key(ebiten.KeyDigit1), true
	case "digit2":
		return Key(ebiten.KeyDigit2), true
	case "digit3":
		return Key(ebiten.KeyDigit3), true
	case "digit4":
		return Key(ebiten.KeyDigit4), true
	case "digit5":
		return Key(ebiten.KeyDigit5), true
	case "digit6":
		return Key(ebiten.KeyDigit6), true
	case "digit7":
		return Key(ebiten.KeyDigit7), true
	case "digit8":
		return Key(ebiten.KeyDigit8), true
	case "digit9":
		return Key(ebiten.KeyDigit9), true
	case "down":
		return Key(ebiten.KeyDown), true
	case "end":
		return Key(ebiten.KeyEnd), true
	case "enter":
		return Key(ebiten.KeyEnter), true
	case "equal":
		return Key(ebiten.KeyEqual), true
	case "escape":
		return Key(ebiten.KeyEscape), true
	case "f1":
		return Key(ebiten.KeyF1), true
	case "f2":
		return Key(ebiten.KeyF2), true
	case "f3":
		return Key(ebiten.KeyF3), true
	case "f4":
		return Key(ebiten.KeyF4), true
	case "f5":
		return Key(ebiten.KeyF5), true
	case "f6":
		return Key(ebiten.KeyF6), true
	case "f7":
		return Key(ebiten.KeyF7), true
	case "f8":
		return Key(ebiten.KeyF8), true
	case "f9":
		return Key(ebiten.KeyF9), true
	case "f10":
		return Key(ebiten.KeyF10), true
	case "f11":
		return Key(ebiten.KeyF11), true
	case "f12":
		return Key(ebiten.KeyF12), true
	case "graveaccent":
		return Key(ebiten.KeyGraveAccent), true
	case "home":
		return Key(ebiten.KeyHome), true
	case "insert":
		return Key(ebiten.KeyInsert), true
	case "kp0":
		return Key(ebiten.KeyKP0), true
	case "kp1":
		return Key(ebiten.KeyKP1), true
	case "kp2":
		return Key(ebiten.KeyKP2), true
	case "kp3":
		return Key(ebiten.KeyKP3), true
	case "kp4":
		return Key(ebiten.KeyKP4), true
	case "kp5":
		return Key(ebiten.KeyKP5), true
	case "kp6":
		return Key(ebiten.KeyKP6), true
	case "kp7":
		return Key(ebiten.KeyKP7), true
	case "kp8":
		return Key(ebiten.KeyKP8), true
	case "kp9":
		return Key(ebiten.KeyKP9), true
	case "kpadd":
		return Key(ebiten.KeyKPAdd), true
	case "kpdecimal":
		return Key(ebiten.KeyKPDecimal), true
	case "kpdivide":
		return Key(ebiten.KeyKPDivide), true
	case "kpenter":
		return Key(ebiten.KeyKPEnter), true
	case "kpequal":
		return Key(ebiten.KeyKPEqual), true
	case "kpmultiply":
		return Key(ebiten.KeyKPMultiply), true
	case "kpsubtract":
		return Key(ebiten.KeyKPSubtract), true
	case "left":
		return Key(ebiten.KeyLeft), true
	case "leftbracket":
		return Key(ebiten.KeyLeftBracket), true
	case "menu":
		return Key(ebiten.KeyMenu), true
	case "meta":
		return Key(ebiten.KeyMeta), true
	case "metaleft":
		return Key(ebiten.KeyMetaLeft), true
	case "metaright":
		return Key(ebiten.KeyMetaRight), true
	case "minus":
		return Key(ebiten.KeyMinus), true
	case "numlock":
		return Key(ebiten.KeyNumLock), true
	case "numpad0":
		return Key(ebiten.KeyNumpad0), true
	case "numpad1":
		return Key(ebiten.KeyNumpad1), true
	case "numpad2":
		return Key(ebiten.KeyNumpad2), true
	case "numpad3":
		return Key(ebiten.KeyNumpad3), true
	case "numpad4":
		return Key(ebiten.KeyNumpad4), true
	case "numpad5":
		return Key(ebiten.KeyNumpad5), true
	case "numpad6":
		return Key(ebiten.KeyNumpad6), true
	case "numpad7":
		return Key(ebiten.KeyNumpad7), true
	case "numpad8":
		return Key(ebiten.KeyNumpad8), true
	case "numpad9":
		return Key(ebiten.KeyNumpad9), true
	case "numpadadd":
		return Key(ebiten.KeyNumpadAdd), true
	case "numpaddecimal":
		return Key(ebiten.KeyNumpadDecimal), true
	case "numpaddivide":
		return Key(ebiten.KeyNumpadDivide), true
	case "numpadenter":
		return Key(ebiten.KeyNumpadEnter), true
	case "numpadequal":
		return Key(ebiten.KeyNumpadEqual), true
	case "numpadmultiply":
		return Key(ebiten.KeyNumpadMultiply), true
	case "numpadsubtract":
		return Key(ebiten.KeyNumpadSubtract), true
	case "pagedown":
		return Key(ebiten.KeyPageDown), true
	case "pageup":
		return Key(ebiten.KeyPageUp), true
	case "pause":
		return Key(ebiten.KeyPause), true
	case "period":
		return Key(ebiten.KeyPeriod), true
	case "printscreen":
		return Key(ebiten.KeyPrintScreen), true
	case "quote":
		return Key(ebiten.KeyQuote), true
	case "right":
		return Key(ebiten.KeyRight), true
	case "rightbracket":
		return Key(ebiten.KeyRightBracket), true
	case "scrolllock":
		return Key(ebiten.KeyScrollLock), true
	case "semicolon":
		return Key(ebiten.KeySemicolon), true
	case "shift":
		return Key(ebiten.KeyShift), true
	case "shiftleft":
		return Key(ebiten.KeyShiftLeft), true
	case "shiftright":
		return Key(ebiten.KeyShiftRight), true
	case "slash":
		return Key(ebiten.KeySlash), true
	case "space":
		return Key(ebiten.KeySpace), true
	case "tab":
		return Key(ebiten.KeyTab), true
	case "up":
		return Key(ebiten.KeyUp), true
	}
	return 0, false
}
