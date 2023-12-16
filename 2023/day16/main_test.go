package main

import "testing"

func TestPart1(t *testing.T) {
    input := []string{
	`.|...\....`,
	`|.-.\.....`,
	`.....|-...`,
	`........|.`,
	`..........`,
	`.........\`,
	`..../.\\..`,
	`.-.-/..|..`,
	`.|....-|.\`,
	`..//.|....`,
    }
    output := part1(input)
    expected := 46
    if output != expected {
	t.Errorf("Day16 TestPart1 => Output: %d, Expected: %d", output, expected)
    }
}
