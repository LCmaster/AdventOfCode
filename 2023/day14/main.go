package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

func part1(input []string) int {
    output := 0
    for i := 0; i < len(input); i++ {
        for j:= 0; j < len(input[i]); j++ {
            if i > 0 && input[i][j] == 'O' && input[i-1][j] == '.'{
                k := 1
                for (i-(k+1)) >= 0 && input[i-(k+1)][j] == '.' {
                    k++
                }
                south := []rune(input[i])
                north := []rune(input[i-k])
                south[j] = '.'
                north[j] = 'O'
                input[i] = string(south)
                input[i-k] = string(north)
            }
        }
    }
    for i, line := range input {
        output += strings.Count(line, "O") * (len(input) - i)
    }
    return output
}

func main() {
    input, err := utils.ReadInputFile(2023, 14)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Answer to part 1:", part1(input))
}
