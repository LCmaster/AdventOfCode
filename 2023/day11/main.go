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

func getEmptyRows(spaceMap [][]int) []int {
    result := []int{}
    for i, line := range spaceMap {
        diff := 0
        for _, tile := range line {
            diff += tile
        }
        if diff == 0 {
            result = append(result, i)
        }
    }
    
    return result
}

func part1(input []string, expansionRate int) int {
    spaceMap := generateMapFromInput(input)
    emptyRows := getEmptyRows(spaceMap)
    emptyColumns := getEmptyRows(transpose(spaceMap))

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
            g1Coords := []int{gallaxy[g1][0], gallaxy[g1][1]}
            g2Coords := []int{gallaxy[g2][0], gallaxy[g2][1]}

            g1Diff := make([]int, 2)
            g2Diff := make([]int, 2)
            for _, x := range emptyColumns {
                if x < g1Coords[0] {
                    g1Diff[0] += expansionRate-1
                }
                if x < g2Coords[0] {
                    g2Diff[0] += expansionRate-1
                }
            }

            for _, y := range emptyRows {
                if y < g1Coords[1] {
                    g1Diff[1] += expansionRate-1
                }
                if y < g2Coords[1] {
                    g2Diff[1] += expansionRate-1
                }
            }

            g1Coords[0] += g1Diff[0]
            g1Coords[1] += g1Diff[1]
            g2Coords[0] += g2Diff[0]
            g2Coords[1] += g2Diff[1]

            distX := math.Abs(float64(g2Coords[0] - g1Coords[0]))
            distY := math.Abs(float64(g2Coords[1] - g1Coords[1]))
            dist := int(distX + distY)

            output += dist
        }
    }
    return output
}

func part2(input []string, expansionRate int) int {
    spaceMap := generateMapFromInput(input)
    emptyRows := getEmptyRows(spaceMap)
    emptyColumns := getEmptyRows(transpose(spaceMap))

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
            g1Coords := []int{gallaxy[g1][0], gallaxy[g1][1]}
            g2Coords := []int{gallaxy[g2][0], gallaxy[g2][1]}

            g1Diff := make([]int, 2)
            g2Diff := make([]int, 2)
            for _, x := range emptyColumns {
                if x < g1Coords[0] {
                    g1Diff[0] += expansionRate-1
                }
                if x < g2Coords[0] {
                    g2Diff[0] += expansionRate-1
                }
            }

            for _, y := range emptyRows {
                if y < g1Coords[1] {
                    g1Diff[1] += expansionRate-1
                }
                if y < g2Coords[1] {
                    g2Diff[1] += expansionRate-1
                }
            }

            g1Coords[0] += g1Diff[0]
            g1Coords[1] += g1Diff[1]
            g2Coords[0] += g2Diff[0]
            g2Coords[1] += g2Diff[1]

            distX := math.Abs(float64(g2Coords[0] - g1Coords[0]))
            distY := math.Abs(float64(g2Coords[1] - g1Coords[1]))
            dist := int(distX + distY)

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

    fmt.Println("Answer to par 1", part1(input, 2))
    fmt.Println("Answer to par 2", part2(input, 1000000))
}
