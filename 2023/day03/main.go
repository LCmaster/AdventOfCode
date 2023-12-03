package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/LCmaster/AdventOfCode/utils"
)

func extractNumbers(input []string) map[int]map[int]string {
	numberMatcher := regexp.MustCompile("[0-9]+")

	numbers := make(map[int]map[int]string)
	for i, line := range input {
		for _, entry := range numberMatcher.FindAllStringIndex(line, -1) {
			if len(entry) > 0 {
				_, ok := numbers[i]
				if !ok {
					numbers[i] = make(map[int]string)
				}

				numbers[i][entry[0]] = line[entry[0]:entry[1]]
			}
		}
	}

	return numbers
}

func extractNumbersAroundSymbolAt(numbers *map[int]map[int]string, row int, column int) []int {
	vOffset, hOffset := -1, -1
	vCount, hCount := 3, 3
	if row == 0 {
		vOffset = 0
		vCount = 2
	}

	if column == 0 {
		hOffset = 0
		hCount = 2
	}

	numbersAround := []int{}
	for x := 0; x < hCount; x++ {
		for y := 0; y < vCount; y++ {
			checkRow := row + vOffset + y
			checkColumn := column + hOffset + x

			for start, numberStr := range (*numbers)[checkRow] {
				end := len(numberStr) + start - 1
				if checkColumn >= start && checkColumn <= end {
					number, err := strconv.Atoi(numberStr)
					if err == nil {
						numbersAround = append(numbersAround, number)
						delete((*numbers)[checkRow], start)
					}
				}
			}
		}
	}

	return numbersAround
}

func part1(input []string) int {
	output := 0

	numbers := extractNumbers(input)
	for i, line := range input {
		for j, ch := range line {
			switch ch {
			case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				continue
			default:
				numbersAround := extractNumbersAroundSymbolAt(&numbers, i, j)
				for _, number := range numbersAround {
					output += number
				}
			}
		}
	}

	return output
}

func part2(input []string) int {
	output := 0

	numbers := extractNumbers(input)
	for i, line := range input {
		for j, ch := range line {
			switch ch {
			case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				continue
			default:
				numbersAround := extractNumbersAroundSymbolAt(&numbers, i, j)

				if ch == '*' && len(numbersAround) == 2 {
					output += numbersAround[0] * numbersAround[1]
				}
			}
		}
	}

	return output
}

func main() {
	input, err := utils.ReadInputFile(2023, 3)
	if err != nil {
		log.Fatal(err)
	}
	answer1 := part1(input)
	fmt.Println("Answer to part 1:", answer1)
	answer2 := part2(input)
	fmt.Println("Answer to part 2:", answer2)
}
