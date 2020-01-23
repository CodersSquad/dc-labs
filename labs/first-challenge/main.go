package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	letters := map[rune]bool{}
	counter := int(0)
	max := int(0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	for _, character := range text {
		if letters[character] {
			if counter > max {
				max = counter
			}
			counter = 0
		}
		letters[character] = true
		counter++
	}

	fmt.Printf("%v", max)
}
