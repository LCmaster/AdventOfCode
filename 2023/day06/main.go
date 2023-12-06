package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

func extractTimes(input []string) []int {
    times := []int{}

    for _, line := range input {
        line = strings.TrimSpace(line)
        if strings.HasPrefix(line, "Time:") {
            tokens := strings.Split(line, " ")
            for _, token := range tokens {
                number, err := strconv.Atoi(strings.TrimSpace(token))
                if err == nil {
                    times = append(times, number)
                }
            }
        }
    }

    return times
}

func extractDistances(input []string) []int {
    distances := []int{}
    
    for _, line := range input {
        line = strings.TrimSpace(line)
        if strings.HasPrefix(line, "Distance:") {
            tokens := strings.Split(line, " ")
            for _, token := range tokens {
                number, err := strconv.Atoi(strings.TrimSpace(token))
                if err == nil {
                    distances = append(distances, number)
                }
            }
        }
    }

    return distances
}

func part1(input []string) int {
    times := extractTimes(input)
    distances := extractDistances(input)

    output := 1
    for i := 0; i < len(times); i++ {
        time := times[i]
        recordDistance := distances[i]

        waysToWin := 0
        for j := 0; j < time; j++ {
            timeRemaining := time - j
            distance := timeRemaining * j
            if distance > recordDistance {
                waysToWin++
            }
        }

        output *= waysToWin
    }

    return output
}

func part2(input []string) int {
    times := extractTimes(input)
    distances := extractDistances(input)

    timeStr := ""
    distanceStr := ""
    for i := 0; i < len(times); i++ {
        timeStr += strconv.Itoa(times[i])
        distanceStr += strconv.Itoa(distances[i])
    }

    time, _ := strconv.Atoi(timeStr)
    recordDistance, _ := strconv.Atoi(distanceStr)

    waysToWin := 0
    for j := 0; j < time; j++ {
        timeRemaining := time - j
        distance := timeRemaining * j
        if distance > recordDistance {
            waysToWin++
        }
    }

    return waysToWin
}


func main() {
    input, err := utils.ReadInputFile(2023, 6)
    if err != nil {
        log.Fatal(err)
    }

    answer1 := part1(input)
    answer2 := part2(input)
    fmt.Println("Answer to part 1:", answer1)
    fmt.Println("Answer to part 2:", answer2)
}
