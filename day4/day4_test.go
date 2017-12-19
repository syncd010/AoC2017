package main

import (
	"testing"

	"github.com/syncd010/AoC2017/helpers"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{in: []string{
			"aa bb cc dd ee",  // Valid
			"aa bb cc dd aa",  // Not Valid
			"aa bb cc dd aaa", // Valid
		}, want: 2},
	}

	for _, c := range cases {
		got := solvePart1(convert(c.in))
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}

	got := solvePart1(convert(helpers.ReadInput("input")))
	want := 477
	if got != want {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{in: []string{
			"abcde fghij",              // Valid
			"abcde xyz ecdab",          // Not valid
			"a ab abc abd abf abj",     // Valid
			"iiii oiii ooii oooi oooo", // Valid
			"oiii ioii iioi iiio",      // Not valid
		}, want: 3},
	}

	for _, c := range cases {
		got := solvePart2(convert(c.in))
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}

	got := solvePart2(convert(helpers.ReadInput("input")))
	want := 167
	if got != want {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got, want)
	}
}
