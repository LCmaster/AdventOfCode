package main

import (
	"fmt"
	"log"

	"github.com/LCmaster/AdventOfCode/utils"
)

func getDifferenceCount(line1, line2 string) int {
    diff := 0
    for i := 0; i < len(line1); i++ {
        if line1[i] != line2[i] {
            diff++
        }
    }
    return diff
}

func getColumn(buffer []string, col int) string {
    column := ""
    for j := 0; j < len(buffer); j++ {
        column += string(buffer[j][col])
    }
    return column
}

func checkMirroredColumns(buffer []string, smudge int) int {
    prevColumn := ""
    for i := 0; i < len(buffer[0]); i++ {
        column := getColumn(buffer, i)

        if prevColumn != "" {
            diff := getDifferenceCount(prevColumn, column)
            if diff <= smudge {
                j := i-2
                k := i+1
                for j >= 0 && k < len(buffer[0]) {
                    diff += getDifferenceCount(getColumn(buffer, j), getColumn(buffer, k))
                    j--
                    k++
                }
                if diff == smudge {
                    return i
                }
            }
        }
        prevColumn = column
    }
    
    return 0
}

func checkMirroredRows(buffer []string, smudge int) int {
    prevLine := ""
    for i := 0; i < len(buffer); i++ {
        line := buffer[i]

        if  prevLine != ""{
            diff := getDifferenceCount(prevLine, line)
            if diff <= smudge {
                j := i-2
                k := i+1
                for j >= 0 && k < len(buffer) {
                    diff += getDifferenceCount(buffer[j], buffer[k])
                    j--
                    k++
                }
                if diff == smudge {
                    return i
                }
            }
        }
        prevLine = line
    }

    return 0
}

func part1(input []string) int {
    output := 0
    buffer := []string{}

    for _, line := range input {
        if len(line) > 0 {
            buffer = append(buffer, line)
            continue
        } else {
            output += checkMirroredColumns(buffer, 0)
            output += 100 * checkMirroredRows(buffer, 0)
            buffer = []string{}
        }
    }

    return output
}

func part2(input []string) int {
    output := 0
    buffer := []string{}

    for _, line := range input {
        if len(line) > 0 {
            buffer = append(buffer, line)
            continue
        } else {
            output += checkMirroredColumns(buffer, 1)
            output += 100*checkMirroredRows(buffer, 1)
            buffer = []string{}
        }
    }

    return output
}


func main() {
    input, err := utils.ReadInputFile(2023, 13)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Answer to part 1:", part1(input))
    fmt.Println("Answer to part 2:", part2(input))
}
