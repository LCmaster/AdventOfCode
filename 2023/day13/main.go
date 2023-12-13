package main

import (
	"fmt"
	"log"

	"github.com/LCmaster/AdventOfCode/utils"
)

func getColumn(buffer []string, col int) string {
    column := ""
    for j := 0; j < len(buffer); j++ {
        column += string(buffer[j][col])
    }
    return column
}

func checkMirroredColumns(buffer []string) int {
    prevColumn := ""
    for i := 0; i < len(buffer[0]); i++ {
        column := getColumn(buffer, i)
        if column == prevColumn {
            symetric := true
            j := i-1
            k := i
            for j >= 0 && k < len(buffer[0]) {
                symetric = symetric && getColumn(buffer, j) == getColumn(buffer, k)
                j--
                k++
            }
            if symetric {
                return i
            }
        }
        prevColumn = column
    }
    
    return 0
}

func checkMirroredRows(buffer []string) int {
    prevLine := ""
    for i := 0; i < len(buffer); i++ {
        if buffer[i] == prevLine {
            symetric := true
            j := i-1
            k := i
            for j >= 0 && k < len(buffer) {
                symetric = symetric && buffer[j] == buffer[k]
                j--
                k++
            }
            if symetric {
                return i
            }
        }
        prevLine = buffer[i]
    }

    return 0
}

func part1(input []string) int {
    output := 0
    buffer := []string{}
    result := []int{}

    count := 0
    for _, line := range input {
        if len(line) > 0 {
            buffer = append(buffer, line)
            continue
        } else {
            result = append(result, checkMirroredColumns(buffer))
            result = append(result, checkMirroredRows(buffer))
            buffer = []string{}
            count++
        }
    }

    for i := 0; i < len(result); i += 2 {
        output += result[i] + (100 * result[i+1])
    }
    return output
}

func main() {
    input, err := utils.ReadInputFile(2023, 13)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Answer to part 1:", part1(input))
}
