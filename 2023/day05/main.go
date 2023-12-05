package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/LCmaster/AdventOfCode/utils"
)

type Range struct {
    destination int
    source int
    length int
}

type RangeMap struct {
    destination string
    source string
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

func extractMaps(input []string) []RangeMap {
    maps := []RangeMap{}
    mapMatcher := regexp.MustCompile("^([a-z]+)-to-([a-z]+) map:")
    rangeMatcher := regexp.MustCompile("^([0-9]+) ([0-9]+) ([0-9]+)$")
    
    i := 0
    currentMapIndex := -1
    for i < len(input) {
        line := strings.TrimSpace(input[i])
        if line == ""{
            i++
            continue
        }

        if currentMapIndex < 0 {
            match := mapMatcher.FindStringSubmatch(line)
            if match == nil {
                i++
                continue
            }
            
            sourceStr := match[1]
            destinationStr := match[2]

            rangeMap := RangeMap{
                destination: destinationStr,
                source: sourceStr,
                ranges: []Range{},
            }
            maps = append(maps, rangeMap)
            currentMapIndex = len(maps)-1

            i++
        } else {
            match := rangeMatcher.FindStringSubmatch(line)
            if match == nil {
                currentMapIndex = -1
                continue
            }
            
            destinationRangeStart, _ := strconv.Atoi(match[1])
            sourceRangeStart, _ := strconv.Atoi(match[2])
            rangeLength, _ := strconv.Atoi(match[3])

            mapRange := maps[currentMapIndex]

            mapRange.ranges = append(
                mapRange.ranges, 
                Range {
                    destination: destinationRangeStart,
                    source: sourceRangeStart,
                    length: rangeLength,
                },
            )
            maps[currentMapIndex] = mapRange

            i++
        }
    }

    return maps
}

func extractSeedLocation(seed int, maps []RangeMap) int {
    source := "seed"
    destination := "location"

    location := seed
    for source != destination {
        for _, rangeMap := range maps {
            if rangeMap.source == source {
                destinationLocation := location
                for _, mapRange := range rangeMap.ranges {
                    if location >= mapRange.source && location < (mapRange.source + mapRange.length) {
                        offset := location - mapRange.source
                        destinationLocation = mapRange.destination + offset
                    }
                }
                location = destinationLocation
                source = rangeMap.destination
            }
        }
    }

    return location
}

func extractLowestLocationFromSeedRange(start int, length int, maps []RangeMap, channel chan<- int) {
    lowestLocation := math.MaxInt
    for i := 0; i < length; i++ {
        var location int = extractSeedLocation(start + i, maps)
        if location < lowestLocation {
            lowestLocation = location
        }
    }

    channel <- lowestLocation
}

func part1(input []string) int {
    seeds := extractSeeds(input)
    maps := extractMaps(input)

    lowestLocationNumber := math.MaxInt
    for _, seed := range seeds {
        location := extractSeedLocation(seed, maps)
        if location < lowestLocationNumber {
            lowestLocationNumber = location
        }
    }

    return lowestLocationNumber
}

func part2(input []string) int {
    seedRanges := extractSeeds(input)
    maps := extractMaps(input)

    var wg sync.WaitGroup
    locationChannel := make(chan int)
    for i := 0; i < len(seedRanges); i += 2 {
        wg.Add(1)
        start := seedRanges[i]
        length := seedRanges[i+1]
        go func() {
            defer wg.Done()
            extractLowestLocationFromSeedRange(start, length, maps, locationChannel)
        }()
    }
    go func(){
        defer close(locationChannel)
        wg.Wait()
    }()
    
    lowestLocationNumber := math.MaxInt
    for location := range locationChannel {
        if location < lowestLocationNumber {
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
    answer2 := part2(input)
    fmt.Println("Answer to part 2:", answer2)
}
