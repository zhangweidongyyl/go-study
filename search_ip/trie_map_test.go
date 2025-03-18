package search_ip

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTrieMap_Insert(t *testing.T) {
	convey.Convey("test insert", t, func() {
		trieMap := NewTrieMap()
		trieMap.Insert("apple", 100)
		trieMap.Insert("this", 1)
		trieMap.Insert("that", 22)
		trieMap.Insert("myage", 9999)
		trieMap.Insert("banan", 1010)

		trieMap.Insert("maydelete", 11)
		trieMap.Delete("maydelete")
		gotVal1, gotIsExist1 := trieMap.Search("maydelete")
		results := trieMap.PrefixSearch("th")

		results1 := trieMap.GetAllWords()
		fmt.Printf("results1 is %+v \r\n", results1)
		fmt.Printf("trieMap prefix search results is %+v \r\n", results)
		fmt.Printf("gotVal1 is %+v,gotIsExist1 is %t \r\n", gotVal1, gotIsExist1)
		gotVal, gotIsExist := trieMap.Search("this")
		fmt.Printf("val is %+v,isExist is %t \r\n", gotVal, gotIsExist)
	})
}
