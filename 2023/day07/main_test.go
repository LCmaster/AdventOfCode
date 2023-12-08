package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
    input := []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
	"",
    }

    output := part1(input)
    expected := 6440

    if output != expected {
	t.Errorf("Day07 TestPart1 => Output: %d, Expected: %d", output, expected)
    }
}

func TestPart2(t *testing.T) {
    input := []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
	"",
    }

    output := part2(input)
    expected := 5905

    if output != expected {
	t.Errorf("Day07 TestPart2 => Output: %d, Expected: %d", output, expected)
    }
}
