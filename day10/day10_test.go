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
		{in: []string{"18,1,0,161,255,137,254,252,14,95,165,33,181,168,2,188"}, want: 46600},
	}

	for _, c := range cases {
		got := solvePart1(convertPart1(c.in))
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{in: []string{"18,1,0,161,255,137,254,252,14,95,165,33,181,168,2,188"}, want: "23234babdc6afa036749cfa9b597de1b"},
	}

	for _, c := range cases {
		got := solvePart2(convertPart2(c.in))
		if got != c.want {
			t.Errorf("solvePart2(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestFinalInput(t *testing.T) {
	// Test on the final input
	got1 := solvePart1(convertPart1(helpers.ReadInput("input")))
	want1 := 46600
	if got1 != want1 {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got1, want1)
	}

	got2 := solvePart2(convertPart2(helpers.ReadInput("input")))
	want2 := "23234babdc6afa036749cfa9b597de1b"
	if got2 != want2 {
		t.Errorf("solvePart2(inputFile) == %v, want %v", got2, want2)
	}
}
