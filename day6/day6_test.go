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
		{in: []string{"0 2 7 0"}, want: 5},
	}

	for _, c := range cases {
		got := solvePart1(convert(c.in[0]))
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}

	got := solvePart1(convert(helpers.ReadInput("input")[0]))
	want := 4074
	if got != want {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{in: []string{"0 2 7 0"}, want: 4},
	}

	for _, c := range cases {
		got := solvePart2(convert(c.in[0]))
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}

	got := solvePart2(convert(helpers.ReadInput("input")[0]))
	want := 2793
	if got != want {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got, want)
	}
}
