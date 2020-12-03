package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part2() {
	data, _ := readLines("input.txt")

	valid := 0
	for _, line := range data {

		split := strings.Split(line, " ")

		minMax := splitDash(split[0])
		letter := removeChar(split[1], ":")

		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		validTmp := false
		for i, singleChar := range split[2] {
			if (i + 1) == min {
				if string(singleChar) == letter {
					validTmp = true
				}
			}
			if (i + 1) == max {
				if string(singleChar) == letter {
					if validTmp {
						validTmp = false
					} else {
						validTmp = true
					}
				}
			}
		}
		if validTmp {
			valid++
		}

	}

	fmt.Println(valid)
}

func part1() {
	data, _ := readLines("input.txt")

	valid := 0
	for _, line := range data {

		split := strings.Split(line, " ")

		minMax := splitDash(split[0])
		letter := removeChar(split[1], ":")

		letterCount := strings.Count(split[2], letter)

		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])

		if letterCount >= min && letterCount <= max {
			valid++
		}

	}

	fmt.Println(valid)
}

func splitDash(data string) []string {
	split := strings.Split(data, "-")
	return split
}

func removeChar(data string, char string) string {
	t := strings.Replace(data, char, "", -1)
	return t
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
