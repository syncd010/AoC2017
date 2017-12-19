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
			"b inc 5 if a > 1",
			"a inc 1 if b < 5",
			"c dec -10 if a >= 1",
			"c inc -20 if c == 10",
		}, want: 1},
	}

	for _, c := range cases {
		got := solvePart1(convert(c.in))
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{in: []string{
			"b inc 5 if a > 1",
			"a inc 1 if b < 5",
			"c dec -10 if a >= 1",
			"c inc -20 if c == 10",
		}, want: 10},
	}

	for _, c := range cases {
		got := solvePart2(convert(c.in))
		if got != c.want {
			t.Errorf("solvePart2(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestFinalInput(t *testing.T) {
	// Test on the final input
	got1 := solvePart1(convert(helpers.ReadInput("input")))
	want1 := 5966
	if got1 != want1 {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got1, want1)
	}

	got2 := solvePart2(convert(helpers.ReadInput("input")))
	want2 := 6347
	if got2 != want2 {
		t.Errorf("solvePart2(inputFile) == %v, want %v", got2, want2)
	}
}
