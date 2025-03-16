package array

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	convey.Convey("testMergeIntervals", t, func() {
		intervals := []Interval{{1, 2}, {2, 5}, {6, 10}}
		result := MergeIntervals(intervals)
		exceptResult := []Interval{{1, 5}, {6, 10}}
		convey.So(result, convey.ShouldEqual, exceptResult)
	})
}

func TestMergeSortRange(t *testing.T) {
	convey.Convey("testMergeSortRange", t, func() {
		sortRanges := []SortRange{{1, 2}, {2, 5}, {6, 10}}
		result := MergeSortRange(sortRanges)
		exceptResult := []Interval{{1, 5}, {6, 10}}
		convey.So(result, convey.ShouldEqual, exceptResult)
	})
}
