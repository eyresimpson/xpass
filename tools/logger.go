package tools

import (
	"fmt"
	"syscall"
	"unsafe"
)

// 日志模块

func Info(text string) {
	fmt.Printf("\033[0;34m%s\033[0m\n", "[INFO]:"+text)
}

func Warn(text string) {
	fmt.Printf("\033[0;33m%s\033[0m\n", "[WARN]:"+text)
}

func Success(text string) {
	fmt.Printf("\033[0;32m%s\033[0m\n", "[OK]:"+text)
}

func Err(text string, err error) {
	fmt.Printf("\033[1;31m%s\033[0m\n", "[ERR]:"+text)
	if err != nil {
		fmt.Printf("\033[1;31m%s\033[0m\n", "\t[ERR]:"+err.Error())
	}
}

func IntPtr(n int) uintptr {
	return uintptr(n)
}
func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

func ShowMessage(tittle, text string) {
	user32dll, _ := syscall.LoadLibrary("user32.dll")
	user32 := syscall.NewLazyDLL("user32.dll")
	MessageBoxW := user32.NewProc("MessageBoxW")
	MessageBoxW.Call(IntPtr(0), StrPtr(text), StrPtr(tittle), IntPtr(0))
	defer syscall.FreeLibrary(user32dll)
}
