package main

import (
	"fmt"
	"sort"
)

// 1.找出只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
// 1.1 利用异或运算遍历nums数组
func singleNumber1(nums []int) int {
	single := 0
	for _, n := range nums {
		single ^= n
	}
	return single
}

// 1.2 利用加法器记录每个元素出现的次数，再遍历map找到出现次数为1的数
func singleNumber2(nums []int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	for k, cnt := range freq {
		if cnt == 1 {
			return k
		}
	}
	return -1
}

// 2.回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reveredNumber := 0
	for x > reveredNumber {
		reveredNumber = reveredNumber*10 + x/10
		x /= 10
	}

	return x == reveredNumber || x == reveredNumber/10

}

// 3.有效的括号
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// 4.最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 5.加一
// 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
// 将大整数加 1，并返回结果的数字数组
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

// 6.删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	i := 1
	for j := 1; j < n; j++ {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

// 7.合并区间
func merge(intervals [][]int) [][]int {

	// Interval 仅用于演示，可用 [][]int 直接代替
	type Interval []int

	// 1. 按区间左端点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var merged [][]int
	for _, cur := range intervals {
		// 2. 当前区间与上一区间不重叠，直接追加
		if len(merged) == 0 || merged[len(merged)-1][1] < cur[0] {
			merged = append(merged, []int{cur[0], cur[1]})
		} else {
			// 3. 否则合并：取右端点的最大值
			last := &merged[len(merged)-1]
			if cur[1] > (*last)[1] {
				(*last)[1] = cur[1]
			}
		}
	}
	return merged
}

// 8.两数之和
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil

}

func main() {
	fmt.Println("-------------1. 找出只出现一次的数字------------------")
	fmt.Println("-------------利用异或运算----------------")
	fmt.Println(singleNumber1([]int{1, 2, 2, 3, 3})) // 1
	fmt.Println("-------------利用哈希表----------------")
	fmt.Println(singleNumber2([]int{4, 1, 2, 1, 2})) // 4

	fmt.Println("-------------2. 回文数------------------")
	fmt.Println(isPalindrome(121)) // true
	fmt.Println(isPalindrome(10))  // false

	fmt.Println("-------------3. 有效的括号------------------")
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true

	fmt.Println("-------------4. 最长公共前缀------------------")
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""

	fmt.Println("-------------5. 加一------------------")
	fmt.Println(plusOne([]int{1, 2, 3})) // [1 2 4]
	fmt.Println(plusOne([]int{9, 9, 9})) // [1 0 0 0]

	fmt.Println("-------------6. 删除有序数组中的重复项------------------")
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Println("长度:", k)           // 5
	fmt.Println("去重后数组:", nums[:k]) // [0 1 2 3 4]

	fmt.Println("-------------7. 合并区间------------------")
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1 6] [8 10] [15 18]]
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))                    // [[1 5]]

	fmt.Println("-------------8. 两数之和------------------")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1 2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0 1]
}
