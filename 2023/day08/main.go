package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

func extractNodes(input []string) map[string][]string{
    nodeMatcher := regexp.MustCompile("^([a-zA-Z0-9]{3}) = \\(([a-zA-Z0-9]{3}), ([a-zA-Z0-9]{3})\\)$")
    nodes := make(map[string][]string)
    for _, line := range input {
        match := nodeMatcher.FindStringSubmatch(strings.TrimSpace(line))
        if match != nil {
            nodes[match[1]] = []string{match[2], match[3]}
        }
    }

    return nodes
}

func gcd(a, b int) int {
    for b != 0 {
        t := b
        b = a % b
        a = t
    }

    return a
}

func lcm(a, b int) int {
    return a * b / gcd(a, b)
}


func part1(input []string) int {
    output := 0
    instructions := input[0]
    nodes := extractNodes(input)

    currentNode := "AAA"
    for true {
        for _, direction := range instructions {
            if currentNode == "ZZZ" {
                return output
            }

            if direction == 'L' {
                currentNode = nodes[currentNode][0]
            } else {
                currentNode = nodes[currentNode][1]
            }
            output++
        }
    }

    return 0
}

func part2(input []string) int {
    output := 0
    instructions := input[0]
    nodes := extractNodes(input)

    currentNodes := []string{}
    for node := range nodes {
        if node[2] == 'A' {
            currentNodes = append(currentNodes, node)
        }
    }

    if len(currentNodes) > 0 {
        stepsMultiplier := []int{}
        for _, node := range currentNodes {
            steps := 0
            Route:
            for true {
                for _, direction := range instructions {
                    if node[2] == 'Z' {
                        break Route
                    }
                    if direction == 'L' {
                        node = nodes[node][0]
                    } else {
                        node = nodes[node][1]
                    }

                    steps++
                }
            }
            stepsMultiplier = append(stepsMultiplier, steps)
        }
        result := stepsMultiplier[0]
        for i := 1; i < len(stepsMultiplier); i++ {
            result = lcm(result, stepsMultiplier[i])
        }
        output = result
    }

    return output
}


func main() {
    input, err := utils.ReadInputFile(2023, 8)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Answer to part 1:", part1(input))
    fmt.Println("Answer to part 2:", part2(input))
}
