package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input []string, maxRedAmount int, maxGreenAmount int, maxBlueAmount int) int {
    gameMatcher := regexp.MustCompile("^Game ([0-9]+): (.*;*)")
    groupMatcher := regexp.MustCompile("([0-9]+) (red|green|blue)")

    output := 0

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

        testRedCubeAmount := totalCubes["red"] <= maxRedAmount
        testGreenCubeAmount := totalCubes["green"] <= maxGreenAmount
        testBlueCubeAmount := totalCubes["blue"] <= maxBlueAmount

        if testRedCubeAmount && testGreenCubeAmount && testBlueCubeAmount {
            output += id
        }        
    }

    return output
}

func part2(input []string) int {
    gameMatcher := regexp.MustCompile("^Game ([0-9]+): (.*;*)")
    groupMatcher := regexp.MustCompile("([0-9]+) (red|green|blue)")

    output := 0

    for _, line := range input {
        gameMatch := gameMatcher.FindStringSubmatch(line)
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
        output += totalCubes["red"] * totalCubes["green"] * totalCubes["blue"]
    }

    return output
}



func main() {
    file, err := os.Open("./2023/day02/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    buffer := []string{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        buffer = append(buffer, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    possibleRedCubesAmount := 12
    possibleGreenCubesAmount := 13
    possibleBlueCubesAmount := 14

    answer1 := part1(buffer, possibleRedCubesAmount, possibleGreenCubesAmount, possibleBlueCubesAmount)
    fmt.Println("Answer to part 1:", answer1)
    answer2 := part2(buffer, possibleRedCubesAmount, possibleGreenCubesAmount, possibleBlueCubesAmount)
    fmt.Println("Answer to part 2:", answer2)
}
