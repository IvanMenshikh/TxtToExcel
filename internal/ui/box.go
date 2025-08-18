package ui

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type MessageBox struct{}

func NewMessageBox() *MessageBox {
	return &MessageBox{}
}

func (m *MessageBox) ShowInfo(title, msg string) {
	showMessageBox(title, msg, 0)
}

func (m *MessageBox) ShowError(title, msg string) {
	showMessageBox(title, msg, 100)
}

func showMessageBox(title, message string, style uint) {
	user32 := windows.NewLazySystemDLL("user32.dll")
	procMessageBoxW := user32.NewProc("MessageBoxW")

	_, _, _ = procMessageBoxW.Call(
		0,
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(message))),
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(title))),
		uintptr(style),
	)
}
