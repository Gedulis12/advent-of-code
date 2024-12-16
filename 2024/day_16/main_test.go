package main

import (
	"testing"
)

func TestE1(t *testing.T) {
	input := "example1"
	want1, want2 := 7036, 45
	minimum, locs := SolvePart1And2(input)
	if want1 != minimum || want2 != locs {
		t.Fatalf("Expected: %d,%d Got: %d,%d ", want1, want2, minimum, locs)
	}
	t.Logf("Expected: %d,%d Got: %d,%d ", want1, want2, minimum, locs)
}

func TestE2(t *testing.T) {
	input := "example2"
	want1, want2 := 11048, 64
	minimum, locs := SolvePart1And2(input)
	if want1 != minimum || want2 != locs {
		t.Fatalf("Expected: %d,%d Got: %d,%d ", want1, want2, minimum, locs)
	}
	t.Logf("Expected: %d,%d Got: %d,%d ", want1, want2, minimum, locs)
}