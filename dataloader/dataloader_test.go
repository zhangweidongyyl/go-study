package dataloader

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestFrom(t *testing.T) {
	convey.Convey("t", t, func() {
		Do()
	})
}
func Do() {
	var items []int
	func() (val int) {
		time.Sleep(5 * time.Second)
		items = append(items, 11)
		return 11
	}()
	fmt.Printf("items is %+v \r\n", items)
}
