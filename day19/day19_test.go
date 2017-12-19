package main

import (
	"testing"

	"github.com/syncd010/AoC2017/helpers"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{in: []string{
			"    |          ",
			"    |  +--+    ",
			"    A  |  C    ",
			"F---|--|-E---+ ",
			"    |  |  |  D ",
			"    +B-+  +--+ ",
		}, want: "ABCDEF"},
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
			"    |          ",
			"    |  +--+    ",
			"    A  |  C    ",
			"F---|--|-E---+ ",
			"    |  |  |  D ",
			"    +B-+  +--+ ",
		}, want: 38},
	}

	for _, c := range cases {
		got := solvePart2(convert(c.in))
		if got != c.want {
			t.Errorf("solvePart2(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestFinalInput(t *testing.T) {
	got1 := solvePart1(convert(helpers.ReadInput("input")))
	want1 := "DWNBGECOMY"
	if got1 != want1 {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got1, want1)
	}

	got2 := solvePart2(convert(helpers.ReadInput("input")))
	want2 := 17228
	if got2 != want2 {
		t.Errorf("solvePart2(inputFile) == %v, want %v", got2, want2)
	}
}
