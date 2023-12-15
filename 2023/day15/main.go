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

func hash(input string, initialValue int) int {
    output := 17 * (initialValue + int(input[0])) % 256 
    if len(input) > 1 {
        return hash(input[1:], output)
    } 
    return output
}

func part1(input []string) int {
    output := 0
    for _, line := range input {
        if len(line) > 0 {
            for _, str := range strings.Split(strings.TrimSpace(line), ",") {
                output += hash(str, 0)
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
                    boxId := hash(label, 0)

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
                    boxId := hash(label, 0)
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

