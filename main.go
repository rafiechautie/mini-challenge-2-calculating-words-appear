package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan kalimat: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	charCount := make(map[rune]int)

	for _, char := range input {
		if char != '\n' {
			charCount[char]++
		}
	}

	fmt.Print(charCount)

	for _, char := range input {
		if char == ' ' {
			fmt.Println()
			continue 
		}
		fmt.Printf("%c\n", char)
	}

	fmt.Print("map[")
	for char, count := range charCount {
		if char == ' ' {
			fmt.Printf(" :%d ", count)
		} else {
			fmt.Printf("%c:%d ", char, count)
		}
	}
	fmt.Println("]")

}