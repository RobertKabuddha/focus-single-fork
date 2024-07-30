package lee_code

import (
	"fmt"
	"github.com/gogf/gf/v2/test/gtest"
	"sort"
	"testing"
)

//给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//
//请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
//注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。
//为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

func TestArrMerge(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		//t.Assert(minimumOperations("200"), 0)
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2, 5, 6}
		numsRes := []int{1, 2, 2, 3, 5, 6}
		t.Assert(merge(nums1, 6, nums2, 3), numsRes)
	})
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	if m == 0 {
		return nil
	}
	nums3 := make([]int, 0, m+n)
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != 0 {
			nums3 = append(nums3, nums1[i])
		}
	}
	for i := 0; i < len(nums2); i++ {
		if nums2[i] != 0 {
			nums3 = append(nums3, nums2[i])
		}
	}

	//排序
	sort.Ints(nums3)
	nums1 = nums3[0:m] //赋值是无效的
	fmt.Println(nums1)
	return nums1[0:m]
}

func mergeCurrent(nums1 []int, m int, nums2 []int, n int) {
	// 将 nums2 中的元素复制到 nums1 的末尾
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	// 排序 nums1
	sort.Ints(nums1)
}
