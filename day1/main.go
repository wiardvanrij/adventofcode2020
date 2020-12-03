package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//day1part1()
	day1part2()
}

func day1part2() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i, line := range lines {
		for _, lineII := range lines {

			rest := 2020 - line - lineII
			if search(i, lines, rest) {
				fmt.Printf("data 1: %d and data 2: %d and data 3: %d \n", line, lineII, rest)
				fmt.Println(line * lineII * rest)
				return
			}
		}
	}
}

func day1part1() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i, line := range lines {

		rest := 2020 - line
		if search(i, lines, rest) {
			fmt.Printf("data 1: %d and data 2: %d \n", line, rest)
			fmt.Println(line * rest)
			return
		}
	}
}

func search(iSearch int, data []int, searchable int) bool {
	for i, line := range data {
		if i == iSearch {
			continue
		}

		if line == searchable {
			return true
		}
	}
	return false
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intLine, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("failed string conv")
			continue
		}
		lines = append(lines, intLine)
	}
	return lines, scanner.Err()
}
