package lfu

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLFUCache_Put(t *testing.T) {
	convey.Convey("lfuCache put", t, func() {
		lfuCache := NewLFUCache(3)
		lfuCache.Put(1, 1)
		lfuCache.Put(2, 2)
		lfuCache.Put(3, 3)
		lfuCache.Put(4, 4)
		lfuCache.Put(5, 5)
		lfuCache.Put(6, 6)
		lfuCache.Put(7, 7)
		lfuCache.Put(8, 8)
		lfuCache.Put(9, 9)
		lfuCache.Put(10, 10)
		lfuCache.Put(11, 11)

		lfuCache.Get(9)
		lfuCache.Get(10)
		lfuCache.Get(11)
	})
}
