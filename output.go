package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	word := os.Args[2]
	banner := os.Args[len(os.Args)-1]
	flag := os.Args[1][:9]
	filename := os.Args[1][9:] // input dosya ismi olacak

	if flag != "--output=" {
		fmt.Println("FLAG HATASI")
		return
	}
	// küçük çaplı banner boş olursa ne olsun kontrolü
	if banner == "standard" {
		banner = "standard"
	} else if banner == "shadow" {
		banner = "shadow"
	} else if banner == "thinkertoy" {
		banner = "thinkertoy"
	} else {
		banner = "standard"
	}

	dosya, err := os.Open(banner + ".txt")
	if err != nil {
		fmt.Println("açma hatası")
	}
	defer dosya.Close()

	// scan the file and assigment be olum
	var file []string

	scanner := bufio.NewScanner(dosya)

	for scanner.Scan() {
		line := scanner.Text()
		file = append(file, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("dosya okunurken hata oluştu:", err)
	}
	// \n fixed babba
	str := strings.Split(word, "\\n")
	var res []string
	var result []string
	for _, r := range str {
		if r != "" {
			res = art(r, file)
			result = append(result, res...) // res elemanlarını tek tek geçiriyor. Döngüyle dolaşmak yerine yapıldı.
		} else {
			fmt.Println()
		}
	}
	writefile(filename, result)
}

// asciart
func art(word string, file []string) []string {
	kutu := make([]string, 8)

	for i := 0; i < 8; i++ {
		for _, r := range word {
			start := (int(r)-32)*9 + i
			for ind, val := range file {
				if ind == start {
					kutu[i] += string(val)
				}
			}
		}
	}
	return kutu
}

//The process of writing to a file kral

func writefile(dosyadi string, result []string) {
	dosya, err := os.Create(dosyadi) // burası inputa göre değişecek
	if err != nil {
		fmt.Println("Dosya oluşturulurken hata oluştu:", err)
		return
	}
	defer dosya.Close()

	// Browsing through the generated ASCII file.
	for _, kelime := range result {
		_, err = dosya.WriteString(string(kelime) + "\n")
		if err != nil {
			fmt.Println("yazılırken hata oldu", err)
			return
		}

	}
	fmt.Println("BAŞARIYLA YAZILDI")
}
