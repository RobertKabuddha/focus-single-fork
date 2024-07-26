package lee_code

import (
	"fmt"
	"github.com/gogf/gf/v2/test/gtest"
	"strconv"
	"testing"
)

// 1. 创建测试文件
// 测试文件通常以 _test.go 结尾。创建一个新的测试文件，例如 main_test.go。
// 2. 编写测试函数
// 测试函数以 Test 开头，并接收一个 *testing.T 参数。

func TestSpecialNum(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		//t.Assert(minimumOperations("200"), 0)
		t.Assert(minimumOperations("251"), 1)
	})
}

func minimumOperations(num string) int {
	// 删除一个数字 看是否能被25整除 能被整除返回 不能继续删
	// 返回删除的数字个数
	orgS := make([]string, 1)
	orgS = append(orgS, num)
	tier := getTier(orgS, 0)
	fmt.Println(tier)
	return tier

}

func getTier(num []string, tier int) int {
	if canDivisibleList(num) {
		return tier
	}
	possibleSlice := make([]string, 0)
	// 尝试所有可能的子序列，检查它们是否能被 25 整除
	for j := 0; j < len(num); j++ {
		numj := num[j]
		for i := 0; i < len(numj); i++ {
			s := numj[:i] + numj[i+1:]
			possibleSlice = append(possibleSlice, s)
		}
	}
	return getTier(possibleSlice, tier+1)
	return -1
}

func canDivisibleList(nums []string) bool {
	for i := 0; i < len(nums); i++ {
		nums[i] = removeLeadingZeros(nums[i])
		if nums[i] != "" {
			//Atoi 等效于 ParseInt（s， 10， 0），转换为 int 类型。
			val, err := strconv.Atoi(nums[i])
			if err == nil {
				if val%25 == 0 {
					return true
				}
			}
		}
	}
	return false
}

func canDivisible(num string) bool {
	num = removeLeadingZeros(num)
	if num != "" {
		//Atoi 等效于 ParseInt（s， 10， 0），转换为 int 类型。
		val, err := strconv.Atoi(num)
		if err == nil {
			if val%25 == 0 {
				return true
			}
		}
	}
	return false
}

func removeLeadingZeros(s string) string {
	for len(s) > 0 && s[0] == '0' {
		s = s[1:]
	}
	if s == "" {
		return "0"
	}
	return s
}
