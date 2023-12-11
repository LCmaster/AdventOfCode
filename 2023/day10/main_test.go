package main

import "testing"

func TestPart1(t *testing.T) {
    input := []string{
        "..F7.",
        ".FJ|.",
        "SJ.L7",
        "|F--J",
        "LJ...",
    }
    output := part1(input)
    expected := 8

    if output != expected {
        t.Errorf("Day10 TestPart1 => Output: %d, Expected: %d", output, expected)
    }
}

func TestPart2(t *testing.T) {
    input := []string{
        ".F----7F7F7F7F-7....",
        ".|F--7||||||||FJ....",
        ".||.FJ||||||||L7....",
        "FJL7L7LJLJ||LJ.L-7..",
        "L--J.L7...LJS7F-7L7.",
        "....F-J..F7FJ|L7L7L7",
        "....L7.F7||L7|.L7L7|",
        ".....|FJLJ|FJ|F7|.LJ",
        "....FJL-7.||.||||...",
        "....L---J.LJ.LJLJ...",
    }
    output := part2(input)
    expected := 8

    if output != expected {
        t.Errorf("Day10 TestPart2 => Output: %d, Expected: %d", output, expected)
    }
}
