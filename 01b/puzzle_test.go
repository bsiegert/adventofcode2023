package main

import "testing"

func TestFindDigits(t *testing.T) {
	tests := []struct {
		line string
		num string
	}{
		{
			"1abc2",
			"12",
		}, {
			"pqr3stu8vwx",
			"38",
		}, {
			"a1b2c3d4e5f",
			"15",
		}, {
			"treb7uchet",
			"77",
		}, {
			"two1nine",
			"29",
		}, {
			"eightwothree",
			"83",
		}, {
			"abcone2threexyz",
			"13",
		}, {
			"xtwone3four",
			"24",
		}, {
			"4nineeightseven2",
			"42",
		}, {
			"zoneight234",
			"14",
		}, {
			"7pqrstsixteen",
			"76",
		},
	}
	for _, tc := range tests {
		got, want := string(findDigits(tc.line)), tc.num
		if got != want {
			t.Errorf("findDigits(%q) = %q, want %q", tc.line, got, want)
		}
	}
}
