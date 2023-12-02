package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInputFile(year int, day int) ([]string, error) {
    file, err := os.Open(fmt.Sprintf("./%d/day%02d/input.txt", year, day))
    if err != nil {
        return nil, err
    }
    defer file.Close()

    buffer := []string{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        buffer = append(buffer, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return buffer, nil
}
