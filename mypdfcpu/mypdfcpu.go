package mypdfcpu

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"time"
)

func MergePdf(inFilePath, outFilePath string, cnt int) {
	// 输入 PDF 文件列表
	inFiles := []string{}
	for i := 0; i < cnt; i++ {
		inFiles = append(inFiles, inFilePath)
	}

	// 输出合并后的 PDF 文件
	outFile := outFilePath

	start := time.Now()
	// 合并操作
	if err := api.MergeCreateFile(inFiles, outFile, true, nil); err != nil {
		panic(err)
	}
	fmt.Println(time.Since(start))
}

func MergeA() string {
	return fmt.Sprintf("%s-%s", "testA", MergeB())

}
func MergeB() string {
	cValue := MergeC()
	fmt.Printf("testB中testC()的返回值: %s\n", cValue) // 添加日志
	return cValue + ":realB"
}
func MergeC() string {
	return "realC"
}
