package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

func generateKey(template string, groups[]int) string {
    keyParts := make([]string, len(groups)+1)
    keyParts[0] = template
    for i, group := range groups {
        keyParts[i+1] = strconv.Itoa(group)
    }
    return strings.Join(keyParts, ",")
}

func getVariations(template string, groups []int, cache *map[string]int) int {
    key := generateKey(template, groups)

    if cachedValue,ok := (*cache)[key]; ok {
        return cachedValue
    } 
    
    if len(template) == 0 {
        if len(groups) == 0 {
            return 1
        } else {
            return 0
        }
    }

    switch template[0] {
    case '?':
        res := getVariations(strings.Replace(template, "?", ".", 1), groups, cache)
        res += getVariations(strings.Replace(template, "?", "#", 1), groups, cache)
        return res
    case '.':
        res := getVariations(strings.TrimPrefix(template, "."), groups, cache)
        (*cache)[key] = res
        return res
    case '#':
        if len(groups) == 0 {
            (*cache)[key] = 0
            return 0
        }
        if len(template) < groups[0] {
            (*cache)[key] = 0
            return 0
        }
        if strings.Contains(template[0:groups[0]], ".") {
            (*cache)[key] = 0
            return 0
        }
        if len(groups) > 1 {
            if len(template) < groups[0]+1 || string(template[groups[0]]) == "#" {
                (*cache)[key] = 0
                return 0
            }
            res := getVariations(template[groups[0]+1:], groups[1:], cache)
            (*cache)[key] = res
            return res
        } else {
            res := getVariations(template[groups[0]:], groups[1:], cache)
            (*cache)[key] = res
            return res
        }
}

    return 0
}

func part1(input []string) int {
    output := 0
    cache := make(map[string]int)
    for _, line := range input {
        token := strings.Split(line, " ")
        groups := []int{}
        for _, num := range strings.Split(token[1], ",") {
            val, _ := strconv.Atoi(num)
            groups = append(groups, val)
        }

        output += getVariations(token[0], groups, &cache)
    }
    return output
}

func part2(input []string) int {
    newInput := make([]string, len(input))
    for i, line := range input {
        token := strings.Split(line, " ")
        template := make([]string, 5)
        groups := make([]string, 5)
        for j := 0; j < 5; j++ {
            template[j] = token[0]
            groups[j] = token[1]
        }
        newLine := strings.Join(template, "?")+" "+strings.Join(groups, ",")
        newInput[i] = newLine
    }
    return part1(newInput)
}

func main() {
    input, err := utils.ReadInputFile(2023, 12)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Answer to part 1", part1(input))
    fmt.Println("Answer to part 2", part2(input))
}
