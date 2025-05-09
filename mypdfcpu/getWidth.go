package mypdfcpu

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"os"
)

// 获取PDF文件尺寸
func getPDFDimensions(pdfPath string) (width, height float64, err error) {
	// 打开PDF文件
	file, err := os.Open(pdfPath)
	if err != nil {
		return 0, 0, fmt.Errorf("打开PDF文件失败: %v", err)
	}

	// 创建配置
	conf := model.NewDefaultConfiguration()

	//// 创建上下文
	//ctx, err := api.ReadContext(file, conf)
	//if err != nil {
	//	return 0, 0, fmt.Errorf("读取PDF文件失败: %v", err)
	//}
	//defer ctx.Close()

	// 获取PDF信息
	info, err := api.PDFInfo(file, pdfPath, []string{"1"}, conf)
	if err != nil {
		return 0, 0, fmt.Errorf("获取PDF信息失败: %v", err)
	}

	// 获取第一页的尺寸
	if len(info.PageBoundaries) > 0 {
		boundary := info.PageBoundaries[0]
		if boundary.Media != nil && boundary.Media.Rect != nil {
			width = boundary.Media.Rect.Width()
			height = boundary.Media.Rect.Height()

			// 转换为毫米：1点 = 25.4/72 ≈ 0.3528毫米
			const pointsToMm = 25.4 / 72
			width = width * pointsToMm
			height = height * pointsToMm

			return width, height, nil
		}
	}

	return 0, 0, fmt.Errorf("PDF文件没有页面")
}

// 获取所有页面的尺寸
func getAllPagesDimensions(pdfPath string) ([]struct{ Width, Height float64 }, error) {
	// 打开PDF文件
	file, err := os.Open(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("打开PDF文件失败: %v", err)
	}

	// 创建配置
	conf := model.NewDefaultConfiguration()

	// 创建上下文
	ctx, err := api.ReadContext(file, conf)
	if err != nil {
		return nil, fmt.Errorf("读取PDF文件失败: %v", err)
	}
	//defer ctx.Close()

	// 获取PDF信息
	info, err := pdfcpu.Info(ctx, pdfPath, nil)
	if err != nil {
		return nil, fmt.Errorf("获取PDF信息失败: %v", err)
	}

	const pointsToMm = 25.4 / 72
	dimensions := make([]struct{ Width, Height float64 }, 0, len(info.PageBoundaries))

	for _, boundary := range info.PageBoundaries {
		if boundary.Media != nil && boundary.Media.Rect != nil {
			width := boundary.Media.Rect.Width() * pointsToMm
			height := boundary.Media.Rect.Height() * pointsToMm

			dimensions = append(dimensions, struct{ Width, Height float64 }{
				Width:  width,
				Height: height,
			})
		}
	}

	return dimensions, nil
}

// 使用示例
func GetWidth() {
	pdfPath := "/Users/zyb/code/mycode/go-study/mypdfcpu/5c7669784e9344c7b8666d5f0645d5ab.pdf"
	//pdfPath := "/Users/zyb/code/mycode/go-study/mypdfcpu/20250508-23.pdf"

	width, height, err := getPDFDimensions(pdfPath)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}
	fmt.Printf("PDF尺寸: %.2f x %.2f 毫米\n", width, height)

}
