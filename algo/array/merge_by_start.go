package array

import "sort"

// Interval 定义区间结构
type Interval struct {
	Start int
	End   int
}

// 实现 sort.Interface 接口，用于排序
type ByStart []Interval

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].Start < a[j].Start }

// MergeIntervals 合并重叠区间
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}

	// 按区间的起始点排序
	sort.Sort(ByStart(intervals))

	merged := []Interval{intervals[0]} // 初始化结果列表
	for i := 1; i < len(intervals); i++ {
		last := &merged[len(merged)-1] // 获取结果列表中的最后一个区间
		current := intervals[i]        // 当前区间

		// 如果当前区间与最后一个区间有重叠
		if current.Start <= last.End {
			// 合并区间，取最大的结束点
			if current.End > last.End {
				last.End = current.End
			}
		} else {
			// 否则，将当前区间加入结果列表
			merged = append(merged, current)
		}
	}

	return merged
}

type SortRange struct {
	Start int
	End   int
}
type SortByStart []SortRange

func (this SortByStart) Len() int {
	return len(this)
}
func (this SortByStart) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this SortByStart) Less(i, j int) bool {
	return this[i].Start < this[j].Start
}
func MergeSortRange(datas []SortRange) []SortRange {
	mergedVals := make([]SortRange, 0)
	sort.Sort(SortByStart(datas))
	mergedVals = append(mergedVals, datas[0])
	for i := 1; i < len(datas); i++ {
		// 取栈顶  这个last必须取地址，否则 last修改了 并不会修改mergedVals里面的值  ，注意
		last := &mergedVals[len(mergedVals)-1]
		current := datas[i]
		// 有重合的
		if current.Start < last.End {
			if current.End > last.End {
				last.End = current.End
			}
		} else {
			// 非重合
			mergedVals = append(mergedVals, current)
		}
	}
	return mergedVals
}
