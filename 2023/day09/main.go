package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

func findNextValue(input []int) int {
    allZeros := true
    for _, val := range input {
        allZeros = allZeros && (val == 0)
    }
    if allZeros {
        return 0
    }

    indexOfLastInputElement := len(input)-1
    diff := make([]int, indexOfLastInputElement)
    for i := 0; i < indexOfLastInputElement; i++ {
        diff[i] = input[i+1] - input[i]
    }

    return input[indexOfLastInputElement] + findNextValue(diff)
}

func part1(input []string) int {
    output := 0

    for _, line := range input {
        line = strings.TrimSpace(line)
        if len(line) > 0 {
            tokens := strings.Split(line, " ")
            entry := []int{}
            for _, token := range tokens {
                number, err := strconv.Atoi(token)
                if err == nil {
                    entry = append(entry, number)
                }
            }

            output += findNextValue(entry)
        }
    }

    return output
}

func part2(input []string) int {
    output := 0

    for _, line := range input {
        line = strings.TrimSpace(line)
        if len(line) > 0 {
            tokens := strings.Split(line, " ")
            entry := []int{}
            for i := len(tokens)-1; i >= 0; i-- {
                number, err := strconv.Atoi(tokens[i])
                if err == nil {
                    entry = append(entry, number)
                }
            }

            output += findNextValue(entry)
        }
    }

    return output
}


func main() {
    input, err := utils.ReadInputFile(2023, 9)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Answer to part 1:", part1(input))
    fmt.Println("Answer to part 2:", part2(input))
}
