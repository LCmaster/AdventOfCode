package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func part1(input []string) int {
    firstDigitMatcher := regexp.MustCompile("^[^0-9]*([0-9]{1}).*$")
    lastDigitMatcher := regexp.MustCompile("^.*([0-9]{1})[^0-9]*$")
    output := []int{}
    for _, line := range input {
        firstDigitMatcher := firstDigitMatcher.FindStringSubmatch(line)
        lastDigitMatcher := lastDigitMatcher.FindStringSubmatch(line)

        if firstDigitMatcher != nil || lastDigitMatcher != nil {
            firstDigitStr := firstDigitMatcher[1]
            lastDigitStr := lastDigitMatcher[1]

            resultStr := firstDigitStr + lastDigitStr
            result, err := strconv.Atoi(resultStr)

            if err == nil {
                output = append(output, result)
            }
        } 
    }

    result := 0

    for _, n := range output {
        result += n
    }

    return result
}

func part2(input []string) int {
    firstDigitMatcher := regexp.MustCompile("(zero|one|two|three|four|five|six|seven|eight|nine|[0-9]{1}).*")
    lastDigitMatcher := regexp.MustCompile(".*(zero|one|two|three|four|five|six|seven|eight|nine|[0-9]{1})")

    digitDictionary := map[string]string {
        "zero"  :   "0",
        "one"   :   "1",
        "two"   :   "2",
        "three" :   "3",
        "four"  :   "4",
        "five"  :   "5",
        "six"   :   "6",
        "seven" :   "7",
        "eight" :   "8",
        "nine"  :   "9",
    }

    output := []int{}
    for _, line := range input {
        firstDigitMatcher := firstDigitMatcher.FindStringSubmatch(line)
        lastDigitMatcher := lastDigitMatcher.FindStringSubmatch(line)

        if firstDigitMatcher != nil || lastDigitMatcher != nil {
            firstDigitStr := firstDigitMatcher[1]
            lastDigitStr := lastDigitMatcher[1]

            digitStr := []string{firstDigitStr, lastDigitStr}

            for i, entry := range digitStr{
                value, exists := digitDictionary[entry]
                if exists {
                    digitStr[i] = value
                }
            }

            resultStr := digitStr[0] + digitStr[1]
            result, err := strconv.Atoi(resultStr)

            if err == nil {
                output = append(output, result)
            }
        } 
    }

    result := 0

    for _, n := range output {
        result += n
    }

    return result
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
    fmt.Println("Answer to puzzle 1:", answer1)
    fmt.Println("Answer to puzzle 2:", answer2)
}
