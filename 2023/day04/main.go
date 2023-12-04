package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

func extractWinningNumbers(input []string) [][]int {
    output := make([][]int, len(input))

    for lineIndex, line := range input {
        line = strings.TrimSpace(line)
        if len(line) > 0 {
            parts := strings.Split(line, ":")
            if len(parts) == 2 {
                game := parts[1];
                numbers := strings.Split(game, "|")
                numbersToCheck := strings.Split(numbers[0], " ")
                winningNumbers := strings.Split(numbers[1], " ")

                output[lineIndex] = []int{}
                for _, numberToCheckStr := range numbersToCheck {
                    for _, winningNumberStr := range winningNumbers {
                        numberToCheck, err1 := strconv.Atoi(numberToCheckStr)
                        winningNumber, err2 := strconv.Atoi(winningNumberStr)

                        if err1 == nil && err2 == nil {
                            if numberToCheck == winningNumber {
                                output[lineIndex] = append(output[lineIndex], winningNumber)
                            }
                        }
                    }
                }
            }
        }
    }

    return output
}

func part1(input []string) int {
    output:= 0
    
    games := extractWinningNumbers(input)
    for _, game := range games {
        points := len(game)
        if points > 0 {
            output += int(math.Pow(2, float64(points-1)))
        }
    }

    return output
}

func part2(input []string) int {
    output:= 0
    instances := make(map[int]int)
    games := extractWinningNumbers(input)
    for gameId, game := range games {
        instances[gameId] += 1
        points := len(game)
        if points > 0 {
            for i := 0; i < instances[gameId]; i++ {
                for j:= 0; j < points; j++ {
                    instances[gameId+j+1] += 1
                }
            } 
        }
    }

    for _, count := range instances {
        output += count
    }

    return output
}


func main() {
    input, err := utils.ReadInputFile(2023, 4)
    if err != nil {
        log.Fatal(err)
    }

    answer1 := part1(input)
    answer2 := part2(input)
    fmt.Println("Answer to part 1:", answer1)
    fmt.Println("Answer to part 2:", answer2)
}
