package main

import (
	"testing"
)

func TestPart1E1(t *testing.T) {
	input := "example1"
	want := 7036
	res := SolvePart1(input)
	if want != res {
		t.Fatalf("Expected: %d, Got: %d ", want, res)
	}
	t.Logf("Expected: %d, Got: %d ", want, res)
}

func TestPart1E2(t *testing.T) {
	input := "example2"
	want := 11048
	res := SolvePart1(input)
	if want != res {
		t.Fatalf("Expected: %d, Got: %d ", want, res)
	}
	t.Logf("Expected: %d, Got: %d ", want, res)
}
