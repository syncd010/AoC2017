package main

import (
	"testing"

	"github.com/syncd010/AoC2017/helpers"
)

func TestPatternOps(t *testing.T) {
	cases := []struct {
		op   func(string) string
		in   string
		want string
	}{
		{op: transpose, in: "#..#", want: "#..#"},
		{op: transpose, in: "#..#.#..#", want: "##.....##"},
		{op: flipH, in: "#..#", want: ".##."},
		{op: flipH, in: "#..#.#..#", want: "..##.##.."},
		{op: flipV, in: "#...", want: "..#."},
		{op: flipV, in: "#..#.#..#", want: "..##.##.."},
	}

	for _, c := range cases {
		got := c.op(c.in)
		if got != c.want {
			t.Errorf("TestTranspose(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestConvert(t *testing.T) {
	cases := []struct {
		in   []string
		want []string
	}{
		{in: []string{"#./.. => 0"},
			want: []string{
				"#...",
				".#..",
				"..#.",
				"...#",
			}},
		{in: []string{".#./..#/### => 0"},
			want: []string{
				"####...#.",
				"###..#.#.",
				"..##.#.##",
				".###.#..#",
				".#...####",
				".#.#..###",
				"##.#.##..",
				"#..#.###.",
			}},
	}

	for _, c := range cases {
		got := convert(c.in)
		for _, want := range c.want {
			if _, ok := got[want]; !ok {
				t.Errorf("TestConvert. Didn;t find %v", want)
			}
		}
	}
}

func TestSolve(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{in: []string{
			"../.# => ##./#../...",
			".#./..#/### => #..#/..../..../#..#",
		}, want: 12},
	}

	for _, c := range cases {
		got := solve(convert(c.in), 2)
		if got != c.want {
			t.Errorf("solve(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestFinalInput(t *testing.T) {
	// Test on the final input
	got1 := solvePart1(convert(helpers.ReadInput("input")))
	want1 := 160
	if got1 != want1 {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got1, want1)
	}

	got2 := solvePart2(convert(helpers.ReadInput("input")))
	want2 := 2271537
	if got2 != want2 {
		t.Errorf("solvePart2(inputFile) == %v, want %v", got2, want2)
	}
}
