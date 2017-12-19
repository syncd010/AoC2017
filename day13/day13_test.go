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
		{in: []string{"0: 3", "1: 2", "4: 4", "6: 4"}, want: 24},
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
	want := 2688
	if got != want {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{in: []string{"0: 3", "1: 2", "4: 4", "6: 4"}, want: 10},
	}

	for _, c := range cases {
		got := solvePart2(convert(c.in))
		if got != c.want {
			t.Errorf("solvePart2(%q) == %v, want %v", c.in, got, c.want)
		}
	}

	got := solvePart2(convert(helpers.ReadInput("input")))
	want := 3876272
	if got != want {
		t.Errorf("solvePart2(inputFile) == %v, want %v", got, want)
	}
}
