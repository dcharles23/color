package main

import (
	"fmt"
	"os"
	"strings"
)

// ColorCodes maps color names to ANSI escape codes.
var ColorCodes = map[string]string{
	"Reset":        "\033[0m",
	"Black":        "\033[30m",
	"Red":          "\033[31m",
	"Green":        "\033[32m",
	"Yellow":       "\033[33m",
	"Blue":         "\033[34m",
	"Magenta":      "\033[35m",
	"Cyan":         "\033[36m",
	"LightGray":    "\033[37m",
	"Gray":         "\033[90m",
	"LightRed":     "\033[91m",
	"LightGreen":   "\033[92m",
	"LightYellow":  "\033[93m",
	"LightBlue":    "\033[94m",
	"LightMagenta": "\033[95m",
	"LightCyan":    "\033[96m",
	"White":        "\033[97m",
	"Orange":       "\033[38;5;208m",
}

func main() {
	// Check if there are more than 2 command-line arguments.
	if len(os.Args) > 2 {
		// Retrieve the color argument from the command line.
		ColorFromOutside := os.Args[1]
		textColor := ""
		// Check if the color argument is longer than 8 characters.
		if len(ColorFromOutside) > 8 {
			// Extract the text color from the color argument (characters after the first 8).
			textColor = ColorFromOutside[8:]
		}

		// Check if there are exactly 3 command-line arguments.
		if len(os.Args) == 3 {
			// Retrieve the text argument from the command line.
			textFromOutside := os.Args[2]
			// Print the text in the specified color using the printColored function.
			printColored(textFromOutside, textColor)
			// Print a newline to separate the output.
			fmt.Println("")

		} else if len(os.Args) == 4 {
			// Retrieve the text and substring arguments from the command line.
			textFromOutside := os.Args[3]
			subStringFromOutside := os.Args[2]

			// Count the number of times the substring appears in the text.
			letterCount := strings.Count(textFromOutside, subStringFromOutside)

			// Loop through the text to find and color each occurrence of the substring.
			for i := 0; i < letterCount; i++ {
				// Determine the length and index of the substring within the text.
				lenOfSubStr := len(subStringFromOutside)
				idxOfSubStrInStr := strings.Index(textFromOutside, subStringFromOutside)

				// Print the portion of the text before the substring in its original color.
				fmt.Print(textFromOutside[:idxOfSubStrInStr])
				// Print the substring in the specified color using the printColored function.
				printColored(subStringFromOutside, textColor)
				// Update the text to exclude the processed portion of the substring.
				textFromOutside = textFromOutside[idxOfSubStrInStr+lenOfSubStr:]
			}
			// Print the remaining part of the text in its original color.
			fmt.Println(textFromOutside)
		}
	}
}

// printColored takes text and a color name, and prints the text in the specified color.
func printColored(text, textColor string) {
	titleCase := strings.Title(textColor)
	fmt.Print(ColorCodes[titleCase] + text + ColorCodes["Reset"])
}
