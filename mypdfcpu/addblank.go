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

// 像素转毫米的转换函数
func pxToMm(px float64) float64 {
	// 1英寸 = 25.4毫米
	// 1英寸 = 96像素 (标准DPI)
	// 因此：1像素 = 25.4/96 ≈ 0.264583333毫米
	return px * 25.4 / 96
}

// CreateEmptyPDFWithSize 创建一个指定大小的空页PDF文件
func CreateEmptyPDFWithSize(outputPath string, width, height float64, needContert bool) error {
	widthMm := width
	heightMm := height
	if needContert {
		// 将像素转换为毫米
		widthMm = pxToMm(width)
		heightMm = pxToMm(height)
	}

	// 创建新的PDF文档，使用自定义大小（单位：毫米）
	pdf := gofpdf.NewCustom(&gofpdf.InitType{
		UnitStr: "mm",
		Size:    gofpdf.SizeType{Wd: widthMm, Ht: heightMm},
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
	if err := CreateEmptyPDFWithSize("custom.pdf", 719, 888, true); err != nil {
		fmt.Printf("创建自定义大小PDF失败: %v\n", err)
		return
	}

	if err := CreateEmptyPDFWithSize("custom-new.pdf", 184.49, 260.01, false); err != nil {
		fmt.Printf("创建自定义大小PDF失败: %v\n", err)
		return
	}

	fmt.Println("PDF文件创建成功")
}
