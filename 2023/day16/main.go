package main

import (
	"fmt"
	"log"
	"math"
	"sync"

	"github.com/LCmaster/AdventOfCode/utils"
)

type any interface{}
var empty any


var (
    Up = [2]int{0, -1}
    Down = [2]int{0, 1}
    Left = [2]int{-1, 0}
    Right = [2]int{1, 0}
)

type Beam struct {
    x, y int
    direction [2]int
}

func (b *Beam) LookLeft() {
    b.direction = Left
}

func (b *Beam) LookRight() {
    b.direction = Right
}

func (b *Beam) LookUp() {
    b.direction = Up
}

func (b *Beam) LookDown() {
    b.direction = Down
}

func (b *Beam) Move() {
    b.x += b.direction[0]
    b.y += b.direction[1]
}

func generateGrid(input []string) [][]rune {
    grid := [][]rune{}
    for _, line := range input {
        grid = append(grid, []rune(line))
    }
    return grid
}

func getBeamEnergyGrid(beam *Beam, grid *[][]rune, wg *sync.WaitGroup, ch chan [2]int, cache *map[string]any, lock *sync.Mutex) {
    for beam.x >= 0 && beam.x < len((*grid)[0]) && beam.y >= 0 && beam.y < len(*grid)  {
        infiniteLoop := false
        key := fmt.Sprintf("(%d,%d) => [%d,%d]", beam.x, beam.y, beam.direction[0], beam.direction[1])
        
        lock.Lock()
        if _, ok := (*cache)[key]; ok {
            infiniteLoop = true
        } else {
            (*cache)[key] = empty
        }
        lock.Unlock()

        if infiniteLoop {
            break
        }

        ch <- [2]int{beam.x, beam.y}
        switch (*grid)[beam.y][beam.x] {
        case '/':
            if beam.direction == Right {
                beam.LookUp()
            } else if beam.direction == Left {
                beam.LookDown()
            } else if beam.direction == Down {
                beam.LookLeft()
            } else {
                beam.LookRight()
            }
        case '\\' :
            if beam.direction == Right {
                beam.LookDown()
            } else if beam.direction == Left {
                beam.LookUp()
            } else if beam.direction == Down {
                beam.LookRight()
            } else {
                beam.LookLeft()
            }
        case '-':
            if beam.direction == Up || beam.direction == Down{
                beam.LookLeft()
                wg.Add(1)
                go getBeamEnergyGrid(&Beam{beam.x+1, beam.y, Right}, grid, wg, ch, cache, lock)
            }
        case '|':
            if beam.direction == Left || beam.direction == Right{
                beam.LookUp()
                wg.Add(1)
                go getBeamEnergyGrid(&Beam{beam.x, beam.y+1, Down}, grid, wg, ch, cache, lock)
            }
        }
        beam.Move()
    }
    wg.Done()
}

func countEnergizedTiles(x, y int, direction [2]int, grid *[][]rune) int {
    var cache = make(map[string]any)
    var lock sync.Mutex

    energyGrid := make([][]bool, len(*grid))
    for y, row := range *grid {
        energyGrid[y] = make([]bool, len(row))
    }

    var wg sync.WaitGroup
    ch := make(chan [2]int)
    wg.Add(1)
    go getBeamEnergyGrid(&Beam{x, y, direction}, grid, &wg, ch, &cache, &lock)
    go func() {
        wg.Wait()
        close(ch)
    }()

    output := 0
    for coords := range ch {
        x, y := coords[0], coords[1]
        if x >= 0 && x < len(*grid) && y >= 0 && y < len((*grid)[0]) {
            if !energyGrid[y][x] {
                output++
            }
            energyGrid[y][x] = true
        }
    }

    return output
}

func part1(input []string) int {
    grid := generateGrid(input)
    return countEnergizedTiles(0, 0, Right, &grid)
}

func part2(input []string) int {
    grid := generateGrid(input)
    output := 0
    for y := 0; y < len(grid); y++ {
        for x := 0; x < len(grid[0]); x++ {
            result := 0
            if x == 0 && y == 0 {
                res1 := float64(countEnergizedTiles(x, y, Right, &grid))
                res2 := float64(countEnergizedTiles(x, y, Down, &grid))
                result = int(math.Max(res1, res2))
            } else if x == len(grid[0])-1 && y == 0 {
                res1 := float64(countEnergizedTiles(x, y, Left, &grid))
                res2 := float64(countEnergizedTiles(x, y, Down, &grid))
                result = int(math.Max(res1, res2))
            } else if x == 0 && y == len(grid) - 1 {
                res1 := float64(countEnergizedTiles(x, y, Left, &grid))
                res2 := float64(countEnergizedTiles(x, y, Down, &grid))
                result = int(math.Max(res1, res2))
            } else if x == 0 && y == len(grid) - 1 {
                res1 := float64(countEnergizedTiles(x, y, Left, &grid))
                res2 := float64(countEnergizedTiles(x, y, Down, &grid))
                result = int(math.Max(res1, res2))
            } else {
                if x == 0 {
                    result = countEnergizedTiles(x, y, Right, &grid)
                }
                if y == 0 {
                    result = countEnergizedTiles(x, y, Down, &grid)
                }
                if y == len(grid) - 1 {
                    result = countEnergizedTiles(x, y, Left, &grid)
                }
                if x == len(grid[0]) - 1 {
                    result = countEnergizedTiles(x, y, Up, &grid)
                }
            }

            if result > output {
                output = result
            }
        }
    }
    return output
}

func main() {
    input, err := utils.ReadInputFile(2023, 16)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Answer to part 1:", part1(input))
    fmt.Println("Answer to part 2:", part2(input))
}

