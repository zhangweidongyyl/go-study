package search_ip

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTrieSet_Insert(t *testing.T) {
	convey.Convey("test insert", t, func() {
		trieSet := NewTrieSet()
		trieSet.Insert("this")
		trieSet.Insert("that")
		trieSet.Insert("myage")

		trieSet.Insert("maydelete")
		trieSet.Delete("maydelete")
		gotIsExist1 := trieSet.Contains("maydelete")
		fmt.Printf("gotIsExist1 is %t \r\n", gotIsExist1)
		gotIsExist := trieSet.Contains("this")
		fmt.Printf("isExist is %t \r\n", gotIsExist)
	})
}
