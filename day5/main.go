package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 5")

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

	//var stacks [9]string
	for scanner.Scan() {
		row := scanner.Text()

		if !strings.Contains(scanner.Text(), "1") {
			fmt.Printf("%v,%v,%v,%v,%v,%v,%v,%v,%v\n", string(row[1]), string(row[5]), string(row[9]), string(row[13]), string(row[17]), string(row[21]), string(row[25]), string(row[29]), string(row[33]))
			fmt.Printf("row size: %v\n", len(row))
			//rowArr := strings.Split(row, " ")
			//fmt.Println(rowArr)
		}
	}
}
