package search_ip

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTrieTree_Search(t *testing.T) {
	convey.Convey("search by ip", t, func() {
		originIps := []string{
			"192.168.1.1", "10.3.123.4", "192.167.2.1", "10.36.25.14",
		}
		trieTree := NewTrieTree()
		for _, ip := range originIps {
			trieTree.Insert(ip)
		}

		searchIp := "192.167.2.3"
		isExist := trieTree.Search(searchIp)
		fmt.Printf("searchIp %s isExist is %t \r\n", searchIp, isExist)
	})
}
