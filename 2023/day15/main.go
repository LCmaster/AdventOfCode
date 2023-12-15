package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/LCmaster/AdventOfCode/utils"
)

type Lens struct {
    slot int
    focalLength int
}

type Box map[string]Lens

func hash(input string) int {
    output := 0
    for _, ch := range input {
        output = (17 * (output + int(ch))) % 256
    }
    return output
}

func part1(input []string) int {
    output := 0
    for _, line := range input {
        if len(line) > 0 {
            for _, str := range strings.Split(strings.TrimSpace(line), ",") {
                output += hash(str)
            }
        }
    }
    return output
}

func part2(input []string) int {
    output := 0
    box := make([]Box, 256)
    for i := 0; i < len(box); i++ {
        box[i] = make(Box)
    }

    for _, line := range input {
        if len(line) > 0 {
            for _, str := range strings.Split(strings.TrimSpace(line), ",") {
                if str[len(str)-1] == '-' {
                    label := str[:len(str)-1]
                    boxId := hash(label)

                    lens, ok := box[boxId][label]
                    if ok {
                        delete(box[boxId], label)
                        for label, el := range box[boxId] {
                            if el.slot > lens.slot {
                                box[boxId][label] = Lens{slot: el.slot-1, focalLength: el.focalLength}
                            }
                        }
                    }
                } else if label, focalStr, ok := strings.Cut(str, "="); ok {
                    boxId := hash(label)
                    focalLength, _ := strconv.Atoi(focalStr)
                    
                    lens, ok := box[boxId][label]
                    if ok {
                        box[boxId][label] = Lens{slot: lens.slot, focalLength: focalLength}
                    } else {
                        box[boxId][label] = Lens{
                            slot: len(box[boxId]),
                            focalLength: focalLength,
                        }
                    }
                }
            }
        }
    }
    for i, b := range box {
        for _, lens := range b {
            output += (i+1) * (lens.slot + 1) * lens.focalLength
        }
    }

    return output
}

func main() {
    input, err := utils.ReadInputFile(2023, 15)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Answer to part 1:", part1(input))
    fmt.Println("Answer to part 2:", part2(input))
}

