package bigfile_external_sorting

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestExternalSorting(t *testing.T) {

	convey.Convey("external sorting", t, func() {
		inputFile := "/Users/zyb/code/mycode/go-study/input.txt"
		outputFile := "/Users/zyb/code/mycode/go-study/output.txt"
		tempDir := "/Users/zyb/code/mycode/go-study/temp"
		lineSize := 2 // 每次读取的行数

		// 确保临时目录存在
		if _, err := os.Stat(tempDir); os.IsNotExist(err) {
			os.Mkdir(tempDir, 0755)
		}

		ExternalSorting(inputFile, outputFile, tempDir, lineSize)
		fmt.Println("External sort completed!")
	})
}
