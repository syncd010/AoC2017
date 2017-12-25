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
		{in: []string{"", ""}, want: 0},
	}

	for _, c := range cases {
		got := solvePart1(convert(c.in))
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestFinalInput(t *testing.T) {
	// Test on the final input
	got1 := solvePart1(convert(helpers.ReadInput("inputTest")))
	want1 := 3
	if got1 != want1 {
		t.Errorf("Solve on Input Test == %v, want %v", got1, want1)
	}

	got2 := solvePart1(convert(helpers.ReadInput("input")))
	want2 := 2526
	if got2 != want2 {
		t.Errorf("Solve on Input == %v, want %v", got2, want2)
	}
}
