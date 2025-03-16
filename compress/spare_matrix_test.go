package compress

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOriginSpareMatrixNums_Compress(t *testing.T) {
	convey.Convey("test compress", t, func() {
		datas := []int{4, 4, 3, 2, 4, 4, 42}
		originDatas := OriginSpareMatrixNums(datas)
		gotEncoded := originDatas.Compress(4)
		fmt.Printf("gotEncoded is %+v \r\n", gotEncoded)

		decompressDatas := gotEncoded.Decompress()
		fmt.Printf("decompressDatas is %+v \r\n", decompressDatas)
	})
}
