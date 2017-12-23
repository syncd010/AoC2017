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
		{in: []string{
			"p=<1376,2307,-2194>, v=<61,-86,-48>, a=<10,1,12>",
			"p=<-1982,30,-2355>, v=<15,-83,19>, a=<6,7,7>",
		}, want: 1},
		{in: []string{
			"p=<1376,2307,-2194>, v=<0,1,0>, a=<0,0,0>",
			"p=<-1982,30,-2355>, v=<1,1,0>, a=<2,0,0>",
			"p=<-1982,30,-2355>, v=<0,0,0>, a=<1,-1,0>",
		}, want: 0},
		{in: []string{
			"p=<0,1,1>, v=<0,1,0>, a=<0,-1,0>",
			"p=<0,0,0>, v=<1,1,0>, a=<1,1,0>",
			"p=<-1,0,0>, v=<0,1,0>, a=<1,2,0>",
			"p=<0,30,0>, v=<0,0,0>, a=<1,-1,0>",
		}, want: 0},
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
			"p=<1376,2307,-2194>, v=<61,-86,-48>, a=<-10,-1,12>",
			"p=<-1982,30,-2355>, v=<15,-83,19>, a=<6,7,7>",
		}, want: 0},
		{in: []string{
			"p=<-6,0,0>, v=<3,0,0>, a=<0,0,0>",
			"p=<-4,0,-0>, v=<2,0,0>, a=<0,0,0>",
			"p=<-2,0,0>, v=<1,0,0>, a=<0,0,0>",
			"p=<3,0,0>, v=<-1,0,0>, a=<0,0,0>",
		}, want: 1},
		{in: []string{
			"p=<2,3,0>, v=<-1,0,0>, a=<0,-1,0>",
			"p=<0,0,0>, v=<0,0,0>, a=<0,0,0>",
			"p=<0,0,0>, v=<2,0,0>, a=<0,0,0>",
			"p=<2,0,0>, v=<0,0,0>, a=<0,0,0>",
		}, want: 0},
	}

	for _, c := range cases {
		got := solvePart2(convert(c.in))
		if got != c.want {
			t.Errorf("solvePart2(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestFinalInput(t *testing.T) {
	// Test on the final input
	got1 := solvePart1(convert(helpers.ReadInput("input")))
	want1 := 150
	if got1 != want1 {
		t.Errorf("solvePart1(inputFile) == %v, want %v", got1, want1)
	}

	got2 := solvePart2(convert(helpers.ReadInput("input")))
	want2 := 657
	if got2 != want2 {
		t.Errorf("solvePart2(inputFile) == %v, want %v", got2, want2)
	}
}
