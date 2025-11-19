package lib

import (
	"fmt"

	"github.com/fatih/color"
)

func LogInfo(message string) {
	color.Set(color.FgBlue, color.Bold)
	defer color.Unset()
	fmt.Println(message)
}

func LogChecklist(message string, status bool) {
	defer color.Unset()

	if status {
		color.Set(color.FgGreen, color.Bold)
		fmt.Println(message + " [✓]")
	} else {
		color.Set(color.FgRed, color.Bold)
		fmt.Println(message + " [✗]")
	}
}

func LogWarning(message string) {
	color.Set(color.FgYellow, color.Bold)
	defer color.Unset()
	fmt.Println(message)
}

func LogError(message string) {
	color.Set(color.FgRed, color.Bold)
	defer color.Unset()
	fmt.Println(message + " [⚠️]")
}
