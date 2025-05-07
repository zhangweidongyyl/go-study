package mypdfcpu

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

// CreateEmptyPDF 创建一个空页的PDF文件
func CreateEmptyPDF(outputPath string) error {
	// 创建新的PDF文档，使用A4大小
	pdf := gofpdf.New("P", "mm", "A4", "")

	// 添加一个空白页
	pdf.AddPage()

	// 将PDF写入文件
	if err := pdf.OutputFileAndClose(outputPath); err != nil {
		return fmt.Errorf("写入PDF文件失败: %w", err)
	}

	return nil
}

// CreateEmptyPDFWithSize 创建一个指定大小的空页PDF文件
func CreateEmptyPDFWithSize(outputPath string, width, height float64) error {
	// 创建新的PDF文档，使用自定义大小（单位：毫米）
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "mm",
		Size:    gofpdf.SizeType{Wd: width, Ht: height},
	})

	// 添加一个空白页
	pdf.AddPage()

	// 将PDF写入文件
	if err := pdf.OutputFileAndClose(outputPath); err != nil {
		return fmt.Errorf("写入PDF文件失败: %w", err)
	}

	return nil
}

// Example 使用示例
func Example() {
	// 创建A4大小的空页PDF
	if err := CreateEmptyPDF("empty.pdf"); err != nil {
		fmt.Printf("创建PDF失败: %v\n", err)
		return
	}

	// 创建自定义大小的空页PDF (例如: 100x150 毫米)
	if err := CreateEmptyPDFWithSize("custom.pdf", 100, 150); err != nil {
		fmt.Printf("创建自定义大小PDF失败: %v\n", err)
		return
	}

	fmt.Println("PDF文件创建成功")
}
