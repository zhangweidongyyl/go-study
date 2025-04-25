package mypdfcpu

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestGoMonkey(t *testing.T) {
	convey.Convey("TestGoMonkey", t, func() {
		patches := gomonkey.NewPatches()
		defer patches.Reset()

		// Mock 函数 C 的行为
		patches.ApplyFunc(MergeC, func() string {
			return "failC"
		})
		t.Logf("c is %s", MergeC())
		// 验证 testC() 是否已被 Mock
		convey.Convey("Verify testC is mocked", func() {
			convey.So(MergeC(), convey.ShouldEqual, "failC") // 先确认 Mock 生效
		})
		res := MergeA()
		t.Logf("res is %s", res)
	})
}

func TestMergePdf(t *testing.T) {
	convey.Convey("test merge", t, func() {

		dataCompleteProgress := fmt.Sprintf("%.f", math.Ceil(float64(33*100)/float64(100)))
		t.Logf("%s", (dataCompleteProgress + "%"))
		MergePdf("/Users/zyb/code/mycode/go-study/mypdfcpu/test1.pdf",
			"/Users/zyb/code/mycode/go-study/mypdfcpu/merged-90.pdf", 22)

		MergePdf("/Users/zyb/code/mycode/go-study/mypdfcpu/11月分享会-部分内容.pdf",
			"/Users/zyb/code/mycode/go-study/mypdfcpu/merged-36.pdf", 22)
		i := 0
		for {
			if i >= 10 {
				break
			}
			MergePdf("/Users/zyb/code/mycode/go-study/mypdfcpu/第 1 讲抽象概括能力提升（一） (1).pdf",
				"/Users/zyb/code/mycode/go-study/mypdfcpu/merged-1.pdf", 22)
			i++
		}

	})
}
