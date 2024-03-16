package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("hata: kelime argümanı eksik")
		return
	}

	text := os.Args[1]

	dosya, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("açılırken hata oluştu")
		panic(err)
	}
	defer dosya.Close()

	var lines []string
	scanner := bufio.NewScanner(dosya)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	words := strings.Split(text, "\\n")
	fmt.Println(words)

	for _, r := range words {
		if r == "" {
			fmt.Println()
		} else {
			writeTerminal(r, lines)

		}
	}

}

func writeTerminal(oha string, lines []string) {
	result := art(oha, lines)
	for _, line := range result {
		fmt.Println(line)
	}
}

func art(word string, lines []string) []string {
	kutu := make([]string, 8)

	for i := 0; i < 8; i++ {
		for _, r := range word {
			start := (int(r)-32)*9 + i + 1
			for ind, val := range lines {
				if ind == start {
					kutu[i] += val
				}
			}
		}
	}
	return kutu
}
