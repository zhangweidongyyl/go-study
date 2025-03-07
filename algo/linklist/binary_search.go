package linklist

//	func ShipWithDays(weights []int,target int) int {
//		left :=0
//		right:=0
//		for _, weight := range weights {
//			left = min(left,weight)
//			right += weight
//		}
//	}
//
// 函数定义 根据船的运载能力 计算这堆货物 需要在多少天运完
// 注意这里的限制条件 装载重量不会大于船的运载重量，因此 决定了 weights里的值不会大于shipCap 否则 会有死循环
func getDaysByShipCap(weights []int, shipCap int) int {
	days := 0
	for i := 0; i < len(weights); {
		cap := shipCap
		// 如果船的运载量能 运多堆货物 则 i 需要跳多次 而 days 只需要加 1
		// 如果 运载量大于某堆
		for i < len(weights) {

			if cap < weights[i] {
				break
			} else {
				cap = cap - weights[i]
			}
			i++
		}
		days++
	}

	return days
}

// MinEatSpeed 珂珂吃香蕉，求最小速度
func MinEatSpeed(piles []int, target int) int {
	left := 0
	right := 100000000 + 1
	for left < right {
		mid := left + (right-left)/2
		midVal := eatNeedHours(piles, mid)
		// 注意 eatNeedHours函数是单调递减的，注意收缩方向
		if midVal == int64(target) {
			right = mid
		} else if midVal > int64(target) {
			left = mid + 1
		} else if midVal < int64(target) {
			right = mid
		}

	}
	return left
}

// 计算
// 以speed 吃掉 香蕉堆 piles 需要的小时数

func eatNeedHours(piles []int, speed int) int64 {
	hours := int64(0)
	for i := 0; i < len(piles); i++ {
		hours += int64(piles[i]) / int64(speed)
		if int64(piles[i])%int64(speed) > 0 {
			hours++
		}
	}
	return hours
}

// LeftBound 找出左侧边界
// 思想：一直往左侧逼近
func LeftBound(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		midVal := nums[mid]
		if midVal == target {
			// 此时 mid 不一定是最左侧的，还需继续往左
			right = mid - 1
		} else if midVal < target {
			left = mid + 1
		} else if midVal > target {
			right = mid - 1
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}

// RightBound 找出右侧边界
// 思想：一直往右侧逼近
func RightBound(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		midVal := nums[mid]
		if midVal == target {
			// 此时 mid 不一定是最右侧的，还需继续往you
			left = mid + 1
		} else if midVal < target {
			left = mid + 1
		} else if midVal > target {
			right = mid - 1
		}
	}
	if nums[right] == target {
		return right
	}
	return -1
}
func BinarySearch(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	res := make([]int, 0)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		midVal := nums[mid]
		if midVal == target {
			res = append(res, mid)
			left = mid + 1
		} else if midVal < target {
			left = mid + 1
		} else if midVal > target {
			right = mid - 1
		}
	}
	return res
}
