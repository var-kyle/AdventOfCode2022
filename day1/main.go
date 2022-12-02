package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Day 1")

	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	var values []int

	for scanner.Scan() {
		val := scanner.Text()
		if val != "" {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			sum += intVal
		} else {
			values = append(values, sum)
			sum = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	fmt.Printf("First: %v\nSecond: %v\nThird: %v\n", values[0], values[1], values[2])
	fmt.Printf("Total from the top 3: %v\n", values[0]+values[1]+values[2])
}
