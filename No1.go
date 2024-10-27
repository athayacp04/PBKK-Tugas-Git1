package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&input)

	if input == 42 {
		fmt.Println("Hello Universe")
	} else {
		fmt.Println(input)
	}
}
