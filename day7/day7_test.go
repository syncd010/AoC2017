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
			"pbga (66)",
			"xhth (57)",
			"ebii (61)",
			"havc (66)",
			"ktlj (57)",
			"fwft (72) -> ktlj, cntj, xhth",
			"qoyq (66)",
			"padx (45) -> pbga, havc, qoyq",
			"tknk (41) -> ugml, padx, fwft",
			"jptl (61)",
			"ugml (68) -> gyxo, ebii, jptl",
			"gyxo (61)",
			"cntj (57)",
		}, want: "tknk"},
	}

	for _, c := range cases {
		got := validate(c.in)
		if got != nil {
			t.Errorf("validate(%q) == %v, want nil", c.in, got)
		}
	}

	for _, c := range cases {
		got := solvePart1(convert(c.in))
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}

	got := solvePart1(convert(helpers.ReadInput("input")))
	want := "rqwgj"
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
			"pbga (66)",
			"xhth (57)",
			"ebii (61)",
			"havc (66)",
			"ktlj (57)",
			"fwft (72) -> ktlj, cntj, xhth",
			"qoyq (66)",
			"padx (45) -> pbga, havc, qoyq",
			"tknk (41) -> ugml, padx, fwft",
			"jptl (61)",
			"ugml (68) -> gyxo, ebii, jptl",
			"gyxo (61)",
			"cntj (57)",
		}, want: 60},
	}

	for _, c := range cases {
		got := solvePart2(convert(c.in))
		if got != c.want {
			t.Errorf("solvePart1(%q) == %v, want %v", c.in, got, c.want)
		}
	}

	got := solvePart2(convert(helpers.ReadInput("input")))
	want := 333
	if got != want {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got, want)
	}
}
