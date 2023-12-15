package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

func extractGameSets(input []string) map[int][]int{
    output := make(map[int][]int)
    for _, result := range utils.Map[string, []int](input, func(line string, index int) []int {
        cubes := make(map[string]int)
        id := -1
        if gameId, games, ok := strings.Cut(line, ":"); ok {
            idStr := ""
            for _, ch := range gameId {
                if ch >= '0' && ch <= '9' {
                    idStr += string(ch)
                }
            }
            id, _ = strconv.Atoi(idStr)
            for _, game := range strings.Split(games, ";") {
                for _, group := range strings.Split(strings.TrimSpace(game), ",") {
                    part := strings.Split(strings.TrimSpace(group), " ")
                    color := part[1]
                    amount, _ := strconv.Atoi(part[0])
                    if cubes[color] < amount {
                        cubes[color] = amount
                    }
                }
            }
        }
        return []int{id, cubes["red"], cubes["green"], cubes["blue"]}
    }) {
        output[result[0]] = result[1:]
    }
    return output
}

func part1(input []string) int {
    output := 0
    maxAmount := []int{12, 13, 14}
    for id, cubes := range extractGameSets(input) {
        test := true
        for i, amount := range cubes {
            test = test && amount <= maxAmount[i]
        }
        if test {
            output += id
        }        
    }

    return output
}

func part2(input []string) int {
    output := 0
    for _, cubes := range extractGameSets(input) {
        output += cubes[0] * cubes[1] * cubes[2]
    }

    return output
}

func main() {
    input, err := utils.ReadInputFile(2023, 2)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Answer to part 1:", part1(input))
    fmt.Println("Answer to part 2:", part2(input))
}
