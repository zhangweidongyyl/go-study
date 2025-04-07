package mypdfcpu

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMergePdf(t *testing.T) {
	convey.Convey("test merge", t, func() {
		//MergePdf("/Users/zyb/code/mycode/go-study/mypdfcpu/test1.pdf",
		//	"/Users/zyb/code/mycode/go-study/mypdfcpu/merged-90.pdf", 22)
		//
		//MergePdf("/Users/zyb/code/mycode/go-study/mypdfcpu/11月分享会-部分内容.pdf",
		//	"/Users/zyb/code/mycode/go-study/mypdfcpu/merged-36.pdf", 22)
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
