package main

import (
	"fmt"
	"log"
	"math"
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

func bhaskaraSolver(a float64, b float64, c float64) []float64 {
    output := make([]float64, 2)
    output[0] = (-b + math.Sqrt(math.Pow(b, 2) - 4*a*c))/(2*a)
    output[1] = (-b - math.Sqrt(math.Pow(b, 2) - 4*a*c))/(2*a)
    return output
}

func part1(input []string) int {
    times := extractTimes(input)
    distances := extractDistances(input)

    output := 1
    for i := 0; i < len(times); i++ {
        limit := bhaskaraSolver(-1, float64(times[i]), -float64(distances[i]+1))
        output *= int(math.Floor(limit[1])) - int(math.Ceil(limit[0])) + 1
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
    distance, _ := strconv.Atoi(distanceStr)

    output := 1
    limit := bhaskaraSolver(-1, float64(time), -float64(distance+1))
    output *= int(math.Floor(limit[1])) - int(math.Ceil(limit[0])) + 1

    return output
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
