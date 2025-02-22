package linklist

// CarPooling 车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）
//
// 给定整数 capacity 和一个数组 trips ,  trip[i] = [numPassengersi, fromi, toi] 表示第 i 次旅行有 numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。
//
// 当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。
func CarPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = 0
	}
	diff := NewDifference(nums)
	for _, trip := range trips {
		// 旅客数量
		userCount := trip[0]
		// 从from开始
		from := trip[1]
		// end结束
		end := trip[2] - 1
		// 从 from 到end 加 userCount个旅客
		diff.increment(from, end, userCount)
	}
	res := diff.getOriginResult()
	for _, re := range res {
		if capacity < re {
			return false
		}
	}
	return true
}

// CalcFlightBookings 计算飞机预定数
// 题目是这样的
// 这里有 n 个航班，它们分别从 1 到 n 进行编号。
//
// 有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
//
// 请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
//
// 示例 1：
//
// 输入：bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5  这个输入的意思是 第 1到第2 航班 预定了 10个座位 第2-第3 航班预定了 20个座位，第2-第5航班预定了25个座位
// 要读懂题目很重要
// 输出：[10,55,45,25,25]
// 解释：
// 航班编号        1   2   3   4   5
// 预订记录 1 ：   10  10
// 预订记录 2 ：       20  20
// 预订记录 3 ：       25  25  25  25
// 总座位数：      10  55  45  25  25
// 因此，answer = [10,55,45,25,25]
// 示例 2：
//
// 输入：bookings = [[1,2,10],[2,2,15]], n = 2
// 输出：[10,25]
// 解释：
// 航班编号        1   2
// 预订记录 1 ：   10  10
// 预订记录 2 ：       15
// 总座位数：      10  25
// 因此，answer = [10,25]
// 最终转换结果为
// 给你输入一个长度为 n 的数组 nums，其中所有元素都是 0。再给你输入一个 bookings，里面是若干三元组 (i, j, k)，每个三元组的含义就是要求你给 nums 数组的闭区间 [i-1,j-1] 中所有元素都加上 k。请你返回最后的 nums 数组是多少？
func CalcFlightBookings(bookings [][]int, needAnswerLength int) []int {

	nums := make([]int, needAnswerLength)
	for i := 0; i < needAnswerLength; i++ {
		nums[i] = 0
	}
	differemce := NewDifference(nums)
	for _, booking := range bookings {
		differemce.increment(booking[0]-1, booking[1]-1, booking[2])
	}
	return differemce.getOriginResult()[:needAnswerLength]
}

// Difference 通过差分数组 得到原始结果
type Difference struct {
	DiffNums []int
}

func NewDifference(nums []int) Difference {
	if len(nums) == 0 {
		panic("nums is nil")
	}
	diffNums := make([]int, len(nums))
	diffNums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diffNums[i] = nums[i] - nums[i-1]
	}
	return Difference{DiffNums: diffNums}
}
func (difference Difference) increment(i, j, val int) {
	difference.DiffNums[i] += val
	if j+1 < len(difference.DiffNums) {
		difference.DiffNums[j+1] -= val
	}
}
func (difference Difference) getOriginResult() []int {
	res := make([]int, len(difference.DiffNums))
	res[0] = difference.DiffNums[0]
	for i := 1; i < len(difference.DiffNums); i++ {
		res[i] = res[i-1] + difference.DiffNums[i]
	}
	return res
}

type NumArray struct {
	Nums     []int
	PreSum   []int
	DiffNums []int
}

func NewNumArray(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	// 差分数组
	diffNums := make([]int, len(nums))
	diffNums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diffNums[i] = nums[i] - nums[i-1]
	}
	return NumArray{Nums: nums, PreSum: preSum, DiffNums: diffNums}
}

// Increment increments a closed interval [i, j] by val (can be negative)
// 给闭区间 [i, j] 增加 val（可以是负数）
// 简直是 精妙
func (numArray NumArray) Increment(i, j, val int) {
	numArray.DiffNums[i] += val
	if j+1 < len(numArray.DiffNums) {
		numArray.DiffNums[j+1] -= val
	}
}
func (numArray NumArray) GetOriginNumsByDiffs() []int {
	res := make([]int, len(numArray.DiffNums))
	res[0] = numArray.DiffNums[0]
	for i := 1; i < len(numArray.DiffNums); i++ {
		res[i] = res[i-1] + numArray.DiffNums[i]
	}
	return res
}
func (numArray NumArray) SumRange(i, j int) int {
	return numArray.PreSum[j+1] - numArray.PreSum[i]
}
