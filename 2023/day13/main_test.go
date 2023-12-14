package main

import "testing"

func TestPart1(t *testing.T) {
    input := []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
	"",
    }
    output := part1(input)
    expected := 405

    if output != expected {
	t.Errorf("Day13 TestPart1 => Output: %d, Expected: %d", output, expected)
    }
}

func TestPart2(t *testing.T) {
    input := []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
	"",
    }
    output := part2(input)
    expected := 400

    if output != expected {
	t.Errorf("Day13 TestPart2 => Output: %d, Expected: %d", output, expected)
    }
}
