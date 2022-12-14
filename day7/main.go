package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	name           string
	subDirectories []Directory
	parent         *Directory
	files          []File
}

type File struct {
	name string
	size int
}

func main() {
	fmt.Println("Day 7")

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

	var root *Directory
	root.name = "/"

	var currentDir Directory

	for scanner.Scan() {
		input := scanner.Text()
		if string(input[0]) == "$" {
			inputArr := strings.Split(input, " ")
			switch inputArr[1] {
			case "cd":
				currentDir = currentDir.changeDirectory(inputArr[2])
			case "ls": //
				break
			}
		} else {
			inputArr := strings.Split(input, " ")
			if intVal, err := strconv.Atoi(inputArr[0]); err == nil { // first string is a number, so it's a file
				var file File
				file.name = inputArr[1]
				file.size = intVal

				currentDir.files = append(currentDir.files, file)
			} else { // it's a directory so add to subdirectories
				var dir Directory
				dir.name = inputArr[1]
				dir.parent = ptrDir

				currentDir.subDirectories = append(currentDir.subDirectories, dir)
			}
		}
	}
	fmt.Println(currentDir)
}

func (d *Directory) changeDirectory(name string) *Directory {
	var dir *Directory

	if d.name == "" {
		dir.name = name
		return dir
	}

	switch name {
	case "..":
		dir = d.parent
	default:
		for _, directory := range d.subDirectories {
			if directory.name == name {
				dir = &directory
			}
		}
	}
	return dir
}

func (d *Directory) equals(dir Directory) bool {
	return d.name == dir.name
}
