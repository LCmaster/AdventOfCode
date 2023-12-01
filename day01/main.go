package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(input []string) int {
    output := 0

    for _, line := range input {
        firstDigit := -1
        lastDigit := 0

        for _, ch := range line {
            value, err := strconv.Atoi(string(ch))
            if err == nil {
                if firstDigit < 0 {
                    firstDigit = value
                }
                lastDigit = value

            }
        }

        output += firstDigit * 10 + lastDigit
    }

    return output
}

func part2(input []string) int {
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

    output := 0
    for _, line := range input {
        firstDigit := -1
        lastDigit := 0

        for i, ch := range line {
            value, err := strconv.Atoi(string(ch))

            if err == nil {
                if firstDigit < 0 {
                    firstDigit = value
                }
                lastDigit = value
            } else {
                for name, value := range digitName {
                    if strings.HasPrefix(line[i:], name) {
                        if firstDigit < 0 {
                            firstDigit = value
                        }
                        lastDigit = value
                    }
                }
            }
        }

        output += firstDigit * 10 + lastDigit
    }

    return output
}

func main() {
    file, err := os.Open("./day01/input.txt")
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
