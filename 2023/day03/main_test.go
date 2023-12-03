package main

import "testing"

func TestPart1(t *testing.T) {
    input := []string{
        "467..114..",
        "...*......",
        "..35..633.",
        "......#...",
        "617*......",
        ".....+.58.",
        "..592.....",
        "......755.",
        "...$.*....",
        ".664.598..",
    }

    output := part1(input)
    expected := 4361

    if output != expected {
        t.Errorf("TestPart1: Output: %d, Expected: %d", output, expected)
    }
}

func TestPart2(t *testing.T) {
    input := []string{
        "467..114..",
        "...*......",
        "..35..633.",
        "......#...",
        "617*......",
        ".....+.58.",
        "..592.....",
        "......755.",
        "...$.*....",
        ".664.598..",
    }

    output := part2(input)
    expected := 467835

    if output != expected {
        t.Errorf("TestPart2: Output: %d, Expected: %d", output, expected)
    }
}
