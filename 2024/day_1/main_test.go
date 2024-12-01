package main

import (
    "testing"
)

func TestPart1(t *testing.T) {
    input := "example"
    want := 11
    res := SolvePart1(input)
    if want != res {
        t.Fatalf("Expected: %d, Got: %d ", want, res)
    }
	t.Logf("Expected: %d, Got: %d ", want, res)
}

func TestPart2(t *testing.T) {
    input := "example"
    want := 31
    res := SolvePart2(input)
    if want != res {
        t.Fatalf("Expected: %d, Got: %d ", want, res)
    }
	t.Logf("Expected: %d, Got: %d ", want, res)
}
