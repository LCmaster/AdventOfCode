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
