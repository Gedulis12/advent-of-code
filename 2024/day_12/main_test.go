package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := "example"
	want := 1930
	res := SolvePart1(input)
	if want != res {
		t.Fatalf("Expected: %d, Got: %d ", want, res)
	}
	t.Logf("Expected: %d, Got: %d ", want, res)
}

func TestPart2(t *testing.T) {
	input := "example"
	want := 1206
	res := SolvePart2(input)
	if want != res {
		t.Fatalf("Expected: %d, Got: %d ", want, res)
	}
	t.Logf("Expected: %d, Got: %d ", want, res)
}

func TestPart2E1(t *testing.T) {
	input := "tests/ex1"
	want := 80
	res := SolvePart2(input)
	if want != res {
		t.Fatalf("Expected: %d, Got: %d ", want, res)
	}
	t.Logf("Expected: %d, Got: %d ", want, res)
}

func TestPart2E2(t *testing.T) {
	input := "tests/ex2"
	want := 436
	res := SolvePart2(input)
	if want != res {
		t.Fatalf("Expected: %d, Got: %d ", want, res)
	}
	t.Logf("Expected: %d, Got: %d ", want, res)
}

func TestPart2E3(t *testing.T) {
	input := "tests/ex3"
	want := 236
	res := SolvePart2(input)
	if want != res {
		t.Fatalf("Expected: %d, Got: %d ", want, res)
	}
	t.Logf("Expected: %d, Got: %d ", want, res)
}

func TestPart2E4(t *testing.T) {
	input := "tests/ex4"
	want := 368
	res := SolvePart2(input)
	if want != res {
		t.Fatalf("Expected: %d, Got: %d ", want, res)
	}
	t.Logf("Expected: %d, Got: %d ", want, res)
}

func TestPart2E5(t *testing.T) {
	input := "tests/ex5"
	want := 572
	res := SolvePart2(input)
	if want != res {
		t.Fatalf("Expected: %d, Got: %d ", want, res)
	}
	t.Logf("Expected: %d, Got: %d ", want, res)
}

func TestPart2E6(t *testing.T) {
	input := "tests/ex6"
	want := 572
	res := SolvePart2(input)
	if want != res {
		t.Fatalf("Expected: %d, Got: %d ", want, res)
	}
	t.Logf("Expected: %d, Got: %d ", want, res)
}