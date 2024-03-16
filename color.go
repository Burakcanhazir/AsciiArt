package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	//RENK TANIMA
	var colors = map[string]string{
		"reset":           "\033[0m",  // Reset color
		"bold":            "\033[1m",  // Bold
		"dim":             "\033[2m",  // Dim
		"italic":          "\033[3m",  // Italic
		"underline":       "\033[4m",  // Underline
		"blink":           "\033[5m",  // Blink
		"reverse":         "\033[7m",  // Reverse (swap foreground and background colors)
		"hidden":          "\033[8m",  // Hidden (invisible)
		"strikethrough":   "\033[9m",  // Strikethrough
		"frame":           "\033[51m", // Framed
		"encircle":        "\033[52m", // Encircled
		"overline":        "\033[53m", // Overlined
		"background":      "\033[7m",  // Background color
		"black":           "\033[30m",
		"red":             "\033[31m",
		"green":           "\033[32m",
		"yellow":          "\033[33m",
		"blue":            "\033[34m",
		"magenta":         "\033[35m",
		"cyan":            "\033[36m",
		"white":           "\033[37m",
		"gray":            "\033[90m",
		"brightred":       "\033[91m",
		"brightgreen":     "\033[92m",
		"brightyellow":    "\033[93m",
		"brightblue":      "\033[94m",
		"brightmagenta":   "\033[95m",
		"brightcyan":      "\033[96m",
		"brightwhite":     "\033[97m",
		"bgblack":         "\033[40m",       // Background black
		"bgred":           "\033[41m",       // Background red
		"bggreen":         "\033[42m",       // Background green
		"bgyellow":        "\033[43m",       // Background yellow
		"bgblue":          "\033[44m",       // Background blue
		"bgmagenta":       "\033[45m",       // Background magenta
		"bgcyan":          "\033[46m",       // Background cyan
		"bgwhite":         "\033[47m",       // Background white
		"bggray":          "\033[100m",      // Background gray
		"bgbrightred":     "\033[101m",      // Background bright red
		"bgbrightgreen":   "\033[102m",      // Background bright green
		"bgbrightyellow":  "\033[103m",      // Background bright yellow
		"bgbrightblue":    "\033[104m",      // Background bright blue
		"bgbrightmagenta": "\033[105m",      // Background bright magenta
		"bgbrightcyan":    "\033[106m",      // Background bright cyan
		"bgbrightwhite":   "\033[107m",      // Background bright white
		"blackbg":         "\033[40m",       // Background black
		"redbg":           "\033[41m",       // Background red
		"greenbg":         "\033[42m",       // Background green
		"yellowbg":        "\033[43m",       // Background yellow
		"bluebg":          "\033[44m",       // Background blue
		"magentabg":       "\033[45m",       // Background magenta
		"cyanbg":          "\033[46m",       // Background cyan
		"whitebg":         "\033[47m",       // Background white
		"orange":          "\033[38;5;208m", // orange
	}
	if len(os.Args) < 3 {
		fmt.Println("ERROR:= arguman count is missing")
		return
	}
	// variables
	flag := os.Args[1][:8]
	color := strings.ToLower(os.Args[1][8:])
	letters := ""
	word := os.Args[len(os.Args)-1]
	newword := strings.Split(word, "\\n")

	if flag != "--color=" {
		fmt.Println("ERROR:= The color flag was not used")
		return
	}

	if len(os.Args) == 4 {
		letters += os.Args[2]
	} else {
		letters += os.Args[len(os.Args)-1]
	}

	// FİLE FUNC
	dosya, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error:= An error occurred while opening the file")
		panic(err)
	}
	defer dosya.Close() // dosya kapandı

	var lines []string

	scanner := bufio.NewScanner(dosya)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:= reading file")
		panic(err)
	}
	printColoredAsciiArt(newword, lines, colors[color], letters)

}
func printColoredAsciiArt(newword []string, lines []string, color string, letters string) {
	for i, word := range newword {
		if word == "" {
			if i != 0 {
				fmt.Println()
			}
			continue
		}
		for h := 1; h < 9; h++ {
			for i := 0; i < len(word); i++ {
				paint := false
				for _, char := range letters {
					if string(char) == string(word[i]) {
						paint = true
						break
					}
				}
				for lineIndex, line := range lines {
					if lineIndex == (int(word[i])-32)*9+h {
						if paint {
							fmt.Print(color + line + "\033[0m")
							break
						} else {
							fmt.Print(line)
							break
						}
					}
				}
			}
			fmt.Println()
		}
	}
}
