package bigfile_external_sorting

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestExternalSorting(t *testing.T) {

	convey.Convey("external sorting", t, func() {
		//baseDir := "/Users/zhangweidong/code/go-study/"
		baseDir := "/Users/zyb/code/mycode/go-study/"
		inputFile := baseDir + "input.txt"
		outputFile := baseDir + "output.txt"
		tempDir := baseDir + "temp"
		lineSize := 2 // 每次读取的行数

		// 确保临时目录存在
		if _, err := os.Stat(tempDir); os.IsNotExist(err) {
			os.Mkdir(tempDir, 0755)
		}

		ExternalSorting(inputFile, outputFile, tempDir, lineSize)
		fmt.Println("External sort completed!")
	})
}
