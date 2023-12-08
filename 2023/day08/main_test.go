package main

import "testing"

func TestPart1(t *testing.T) {
    input := []string{
        "LLR",
        "AAA = (BBB, BBB)",
        "BBB = (AAA, ZZZ)",
        "ZZZ = (ZZZ, ZZZ)",
        "",
    }

    output := part1(input)
    expected := 6

    if output != expected {
        t.Errorf("Day08 TestPart1 => Output: %d, Expected: %d", output, expected)
    }
}

func TestPart2(t *testing.T) {
    input := []string{
        "LR",
        "11A = (11B, XXX)",
        "11B = (XXX, 11Z)",
        "11Z = (11B, XXX)",
        "22A = (22B, XXX)",
        "22B = (22C, 22C)",
        "22C = (22Z, 22Z)",
        "22Z = (22B, 22B)",
        "XXX = (XXX, XXX)",
    }

    output := part2(input)
    expected := 6

    if output != expected {
        t.Errorf("Day08 TestPart2 => Output: %d, Expected: %d", output, expected)
    }
}

