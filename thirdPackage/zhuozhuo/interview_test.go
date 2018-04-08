package zhuozhuo

import (
	"testing"
)

func TestGetLongestSubstrings(t *testing.T) {
	tests := []struct {
		str string
	}{
		{
			"aabcceddabc",
		},
		{
			"abcaebee",
		},
	}
	for _, test := range tests {
		result, num := getLongestSubstrings(test.str)
		t.Errorf("merge :%v %v \n", result, num)
	}
}
func TestMerge(t *testing.T) {
	tests := []struct {
		nums1 []int // 单元测试要规定nums1的cap
		nums2 []int
	}{
		{
			[]int{1, 8}, []int{3, 7},
		},
		{
			[]int{1, 3, 4}, []int{1, 1, 2, 5},
		},
	}
	for _, test := range tests {
		result := merge(test.nums1, test.nums2)
		t.Errorf("merge :%v \n", result)
	}
}

func TestJumpGame(t *testing.T) {
	var tests = []struct {
		nums []int
	}{
		{[]int{2, 3, 1, 1, 4}},
		{[]int{3, 2, 1, 0, 4}},
	}
	for _, test := range tests {
		t.Errorf("canJump result :%v \n", canJump(test.nums))
	}
}
