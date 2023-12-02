package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

func extractGameSets(input []string) map[int][]int{
    gameMatcher := regexp.MustCompile("^Game ([0-9]+): (.*;*)")
    groupMatcher := regexp.MustCompile("([0-9]+) (red|green|blue)")

    output := make(map[int][]int)

    for _, line := range input {
        gameMatch := gameMatcher.FindStringSubmatch(line)
        id, err := strconv.Atoi(gameMatch[1])
        if err != nil {
            continue
        }
        plays := gameMatch[2]

        totalCubes := map[string]int {"red": 0, "green": 0, "blue": 0}

        for _, play := range strings.Split(plays, ";") {
            for _, group := range strings.Split(play, ",") {
                groupMatch := groupMatcher.FindStringSubmatch(group)
                number, err := strconv.Atoi(groupMatch[1])
                if err != nil {
                    break
                }
                color := groupMatch[2]

                if number > totalCubes[color] {
                    totalCubes[color] = number
                }
            }
        }

        output[id] = []int{
            totalCubes["red"],
            totalCubes["green"],
            totalCubes["blue"],
        }       
    }

    return output

}

func part1(input []string, maxRedAmount int, maxGreenAmount int, maxBlueAmount int) int {
    gameSets := extractGameSets(input)
    output := 0

    for id, cubes := range gameSets {
        testRedCubeAmount := cubes[0] <= maxRedAmount
        testGreenCubeAmount := cubes[1] <= maxGreenAmount
        testBlueCubeAmount := cubes[2] <= maxBlueAmount

        if testRedCubeAmount && testGreenCubeAmount && testBlueCubeAmount {
            output += id
        }        
    }

    return output
}

func part2(input []string) int {
    gameSets := extractGameSets(input)
    output := 0

    for _, cubes := range gameSets {
        output += cubes[0] * cubes[1] * cubes[2]
    }

    return output
}



func main() {
    input, err := utils.ReadInputFile(2023, 2)
    if err != nil {
        log.Fatal(err)
    }
    
    possibleRedCubesAmount := 12
    possibleGreenCubesAmount := 13
    possibleBlueCubesAmount := 14

    answer1 := part1(input, possibleRedCubesAmount, possibleGreenCubesAmount, possibleBlueCubesAmount)
    fmt.Println("Answer to part 1:", answer1)
    answer2 := part2(input)
    fmt.Println("Answer to part 2:", answer2)
}
