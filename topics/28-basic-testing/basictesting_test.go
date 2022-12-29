package basictesting

import (
	"testing"
)

/*
TEST FUNCTION MUST START WITH 'Test' WORD
*/

func TestSum(t *testing.T) {
	total := Sum(5, 8)

	if total != 13 {
		t.Errorf("Error! Received: %d - Expected: %d", total, 13)
	}
}

func TestSumTable(t *testing.T) {
	table := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{15, 10, 25},
		{48, 50, 98},
		{-5, 4, -1},
	}

	for _, item := range table {
		total := Sum(item.a, item.b)
		if total != item.c {
			t.Errorf("Error! Received: %d - Expected: %d", total, item.c)
		}
	}
}

func TestGetmax(t *testing.T) {
	table := []struct {
		a int
		b int
		c int
	}{
		{10, 20, 20},
		{-5, 10, 10},
		{48, -3, 48},
		{-500, 4, 4},
	}

	for _, item := range table {
		res := GetMax(item.a, item.b)
		if res != item.c {
			t.Errorf("Error! Received: %d - Expected: %d", res, item.c)
		}
	}
}

func TestFibonacci(t *testing.T) {
	table := []struct {
		n int
		r int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{8, 21},
		{45, 1134903170},
	}

	for _, item := range table {
		res := Fibonacci(item.n)
		if res != item.r {
			t.Errorf("Error! Received: %d - Expected: %d", res, item.r)
		}
	}

}
