package main

import (
	"fmt"
	"log"

	"github.com/LCmaster/AdventOfCode/utils"
)

func getStartingPosition(input []string) (int, int) {
    for y, line := range input {
        for x, ch := range line {
            if ch == 'S' {
                return x, y
            }
        }
    }

    return -1, -1
}

func getAdjacentConnectingTiles(tile rune, x, y int) [][]int {
    tiles := [][]int{}

    switch tile {
    case 'F':
        tiles = append(tiles, []int{x+1, y})
        tiles = append(tiles, []int{x, y+1})
    case '7':
        tiles = append(tiles, []int{x-1, y})
        tiles = append(tiles, []int{x, y+1})
    case 'L':
        tiles = append(tiles, []int{x, y-1})
        tiles = append(tiles, []int{x+1, y})
    case 'J':
        tiles = append(tiles, []int{x-1, y})
        tiles = append(tiles, []int{x, y-1})
    case '-':
        tiles = append(tiles, []int{x-1, y})
        tiles = append(tiles, []int{x+1, y})
    case '|', 'S':
        tiles = append(tiles, []int{x, y-1})
        tiles = append(tiles, []int{x, y+1})
    }
    
    return tiles
}

func saveStep(x, y, step int, savedSteps *map[int]map[int]int) bool {
    _, ok := (*savedSteps)[x]
    if !ok {
        (*savedSteps)[x] = make(map[int]int)
    }

    _, ok = (*savedSteps)[x][y]
    if !ok {
        (*savedSteps)[x][y] = step
        return true
    }
    return false
}

func findNextStep(x, y, step int, input []string, savedSteps *map[int]map[int]int) [][]int {
    adjacentPipes := getAdjacentConnectingTiles(
        rune(input[y][x]), 
        x, 
        y)
    savedPipes := [][]int{}
    for _, coords := range adjacentPipes {
        if saveStep(coords[0], coords[1], step+1, savedSteps) {
            savedPipes = append(savedPipes, []int{coords[0], coords[1], step+1})
        }
    }
    return savedPipes
}

func part1(input []string) int {
    output := 0
    
    x, y := getStartingPosition(input)
    pipeDistance := make(map[int]map[int]int)
    pipeDistance[x] = make(map[int]int)
    pipeDistance[x][y] = 0

    savedPipes := [][]int{
        []int{x, y, 0},
    }
    for len(savedPipes) > 0 {
        newSavedPipes := [][]int{}
        for _, coords := range savedPipes {
            newSavedPipes = append(
                newSavedPipes, 
                findNextStep(coords[0], coords[1], coords[2], input, &pipeDistance)...
            )
        }
        savedPipes = newSavedPipes
    }
    for _, rowMap := range pipeDistance {
        for _, step := range rowMap {
            if step > output {
                output = step
            }
        }
    }
    
    return output
}

func part2(input []string) int {
    output := 0
    
    x, y := getStartingPosition(input)
    pipeDistance := make(map[int]map[int]int)
    pipeDistance[x] = make(map[int]int)
    pipeDistance[x][y] = 0

    savedPipes := [][]int{
        []int{x, y, 0},
    }
    for len(savedPipes) > 0 {
        newSavedPipes := [][]int{}
        for _, coords := range savedPipes {
            newSavedPipes = append(
                newSavedPipes, 
                findNextStep(coords[0], coords[1], coords[2], input, &pipeDistance)...
            )
        }
        savedPipes = newSavedPipes
    }

    for i, line := range input {
        intersectionCount := 0
        for j, ch := range line {
            if (x == j && y == i) || pipeDistance[j][i] > 0 {
                if ch == 'F' || ch == '7' || ch == '|' || ch == 'S' {
                    intersectionCount++
                }
            } else {
                if intersectionCount % 2 == 1 {
                    output++
                }
            }
        }
    }
    
    return output
}


func main() {
    input, err := utils.ReadInputFile(2023, 10)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Answer to part 1", part1(input))
    fmt.Println("Answer to part 2", part2(input))
}

