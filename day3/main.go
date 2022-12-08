package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Day 3")

	part1()
	part2()
}

func part1() {
	file, err := os.Open("day3.txt")
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
	sum := 0
	for scanner.Scan() {
		val := scanner.Text()
		strLen := len(val)
		c1 := val[0 : strLen/2]
		c2 := val[strLen/2 : strLen]
		priority := 0
		for i := 0; i < len(c1); i++ {
			if strings.Contains(c2, string(c1[i])) {
				priority = int(c1[i])
			}
		}
		sum += convertPriority(priority)
	}
	fmt.Printf("Sum part 1: %v\n", sum)
}

func part2() {
	file, err := os.Open("day3.txt")
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
	sum := 0
	var groupSlice []string
	for scanner.Scan() {
		groupSlice = append(groupSlice, scanner.Text())

		if len(groupSlice) == 3 {
			for _, char := range groupSlice[0] {
				if strings.Contains(groupSlice[1], string(char)) && strings.Contains(groupSlice[2], string(char)) {
					sum += convertPriority(int(char))
					groupSlice = nil
					break
				}
			}
		}
	}
	fmt.Printf("Sum part 2: %v\n", sum)
}

func convertPriority(priority int) int {
	result := 0
	if string(priority) != "" {
		if unicode.IsUpper(rune(priority)) {
			result = priority - 38
		} else if unicode.IsLower(rune(priority)) {
			result = priority - 96
		}
	}
	return result
}
