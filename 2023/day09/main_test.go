package main

import "testing"

func TestPart1(t *testing.T) {
    input := []string{
        "0 3 6 9 12 15",
        "1 3 6 10 15 21",
        "10 13 16 21 30 45",
    }

    output := part1(input)
    expected := 114
    
    if output != expected {
        t.Errorf("Day09 TestPart1 => Output: %d, Expected: %d", output, expected)
    }
}

func TestPart2(t *testing.T) {
    input := []string{
        "0 3 6 9 12 15",
        "1 3 6 10 15 21",
        "10 13 16 21 30 45",
    }

    output := part2(input)
    expected := 2
    
    if output != expected {
        t.Errorf("Day09 TestPart2 => Output: %d, Expected: %d", output, expected)
    }
}
