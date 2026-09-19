package ui

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)

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
