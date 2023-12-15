package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

func slideNorth(input []string) []string {
    for i := 0; i < len(input); i++ {
        for j:= 0; j < len(input[i]); j++ {
            if i > 0 && input[i][j] == 'O' && input[i-1][j] == '.'{
                k := 1
                for (i-(k+1)) >= 0 && input[i-(k+1)][j] == '.' {
                    k++
                }
                input[i], input[i-k] = utils.SwitchStringCharacters(input[i], input[i-k], j, j)
            }
        }
    }
    return input
}

func slideWest(input []string) []string {
    for i := 0; i < len(input); i++ {
        for j := 0; j < len(input[i]); j++ {
            if j > 0 && input[i][j] == 'O' && input[i][j-1] == '.'{
                k := 1
                for (j-(k+1)) >= 0 && input[i][j-(k+1)] == '.' {
                    k++
                }
                input[i] = utils.SwitchCharacters(input[i], j, j-k)
            }
        }
    }
    return input
}

func slideSouth(input []string) []string {
    for i := len(input)-1; i >= 0; i-- {
        for j:= 0; j < len(input[i]); j++ {
            if i < len(input)-1 && input[i][j] == 'O' && input[i+1][j] == '.'{
                k := 1
                for (i+k+1) <= len(input)-1 && input[i+k+1][j] == '.' {
                    k++
                }
                input[i], input[i+k] = utils.SwitchStringCharacters(input[i], input[i+k], j, j)
            }
        }
    }
    return input
}

func slideEast(input []string) []string {
    for i := 0; i < len(input); i++ {
        for j := len(input[i]) - 1; j >= 0; j-- {
            if j < len(input[i])-1 && input[i][j] == 'O' && input[i][j+1] == '.'{
                k := 1
                for (j+k+1) <= len(input[i])-1 && input[i][j+k+1] == '.' {
                    k++
                }
                input[i] = utils.SwitchCharacters(input[i], j, j+k)
            }
        }
    }
    return input
}

func runCycle(input []string) []string {
    return slideEast(slideSouth(slideWest(slideNorth(input))))
}

func calculateLoad(input []string) int {
    result := 0
    for i, line := range input {
        result += strings.Count(line, "O") * (len(input) - i)
    }
    return result
}

func generateKey(input []string) string {
    return strings.Join(input, ",")
}

func part1(input []string) int {
    return calculateLoad(slideNorth(input))
}

func part2(input []string, cycles int) int {
    cache := make(map[string]string)
    key := generateKey(input)
    i := 0
    for i < cycles {
        if next, ok := cache[key]; ok {
            if i + len(cache) < cycles {
                i += len(cache)
            }else {
                i++
            }
            key = next
        } else {
            result := runCycle(strings.Split(key, ","))
            next := generateKey(result)
            cache[key] = next
            
            key = next
            i++
        }
    }
    return calculateLoad(strings.Split(key, ","))
}


func main() {
    input, err := utils.ReadInputFile(2023, 14)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Answer to part 1:", part1(input))
    fmt.Println("Answer to part 2:", part2(input, 1000000000))
}
