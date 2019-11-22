package main

import (
	"testing"
)

func TestMagicSquareSample1(t *testing.T) {
	sampleData := [][]int32{
		[]int32{5, 3, 4}, []int32{1, 5, 8}, []int32{6, 4, 2},
	}

	cost := formingMagicSquare(sampleData)

	if cost != 7 {
		t.Errorf("got %d instead of 7", cost)
	}
}

func TestMagicSquareSample2(t *testing.T) {
	sampleData := [][]int32{
		[]int32{4, 9, 2}, []int32{3, 5, 7}, []int32{8, 1, 5},
	}

	cost := formingMagicSquare(sampleData)

	if cost != 1 {
		t.Errorf("got %d instead of 1", cost)
	}
}

func TestMagicSquareSample3(t *testing.T) {
	sampleData := [][]int32{
		[]int32{4, 8, 2}, []int32{4, 5, 7}, []int32{6, 1, 6},
	}

	cost := formingMagicSquare(sampleData)

	if cost != 4 {
		t.Errorf("got %d instead of 4", cost)
	}
}
