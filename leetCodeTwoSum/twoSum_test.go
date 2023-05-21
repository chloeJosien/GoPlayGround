package leetCodeTwoSum

import (
	"testing"
)

type addTest struct {
	arg1   []int
	arg2   int
	result []int
}

var addTests = []addTest{
	addTest{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	addTest{[]int{3, 2, 4}, 6, []int{1, 2}},
	addTest{[]int{3, 3}, 6, []int{0, 1}},
	addTest{[]int{3, 2}, 6, nil},
}

func TestTwoSum(t *testing.T) {
	for _, test := range addTests {
		if output := twoSum(test.arg1, test.arg2); !MatchingArray(output, test.result) {
			t.Errorf("Output %q not equal to expected %q", output, test.result)
		}
	}
}

func MatchingArray(found []int, expected []int) bool {
	if found == nil && expected == nil {
		return true
	}
	if found == nil && expected != nil {
		return false
	}
	for i := range expected {
		if found[i] != expected[i] {
			return false
		}
	}
	return true
}

// func TestSum_withThree(t *testing.T) {
// 	got := twoSum([]int{1, 2, 3}, 3)
// 	want := []int{0, 1}

// 	if got[0] != want[0] {
// 		t.Errorf("got %q, wanted %q", got[0], want[0])
// 	}
// }
