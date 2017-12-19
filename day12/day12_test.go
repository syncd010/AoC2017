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
			"0 <-> 2",
			"1 <-> 1",
			"2 <-> 0, 3, 4",
			"3 <-> 2, 4",
			"4 <-> 2, 3, 6",
			"5 <-> 6",
			"6 <-> 4, 5",
		}, want: 6},
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
			"0 <-> 2",
			"1 <-> 1",
			"2 <-> 0, 3, 4",
			"3 <-> 2, 4",
			"4 <-> 2, 3, 6",
			"5 <-> 6",
			"6 <-> 4, 5",
		}, want: 2},
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
	want1 := 239
	if got1 != want1 {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got1, want1)
	}

	got2 := solvePart2(convert(helpers.ReadInput("input")))
	want2 := 215
	if got2 != want2 {
		t.Errorf("solvePart2(inputFile) == %v, want %v", got2, want2)
	}
}
