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

