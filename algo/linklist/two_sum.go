package linklist

func ReverseString(str []rune) string {
	left, right := 0, len(str)-1
	for left < right {
		t := str[left]
		str[left] = str[right]
		str[right] = t
		left++
		right--
	}
	return string(str)
}

// TwoSum 两数之和  nums是一个非递减的数组，查找出nums里两个数的和为target的索引，索引从1 开始
func TwoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}
