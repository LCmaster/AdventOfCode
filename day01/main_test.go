package main

import "testing"


func Test_part1(t *testing.T) {
    input := []string{
        "1abc2",
        "pqr3stu8vwx",
        "a1b2c3d4e5f",
        "treb7uchet",
    }

    output := part1(input)
    expect := 142

    if output != expect {
        t.Errorf("Output: %q, Expected: %q", output, expect)
    }
}

