package excel

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestOpenReader(t *testing.T) {
	convey.Convey("read excel datas", t, func() {
		dstFileName := "/Users/zyb/Downloads/主端学校库 - 可上传.xlsx"
		// read excel file
		excelFile, err := os.Open(dstFileName)
		if err != nil {
			panic(err)
		}
		excelData, err := OpenReader(excelFile)
		if err != nil {
			panic(err)
		}
		fmt.Printf("excelData is %+v \r\n", excelData)
	})
}
