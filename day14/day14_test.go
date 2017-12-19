package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{in: []string{"flqrgnkx"}, want: 8108},
		{in: []string{"uugsqrei"}, want: 8194},
	}

	for _, c := range cases {
		got := solvePart1(c.in[0])
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
		{in: []string{"flqrgnkx"}, want: 1242},
		{in: []string{"uugsqrei"}, want: 1141},
	}

	for _, c := range cases {
		got := solvePart2(c.in[0])
		if got != c.want {
			t.Errorf("solvePart2(%q) == %v, want %v", c.in, got, c.want)
		}
	}

	// got := solvePart2(convert(helpers.ReadInput("input")))
	// want := 3876272
	// if got != want {
	// 	t.Errorf("solvePart2(inputFile) == %v, want %v", got, want)
	// }
}
