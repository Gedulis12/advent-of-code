package main

import (
        "testing"
)

func TestPart1(t *testing.T) {
        input := "example"
        want := "5,7,3,0"
        res := SolvePart1(input)
        if want != res {
                t.Fatalf("Expected: %s, Got: %s ", want, res)
        }
        t.Logf("Expected: %s, Got: %s ", want, res)
}

func TestPart2(t *testing.T) {
        input := "example"
        want := 19050
        res := SolvePart2(input)
        if want != res {
                t.Fatalf("Expected: %d, Got: %d ", want, res)
        }
        t.Logf("Expected: %d, Got: %d ", want, res)
}