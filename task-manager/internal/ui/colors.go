package ui

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)

// Включаем поддержку ANSI-цветов в стандартной консоли Windows
func init() {
	handle := syscall.Handle(os.Stdout.Fd())
	var mode uint32
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getConsoleMode := kernel32.NewProc("GetConsoleMode")
	setConsoleMode := kernel32.NewProc("SetConsoleMode")

	getConsoleMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&mode)))
	mode |= 0x0004 // ENABLE_VIRTUAL_TERMINAL_PROCESSING
	setConsoleMode.Call(uintptr(handle), uintptr(mode))
}

func PrintSuccess(msg string) {
	fmt.Printf("%s✅ %s%s\n", Green, msg, Reset)
}

func PrintError(err error) {
	fmt.Printf("%s❌ Ошибка: %v%s\n", Red, err, Reset)
}

func PrintInfo(msg string) {
	fmt.Printf("%sℹ️  %s%s\n", Cyan, msg, Reset)
}

func PrintWarning(msg string) {
	fmt.Printf("%s⚠️  %s%s\n", Yellow, msg, Reset)
}
