package mypdfcpu

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func MergePdf() {
	// 输入 PDF 文件列表
	inFiles := []string{"/Users/zyb/code/mycode/go-study/mypdfcpu/test1.pdf", "/Users/zyb/code/mycode/go-study/mypdfcpu/mutex.pdf"}

	// 输出合并后的 PDF 文件
	outFile := "/Users/zyb/code/mycode/go-study/mypdfcpu/merged.pdf"

	// 合并操作
	if err := api.MergeCreateFile(inFiles, outFile, true, nil); err != nil {
		panic(err)
	}
}
