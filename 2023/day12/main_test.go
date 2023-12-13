package main

import "testing"

func TestPart1(t *testing.T) {
    input := []string{
        "???.### 1,1,3",
        ".??..??...?##. 1,1,3",
        "?#?#?#?#?#?#?#? 1,3,1,6",
        "????.#...#... 4,1,1",
        "????.######..#####. 1,6,5",
        "?###???????? 3,2,1",
    }
    output := part1(input)
    expected := 21

    if output != expected {
        t.Errorf("Day12 TestPart1 => Output: %d, Expected: %d", output, expected)
    }
}

func TestPart2(t *testing.T) {
    input := []string{
        "???.### 1,1,3",
        ".??..??...?##. 1,1,3",
        "?#?#?#?#?#?#?#? 1,3,1,6",
        "????.#...#... 4,1,1",
        "????.######..#####. 1,6,5",
        "?###???????? 3,2,1",
    }
    output := part2(input)
    expected := 525152

    if output != expected {
        t.Errorf("Day12 TestPart2 => Output: %d, Expected: %d", output, expected)
    }
}
