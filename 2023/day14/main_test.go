package main

import "testing"

func TestPart1(t *testing.T) {
    input := []string {
        "O....#....",
        "O.OO#....#",
        ".....##...",
        "OO.#O....O",
        ".O.....O#.",
        "O.#..O.#.#",
        "..O..#O..O",
        ".......O..",
        "#....###..",
        "#OO..#....",
    }
    output := part1(input)
    expected := 136

    if output != expected {
        t.Errorf("Day14 TestPart1 => Output: %d, Expected: %d", output, expected)
    }
}

func TestPart2(t *testing.T) {
    input := []string {
        "O....#....",
        "O.OO#....#",
        ".....##...",
        "OO.#O....O",
        ".O.....O#.",
        "O.#..O.#.#",
        "..O..#O..O",
        ".......O..",
        "#....###..",
        "#OO..#....",
    }
    output := part2(input, 1000000000)
    expected := 64

    if output != expected {
        t.Errorf("Day14 TestPart2 => Output: %d, Expected: %d", output, expected)
    }
}
