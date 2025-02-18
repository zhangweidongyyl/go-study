package linklist

import "math"

// FindLongestOnes 经过k次翻转，能得到最长的1的个数
func FindLongestOnes(nums []int, k int) int {
	windowOfOneCount := 0
	left, right := 0, 0
	maxLen := 0
	for right < len(nums) {
		if nums[right] == 1 {
			windowOfOneCount++
		}
		right++

		// 收缩窗口的条件
		if right-left-windowOfOneCount > k {
			if nums[left] == 1 {
				windowOfOneCount--
			}
			left++
		}
		maxLen = max(maxLen, right-left)
	}
	return maxLen
}

// FindProductLessThanK 给定一个数组nums，找出里面元素乘积小余K的组合个数
func FindProductLessThanK(nums []int, k int) int {
	left, right := 0, 0
	windowProduct := 1
	resCount := 0
	for right < len(nums) {
		num := nums[right]
		windowProduct *= num
		right++

		// 判断窗口收缩条件
		if windowProduct >= k && right > left {
			leftNum := nums[left]
			windowProduct /= leftNum
			left++
		}
		// 这里为啥是right-left?是因为一个大数组满足，那么其子的组合肯定满足
		resCount += right - left

	}
	return resCount
}

// FindMinOperations 给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要 修改 数组以供接下来的操作使用。
// 注意读题，从边缘删除，那么意思就是剩余的尽可能大，并且是连续的
// 如果可以将 x 恰好 减到 0 ，返回 最小操作数 ；否则，返回 -1 。
// 本问题可以翻译下  数组里  有几个数加起来能 等于x
func FindMinOperations(nums []int, x int) int {
	sums := 0
	for _, num := range nums {
		sums += num
	}
	windowSum := 0
	// 记录 有哪些元素在window的
	window := make(map[int]bool, 0)
	target := sums - x
	left, right := 0, 0
	res := make([]int, 0)
	for right < len(nums) {
		windowSum += nums[right]
		window[right] = true
		right++

		for windowSum > target && right > left {
			if val, ok := window[left]; ok && val {
				window[left] = false
			}
			windowSum = windowSum - nums[left]
			left++
		}
		if windowSum == target {
			for k, v := range window {
				if v == true {
					res = append(res, k)
				}
			}
			return len(nums) - len(res)
		}

	}
	return -1
}
func FindLongestSubString(source string) string {
	windowMap := make(map[rune]int, 0)
	left, right := 0, 0
	res := 0
	for right < len(source) {
		c := rune(source[right])
		windowMap[c]++
		right++

		// 这个时候 窗口里有和c重复的字符了 进行收缩窗口
		for windowMap[c] > 1 {
			// 此时要进行 窗口收缩
			d := rune(source[left])
			left++

			windowMap[d]--
		}
		if res < right-left {
			res = right - left
		}

	}

	return source[left:right]
}
func FindStringPosition(source, includeString string) []int {
	needMap := make(map[rune]int, 0)
	for _, s := range includeString {
		needMap[s]++
	}
	windowMap := make(map[rune]int, 0)
	left, right := 0, 0
	valid := 0

	res := make([]int, 0)
	for right < len(source) {
		c := rune(source[right])
		right++

		if _, ok := needMap[c]; ok {
			windowMap[c]++
			if windowMap[c] == needMap[c] {
				valid++
			}
		}

		if right-left >= len(includeString) {
			if valid == len(needMap) {
				res = append(res, left)
			}

			d := rune(source[left])
			left++
			if _, ok := windowMap[d]; ok {
				if windowMap[d] == needMap[d] {
					valid--
				}
				windowMap[d]--
			}
		}
	}
	return res
}

// CheckContainsString check source是否包含includeString
func CheckContainsString(source, includeString string) bool {
	needMap := make(map[rune]int, 0)
	for _, s := range includeString {
		needMap[s]++
	}
	windowMap := make(map[rune]int, 0)
	left, right := 0, 0
	// 判断windowMap是否有needMap一致字符的标志位
	valid := 0
	for right < len(source) {
		c := rune(source[right])
		right++
		if _, ok := needMap[c]; ok {
			windowMap[c]++
			if windowMap[c] == needMap[c] {
				valid++
			}
		}

		if right-left >= len(includeString) {
			if valid == len(needMap) {
				return true
			}

			// 先取出来，再位移，否则有问题
			d := rune(source[left])
			left++
			if _, ok := windowMap[d]; ok {

				if windowMap[d] == needMap[d] {
					valid--
				}

				windowMap[d]--
			}
		}

	}

	return false
}
func MinString(source, includeString string) string {
	needMap := make(map[byte]int, 0)
	for i := 0; i < len(includeString); i++ {
		needMap[includeString[i]]++
	}

	windowMap := make(map[byte]int, 0)
	left, right := 0, 0
	// 判断 字符出现的次数，如果在两个map出现一样多则 ++
	valid := 0

	// 记录最小覆盖子串的起始索引及长度
	start, length := 0, math.MaxInt32
	for right < len(source) {
		c := source[right]
		right++
		// c 是要判断的字符
		if _, ok := needMap[c]; ok {
			windowMap[c]++
			if windowMap[c] == needMap[c] {
				valid++
			}
		}
		// 此时 windowMap里出现了 needMap里的所有字符，且都是一次
		// 但不表示 此时windowMap里是最短 的包含includeString的串，可能 比需要的结果长
		for valid == len(needMap) {
			// 这里的逻辑是 处理最短的

			if right-left < length {
				start = left
				length = right - left
			}

			// 要尝试移出的左侧
			d := source[left]
			// d 移出后 不满足 for循环条件了，此时就是left的上一位是最短串，即起始位是start
			left++
			if _, ok := needMap[d]; ok {
				if windowMap[d] == needMap[d] {
					valid--
				}
				windowMap[d]--
			}
		}

	}
	if length == math.MaxInt32 {
		return ""
	}
	return source[start : start+length]

}
