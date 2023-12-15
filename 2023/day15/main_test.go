package main

import "testing"

func TestPart1(t *testing.T) {
    input := []string{
	"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
    }
    output := part1(input)
    expected := 1320

    if output != expected {
	t.Errorf("Day15 TestPart1 => Output: %d, Expected: %d", output, expected)
    }
}

func TestPart2(t *testing.T) {
    input := []string{
	"rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
    }
    output := part2(input)
    expected := 145

    if output != expected {
	t.Errorf("Day15 TestPart2 => Output: %d, Expected: %d", output, expected)
    }
}
