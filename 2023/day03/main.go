package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/LCmaster/AdventOfCode/utils"
)



func part1(input []string) int {
    numberMatcher := regexp.MustCompile("[0-9]+")
    output := 0

    included := make(map[int]map[int]int)
    for i, line := range input {
        for j, ch := range line {
            switch ch {
                case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
                    continue
                default:
                    vOffset := -1
                    vCount := 3
                    if i == 0 {
                        vOffset = 0
                        vCount = 2
                    }
                    
                    for k := 0; k < vCount; k++ {
                        line = input[i+vOffset+k]
                        indices := numberMatcher.FindAllStringIndex(line, -1)
                        
                        for _, entry := range indices {
                            if len(entry) > 0 {
                                hOffset := -1
                                hCount := 3
                                if j == 0 {
                                    hOffset = 0
                                    hCount = 2
                                }

                                for l := 0; l < hCount; l++ {
                                    number, err := strconv.Atoi(line[entry[0]:entry[1]])
                                    if err == nil {
                                        _, ok := included[i+vOffset+k]
                                        if !ok {
                                            included[i+vOffset+k] = make(map[int]int)
                                        }
                                        if j+hOffset+l >= entry[0] && j+hOffset+l <= entry[1]-1 {
                                            included[i+vOffset+k][entry[0]] = number
                                        }
                                    }
                                } 
                            }
                        }
                    }
            }
        }
    }

    for _, row := range included {
        for _, value := range row {
            output += value
        }
    }

    return output
}

func part2(input []string) int {
    numberMatcher := regexp.MustCompile("[0-9]+")
    output := 0
    
    numbers := make(map[int]map[int]string)
    for i, line := range input {
        for _, entry := range numberMatcher.FindAllStringIndex(line, -1) {
            if len(entry) > 0 {
                _, ok := numbers[i]
                if !ok {
                    numbers[i] = make(map[int]string)
                }
                
                numbers[i][entry[0]] = line[entry[0]:entry[1]]
            }
        }
    }

    for i, line := range input {
        for j, ch := range line {
            switch ch {
                case '.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' :
                    continue
                default :
                    vOffset, hOffset := -1, -1
                    vCount, hCount := 3, 3
                    if i == 0 || i == len(input)-1 {
                        vCount = 2
                    }
                    if j == 0 || j == len(line)-1 {
                        hCount = 2
                    }
                    
                    numbersAround := []int{}
                    for x := 0; x < hCount; x++ {
                        for y := 0; y < vCount; y++ {
                            row := i+vOffset+y
                            column := j+hOffset+x
                            
                            for start, numberStr := range numbers[row] {
                                end := len(numberStr) + start -1
                                if column >= start && column <= end {
                                    number, err := strconv.Atoi(numberStr)
                                    if err == nil {
                                        numbersAround = append(numbersAround, number)
                                    }
                                    delete(numbers[row], start)
                                }
                            }
                            
                        }
                    }

                    if ch == '*' && len(numbersAround) == 2 {
                        output += numbersAround[0] * numbersAround[1]
                    } 
            }
        }
    }

    return output
}


func main() {
    input, err := utils.ReadInputFile(2023, 3)
    if err != nil {
        log.Fatal(err)
    }
    answer1 := part1(input)
    fmt.Println("Answer to part 1:", answer1)
    answer2 := part2(input)
    fmt.Println("Answer to part 2:", answer2)
}
