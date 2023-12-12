package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

type Record struct {
    template string
    inputs []int
}

func getVariations(line, prefix string, validator *regexp.Regexp) []string {
    result := []string{}
    if len(prefix) < len(line) {
        ch := line[len(prefix)]
        if ch == '?' {
            for _, token := range []string{".", "#"} {
                for _, variant := range getVariations(line, prefix+token, validator) {
                    if validator.MatchString(variant) {
                        result = append(result, variant)
                    }
                }
            }
        } else {
            token := string(ch)
            for _, variant := range getVariations(line, prefix+token, validator) {
                if validator.MatchString(variant) {
                    result = append(result, variant)
                }
            }
        }
    } else {
        if validator.MatchString(prefix) {
            result = append(result, prefix)
        }
    }

    return result
}

func part1(input []string) int {
    output := 0
    for _, line := range input {
        token := strings.Split(line, " ")
        groups := strings.Split(token[1], ",")
        totalBroken := 0
        regex := "\\.*"
        for i, group := range groups {
            if i > 0 {
                regex += "\\.+"
            }
            regex += "#{"+group+"}"

            val, err := strconv.Atoi(group)
            if err == nil {
                totalBroken += val
            }
        }
        regex += "\\.*"
        validator := regexp.MustCompile(regex)

        variants := getVariations(token[0], "", validator)
        for _, variant := range variants {
            broken := 0
            for _, ch := range variant {
                if ch == '#' {
                    broken++
                }
            }
            if broken == totalBroken {
                if validator.MatchString(variant) {
                    output++
                }
            }
        }
    }
    return output
}

func main() {
    input, err := utils.ReadInputFile(2023, 12)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Answer to part 1", part1(input))
}
