package output

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Standard success message
func Standard(text string) {
	msg := changeColor(text, color.FgGreen)
	fmt.Println(msg)
}

// Error ends the running process with a red error message
func Error(text string) {
	msg := changeColor(text, color.FgRed)
	fmt.Println(msg)
	os.Exit(1)
}

func changeColor(text string, code color.Attribute) string {
	c := color.New(code).SprintFunc()
	return c(text)
}

// TableBody output
func TableBody(text string) string {
	return changeColor(text, color.FgWhite)
}

// TableHeader output
func TableHeader(text string) string {
	return changeColor(text, color.FgCyan)
}
