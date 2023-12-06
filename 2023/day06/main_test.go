package main

import "testing"

func TestPart1(t *testing.T) {
    input := []string{
        "Time:      7  15   30",
        "Distance:  9  40  200",
        "",
    }

    output := part1(input)
    expected := 288

    if output != expected {
        t.Errorf("Day06 TestPart1 => Output: %d, Expected: %d", output, expected)
    }
}

func TestPart2(t *testing.T) {
    input := []string{
        "Time:      7  15   30",
        "Distance:  9  40  200",
        "",
    }

    output := part2(input)
    expected := 71503

    if output != expected {
        t.Errorf("Day06 TestPart2 => Output: %d, Expected: %d", output, expected)
    }
}

