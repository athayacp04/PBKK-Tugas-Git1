package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Split(input, " ")
	var reversedWords []string

	for i := len(words) - 1; i >= 0; i-- {
		reversedWords = append(reversedWords, words[i])
	}

	return strings.Join(reversedWords, " / ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama minimal 3 kata: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Cek jumlah kata
	if len(strings.Split(input, " ")) < 3 {
		fmt.Println("Masukkan harus minimal 3 kata.")
		return
	}

	result := reverseWords(input)
	fmt.Println("Nama dibalik:", result)
}
