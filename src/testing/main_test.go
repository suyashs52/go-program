package main

import (
	"testing"

	"src.com/user/testing/acdc"
)

func TestMySum(t *testing.T) {
	//go test to run the test
	//go test -v
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{4, 6}, 10},
		test{[]int{21, 21}, 42},
		test{[]int{1, 1}, 2},
		test{[]int{4, -1}, 3},
	}

	for _, v := range tests {
		x := acdc.Sum(v.data...)

		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}
