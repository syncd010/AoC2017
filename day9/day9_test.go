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
		{in: []string{"{{{}}}"}, want: 6},
		{in: []string{"{{},{}}"}, want: 5},
		{in: []string{"{{{},{},{{}}}}"}, want: 16},
		{in: []string{"{<a>,<a>,<a>,<a>}"}, want: 1},
		{in: []string{"{{<ab>},{<ab>},{<ab>},{<ab>}}"}, want: 9},
		{in: []string{"{{<!!>},{<!!>},{<!!>},{<!!>}}"}, want: 9},
		{in: []string{"{{<a!>},{<a!>},{<a!>},{<ab>}}"}, want: 3},
	}

	for _, c := range cases {
		got := solvePart1(convert(c.in)[0])
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
		{in: []string{"<random characters>"}, want: 17},
		{in: []string{"<<<<>"}, want: 3},
		{in: []string{"<{!>}>"}, want: 2},
		{in: []string{"<!!>"}, want: 0},
		{in: []string{"<!!!>>"}, want: 0},
		{in: []string{"<{o\"i!a,<{i<a>,"}, want: 10},
	}

	for _, c := range cases {
		got := solvePart2(convert(c.in)[0])
		if got != c.want {
			t.Errorf("solvePart2(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestFinalInput(t *testing.T) {
	// Test on the final input
	got1 := solvePart1(convert(helpers.ReadInput("input"))[0])
	want1 := 9662
	if got1 != want1 {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got1, want1)
	}

	got2 := solvePart2(convert(helpers.ReadInput("input"))[0])
	want2 := 4903
	if got2 != want2 {
		t.Errorf("solvePart2(inputFile) == %v, want %v", got2, want2)
	}
}
