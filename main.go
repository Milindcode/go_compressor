package main

import (
	"bufio"
	// "fmt"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatalf("An Error occurred:  %v", err)
	}
}

func main() {
	frequency_map := make(map[rune]int)

	file, err := os.Open("data.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()
		for _, char := range s {
			frequency_map[char]++
		}
		frequency_map['\n']++
	}

}
