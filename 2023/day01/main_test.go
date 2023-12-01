package main

import "testing"


func TestPart1(t *testing.T) {
    input := []string{
        "1abc2",
        "pqr3stu8vwx",
        "a1b2c3d4e5f",
        "treb7uchet",
    }

    output := part1(input)
    expect := 142

    if output != expect {
        t.Errorf("Output: %d, Expected: %d", output, expect)
    }
}

func TestPart2(t *testing.T) {
    input := []string{
        "two1nine",
        "eightwothree",
        "abcone2threexyz",
        "xtwone3four",
        "4nineeightseven2",
        "zoneight234",
        "7pqrstsixteen",
    }

    output := part2(input)
    expect := 281

    if output != expect {
        t.Errorf("Output: %d, Expected: %d", output, expect)
    }
}

