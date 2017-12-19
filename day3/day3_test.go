package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{in: 1, want: 0},
		{in: 12, want: 3},
		{in: 23, want: 2},
		{in: 1024, want: 31},
	}

	for _, c := range cases {
		got := solvePart1(c.in)
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}

	got := solvePart1(289326)
	want := 419
	if got != want {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{in: 55, want: 57},
		{in: 350, want: 351},
		{in: 23, want: 23},
		{in: 800, want: 806},
	}

	for _, c := range cases {
		got := solvePart2(c.in)
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}

	got := solvePart2(289326)
	want := 295229
	if got != want {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got, want)
	}
}
