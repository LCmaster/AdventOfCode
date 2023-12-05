package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

type Range struct {
    destination int
    source int
    length int
}

type RangeMap struct {
    destinationSlotType string
    sourceSlotType string
    ranges []Range
}

func extractSeeds(input []string) []int {
    seedsMatcher := regexp.MustCompile("^seeds:(( [0-9]+)+)")
    seeds := []int{}

    for _, line := range input {
        match := seedsMatcher.FindStringSubmatch(line)
        if match != nil {
            seedsStr := strings.Split(strings.TrimSpace(match[1]), " ")
            for _, seedStr := range seedsStr {
                seed, err := strconv.Atoi(seedStr)
                if err == nil {
                    seeds = append(seeds, seed)
                }
            }
        }
    }

    return seeds
}

func extractMaps(input []string) map[string]RangeMap {
    maps := make(map[string]RangeMap)
    mapMatcher := regexp.MustCompile("^([a-z]+)-to-([a-z]+) map:")
    rangeMatcher := regexp.MustCompile("^([0-9]+) ([0-9]+) ([0-9]+)$")
    
    i := 0
    currentMapName := ""
    for i < len(input) {
        line := strings.TrimSpace(input[i])
        if line == ""{
            i++
            continue
        }

        if currentMapName == "" {
            match := mapMatcher.FindStringSubmatch(line)
            if match == nil {
                i++
                continue
            }
            
            sourceStr := match[1]
            destinationStr := match[2]

            currentMapName = fmt.Sprintf("%s-%s", sourceStr, destinationStr)
            rangeMap := RangeMap{
                destinationSlotType: destinationStr,
                sourceSlotType: sourceStr,
                ranges: []Range{},
            }
            maps[currentMapName] = rangeMap

            i++
        } else {
            match := rangeMatcher.FindStringSubmatch(line)
            if match == nil {
                currentMapName = ""
                continue
            }
            
            destinationRangeStart, _ := strconv.Atoi(match[1])
            sourceRangeStart, _ := strconv.Atoi(match[2])
            rangeLength, _ := strconv.Atoi(match[3])

            mapRange := maps[currentMapName]

            mapRange.ranges = append(
                mapRange.ranges, 
                Range {
                    destination: destinationRangeStart,
                    source: sourceRangeStart,
                    length: rangeLength,
                },
            )
            maps[currentMapName] = mapRange

            i++
        }
    }

    return maps
}

func extractSeedsLocations(seeds []int, maps map[string]RangeMap) map[int]int {
    seedLocation := make(map[int]int)

    for _, seed := range seeds {
        source := "seed"
        destination := "location"

        location := seed
        for source != destination {
            for name, rangeMap := range maps {
                if strings.HasPrefix(name, source) {
                    Ranges:
                        for _, mapRange := range rangeMap.ranges {
                            if location >= mapRange.source && location <= (mapRange.source + mapRange.length) {
                                offset := location - mapRange.source
                                location = mapRange.destination + offset
                                break Ranges
                            }
                        }

                    token := strings.Split(name, "-")
                    source = token[1]
                }
            }
        }
        seedLocation[seed] = location
    }

    return seedLocation
}

func part1(input []string) int {
    seeds := extractSeeds(input)
    maps := extractMaps(input)
    seedsLocation := extractSeedsLocations(seeds, maps)

    lowestLocationNumber := -1
    for _, location := range seedsLocation {
        if lowestLocationNumber < 0 {
            lowestLocationNumber = location
        } else if location < lowestLocationNumber {
            lowestLocationNumber = location
        }
    }

    return lowestLocationNumber
}

func main() {
    input, err := utils.ReadInputFile(2023, 5)
    if err != nil {
        log.Fatal(err)
    }

    answer1 := part1(input)
    fmt.Println("Answer to part 1:", answer1)
}
