package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 6")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()

	for i, char := range input {
		if i >= 13 {
			substr := input[i-13 : i+1]
			fmt.Printf("Char: %v, previous 4 chars: %v\n", string(char), substr)

			startOfPacket := true
			for _, c := range substr {
				if strings.Count(substr, string(c)) > 1 {
					startOfPacket = false
				}
			}

			if startOfPacket {
				fmt.Printf("Start of packet found at: %v\n", i+1)
				return
			}
		}
	}

}
