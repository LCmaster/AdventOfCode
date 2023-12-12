package main

import (
	"fmt"
	"log"
	"math"

	"github.com/LCmaster/AdventOfCode/utils"
)

func generateMapFromInput(input []string) [][]int {
    result := make([][]int, len(input))
    id := 1
    for i, line := range input {
        result[i] = make([]int, len(line))
        for j, ch := range line {
            if ch == '#' {
                result[i][j] = id
                id++
            }
        }
    }
    return result
}

func transpose(spaceMap [][]int) [][]int {
    result := make([][]int, len(spaceMap[0]))
    for i := 0; i < len(spaceMap[0]); i++ {
        result[i] = make([]int, len(spaceMap))
        for j := 0; j < len(spaceMap); j++ {
            result[i][j] = spaceMap[j][i]
        }
    }
    return result
}

func multiplyIdenticalLines(slice [][]int, factor int) [][]int {
    result := [][]int{}
    for _, line := range slice {
        diff := 0
        for _, tile := range line {
            diff += tile
        }
        result = append(result, line)
        if diff == 0 {
            for i := 0; i < factor - 1; i ++ {
                result = append(result, line)
            }
        }
    }
    return result
}

func generateRealMap(spaceMap [][]int, expansionFactor int) [][]int {
    result := multiplyIdenticalLines(spaceMap, expansionFactor)
    result = transpose(result)
    result = multiplyIdenticalLines(result, expansionFactor)
    return transpose(result)
}

func part1(input []string) int {
    spaceMap := generateRealMap(generateMapFromInput(input), 2)
    gallaxy := make(map[int][]int)
    gallaxyId := []int{}
    output := 0

    for y, line := range spaceMap {
        for x, cell := range line {
            if cell > 0 {
                gallaxyId = append(gallaxyId, cell)
                gallaxy[cell] = []int{x, y}
            }
        }
    }

    for i := 0; i < len(gallaxyId); i++ {
        for j := i+1; j < len(gallaxyId); j++ {
            g1 := gallaxyId[i]
            g2 := gallaxyId[j]
            dist := int(math.Abs(float64(gallaxy[g2][0] - gallaxy[g1][0])) + math.Abs(float64(gallaxy[g2][1] - gallaxy[g1][1])))
            // fmt.Printf("%d <=> %d = %d\n", g1, g2, dist)
            output += dist
        }
    }
    return output
}

func part2(input []string) int {
    spaceMap := generateRealMap(generateMapFromInput(input), 1000000)
    gallaxy := make(map[int][]int)
    gallaxyId := []int{}
    output := 0

    for y, line := range spaceMap {
        for x, cell := range line {
            if cell > 0 {
                gallaxyId = append(gallaxyId, cell)
                gallaxy[cell] = []int{x, y}
            }
        }
    }

    for i := 0; i < len(gallaxyId); i++ {
        for j := i+1; j < len(gallaxyId); j++ {
            g1 := gallaxyId[i]
            g2 := gallaxyId[j]
            dist := int(math.Abs(float64(gallaxy[g2][0] - gallaxy[g1][0])) + math.Abs(float64(gallaxy[g2][1] - gallaxy[g1][1])))
            output += dist
        }
    }
    return output
}


func main() {
    input, err := utils.ReadInputFile(2023, 11)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Answer to par 1", part1(input))
    fmt.Println("Answer to par 2", part2(input))
}
