package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

func extractNodes(input []string) map[string][]string{
    nodeMatcher := regexp.MustCompile("^([A-Z]{3}) = \\(([A-Z]{3}), ([A-Z]{3})\\)$")
    nodes := make(map[string][]string)
    for _, line := range input {
        match := nodeMatcher.FindStringSubmatch(strings.TrimSpace(line))
        if match != nil {
            nodes[match[1]] = []string{match[2], match[3]}
        }
    }

    return nodes
}

func part1(input []string) int {
    output := 0
    instructions := input[0]
    nodes := extractNodes(input)

    currentNode := "AAA"
    for currentNode != "ZZZ" {
        for _, direction := range instructions {
            if currentNode == "ZZZ" {
                break
            }

            if direction == 'L' {
                currentNode = nodes[currentNode][0]
            } else {
                currentNode = nodes[currentNode][1]
            }
            output++
        }
    }

    return output
}

func main() {
    input, err := utils.ReadInputFile(2023, 8)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Answer to part 1:", part1(input))
}
