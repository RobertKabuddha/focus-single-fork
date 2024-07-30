package lee_code

import (
	"github.com/gogf/gf/v2/test/gtest"
	"sort"
	"testing"
)

//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。
//
//假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：
//更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
//返回 k。

func TestRemoveElement(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		//t.Assert(minimumOperations("200"), 0)
		nums1 := []int{1, 2, 3, 0, 0, 0}
		t.Assert(removeElement(nums1, 3), 2)
	})
}

func removeElement(nums []int, val int) int {
	size := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i] = 0
		} else {
			size++
		}
	}
	//倒叙
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return size
}
