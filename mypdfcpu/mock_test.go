package mypdfcpu

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	convey.Convey("TestGoMonkey1", t, func() {
		//Example()
		//GetWidth()

		fmt.Println(strconv.FormatInt(130735%20, 10))
	})

}
func testA() string {
	return fmt.Sprintf("%s-%s", "testA", testB())
}

func testB() string {
	return testC() + ":realB"
}

func testC() string {
	return "realC"
}

func TestGoMonkey1(t *testing.T) {
	convey.Convey("TestGoMonkey1", t, func() {
		patches := gomonkey.NewPatches()
		defer patches.Reset()
		patches.ApplyFunc(testC, func() string {
			return "failC"
		})

		res := testA()
		t.Logf("res: %s", res)
		convey.So(res, convey.ShouldEqual, "testA-failC:realB")
	})
}
