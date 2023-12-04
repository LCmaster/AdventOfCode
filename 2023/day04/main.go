package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

func part1(input []string) int {
    output:= 0

    for _, line := range input {
        parts := strings.Split(line, ":")
        if len(parts) == 2 {
            // card := parts[0]
            game := parts[1];
            numbers := strings.Split(game, "|")
            numbersToCheck := strings.Split(numbers[0], " ")
            winningNumbers := strings.Split(numbers[1], " ")

            cardPoints := 0
            for _, numberToCheckStr := range numbersToCheck {
                for _, winningNumberStr := range winningNumbers {
                    numberToCheck, err1 := strconv.Atoi(numberToCheckStr)
                    winningNumber, err2 := strconv.Atoi(winningNumberStr)

                    if err1 == nil && err2 == nil {
                        if numberToCheck == winningNumber {
                            if cardPoints == 0 {
                                cardPoints += 1
                            } else {
                                cardPoints *= 2
                            }
                        }
                    }
                }
            }

            output += cardPoints
        }
    }

    return output
}

func part2(input []string) int {
    output:= 0

    instances := make(map[int]int)
    for id, line := range input {
        gameId := id+1
        instances[gameId] += 1

        parts := strings.Split(line, ":")
        if len(parts) == 2 {
            // card := parts[0]
            game := parts[1];
            numbers := strings.Split(game, "|")
            numbersToCheck := strings.Split(numbers[0], " ")
            winningNumbers := strings.Split(numbers[1], " ")

            cardPoints := 0
            for _, numberToCheckStr := range numbersToCheck {
                for _, winningNumberStr := range winningNumbers {
                    numberToCheck, err1 := strconv.Atoi(numberToCheckStr)
                    winningNumber, err2 := strconv.Atoi(winningNumberStr)

                    if err1 == nil && err2 == nil {
                        if numberToCheck == winningNumber {
                            cardPoints += 1
                        }
                    }
                }
            }

            for i := 0; i < instances[gameId]; i++ {
                for j:= 0; j < cardPoints; j++ {
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
