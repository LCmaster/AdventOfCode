package main

import "testing"

func TestPart1(t *testing.T) {
    input := []string{
        "...#......",
        ".......#..",
        "#.........",
        "..........",
        "......#...",
        ".#........",
        ".........#",
        "..........",
        ".......#..",
        "#...#.....",
    }

    output := part1(input)
    expected := 374

    if output != expected {
        t.Errorf("Day11 TestPart1 => Output: %d, Expected: %d", output, expected)
    }
}

func TestPart2(t *testing.T) {
    input := []string{
        "...#......",
        ".......#..",
        "#.........",
        "..........",
        "......#...",
        ".#........",
        ".........#",
        "..........",
        ".......#..",
        "#...#.....",
    }

    output := part2(input)
    expected := 8410

    if output != expected {
        t.Errorf("Day11 TestPart2 => Output: %d, Expected: %d", output, expected)
    }
}
