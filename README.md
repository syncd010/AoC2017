# Advent of Code 2017, Go

These are my solutions for [Advent of Code 2017](http://adventofcode.com/2017), in [Go](https://golang.org/).

Not much to see here. This code isn't meant to be elegant, efficient, well documented or whatever, it's just a way to try out Go and have some fun.

Notes:
- This worked with *Go*'s version circa 2017. *Go* is still evolving, so this might rot work in the future;
- This repo relies on `$GOPATH` being set appropriately, it doesn't use go modules.

## Structure and usage
Each day's puzzle is contained in a separate directory, with a separate *go* program, associated tests and the day's input (which is the file name `input`).

To run: 
```
cd day{n}
go run day{n}.go {FILE}
go test -v
```
where FILE can be the `input` file on each directory.
