package main

import (
    "testing"
)

func TestPart1(t *testing.T) {
    input := "example"
    want := 6440
    res := SolvePart1(input)
    if want != res {
        t.Fatalf("Expected: %d, Got: %d ", want, res)
    }
}

func TestPart2(t *testing.T) {
    input := "example"
    want := 5905
    res := SolvePart2(input)
    if want != res {
        t.Fatalf("Expected: %d, Got: %d ", want, res)
    }
}
