package mypdfcpu

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMergePdf(t *testing.T) {
	convey.Convey("test merge", t, func() {
		MergePdf()
	})
}
