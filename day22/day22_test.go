package main

import (
	"testing"

	"github.com/syncd010/AoC2017/helpers"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		in    []string
		steps int
		want  int
	}{
		{in: []string{"..#", "#..", "..."}, steps: 7, want: 5},
		{in: []string{"..#", "#..", "..."}, steps: 70, want: 41},
		{in: []string{"..#", "#..", "..."}, steps: 10000, want: 5587},
	}

	for _, c := range cases {
		_, got := simulate(convert(c.in), c.steps, stepPart1)
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in    []string
		steps int
		want  int
	}{
		{in: []string{"..#", "#..", "..."}, steps: 100, want: 26},
		{in: []string{"..#", "#..", "..."}, steps: 10000000, want: 2511944},
	}

	for _, c := range cases {
		_, got := simulate(convert(c.in), c.steps, stepPart2)
		if got != c.want {
			t.Errorf("solvePart2(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestFinalInput(t *testing.T) {
	// Test on the final input
	got1 := solvePart1(convert(helpers.ReadInput("input")))
	want1 := 5433
	if got1 != want1 {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got1, want1)
	}

	got2 := solvePart2(convert(helpers.ReadInput("input")))
	want2 := 2512599
	if got2 != want2 {
		t.Errorf("solvePart2(inputFile) == %v, want %v", got2, want2)
	}
}
