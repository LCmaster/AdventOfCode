package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func extractFirstAndLastDigits(input string, checkForDigitNames bool) (int, int) {
    digitName := map[string]int {
        "zero"  :   0,
        "one"   :   1,
        "two"   :   2,
        "three" :   3,
        "four"  :   4,
        "five"  :   5,
        "six"   :   6,
        "seven" :   7,
        "eight" :   8,
        "nine"  :   9,
    }

    first := -1
    last := 0

    for i, ch := range input {
        value, err := strconv.Atoi(string(ch))
        if err == nil {
            if first < 0 {
                first = value
            }
            last = value
        } else if checkForDigitNames {
            for name, value := range digitName {
                if strings.HasPrefix(input[i:], name) {
                    if first < 0 {
                        first = value
                    }
                    last = value
                }
            }
        }
    }
    return first, last
}

func part1(input []string) int {
    output := 0

    for _, line := range input {
        first, last := extractFirstAndLastDigits(line, false)
        output += first * 10 + last
    }

    return output
}

func part2(input []string) int {
    output := 0

    for _, line := range input {
        first, last := extractFirstAndLastDigits(line, true)        
        output += first * 10 + last
    }

    return output
}

func main() {
    file, err := os.Open("./2023/day01/input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    buffer := []string{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        buffer = append(buffer, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    answer1 := part1(buffer)
    answer2 := part2(buffer)
    fmt.Println("Answer to part 1:", answer1)
    fmt.Println("Answer to part 2:", answer2)
}
