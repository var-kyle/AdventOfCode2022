package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Day 2")

	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println("Part 1")
	sumP1 := 0
	sumP2 := 0

	for scanner.Scan() {
		input := scanner.Text()

		opponent := string(input[0])
		choice := string(input[2])

		sumP1 += part1(choice, opponent)
		sumP2 += part2(choice, opponent)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total score for Part 1 is: %v\n", sumP1)
	fmt.Printf("Total score for Part 2 is: %v\n", sumP2)
}

func part1(choice string, opponent string) int {
	sum := 0
	if choice == "X" {
		sum += 1

		if opponent == "A" {
			sum += 3
		} else if opponent == "C" {
			sum += 6
		}
	} else if choice == "Y" {
		sum += 2

		if opponent == "A" {
			sum += 6
		} else if opponent == "B" {
			sum += 3
		}
	} else if choice == "Z" {
		sum += 3

		if opponent == "B" {
			sum += 6
		} else if opponent == "C" {
			sum += 3
		}
	}
	return sum
}

func part2(choice string, opponent string) int {
	sum := 0
	if choice == "X" {
		if opponent == "A" {
			sum += 3
		} else if opponent == "B" {
			sum += 1
		} else if opponent == "C" {
			sum += 2
		}
	} else if choice == "Y" {
		if opponent == "A" {
			sum += 4
		} else if opponent == "B" {
			sum += 5
		} else if opponent == "C" {
			sum += 6
		}
	} else if choice == "Z" {
		if opponent == "A" {
			sum += 8
		} else if opponent == "B" {
			sum += 9
		} else if opponent == "C" {
			sum += 7
		}
	}
	return sum
}
