package compress

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"unsafe"
)

func TestVarintCompress(t *testing.T) {
	convey.Convey("test compress", t, func() {
		// 示例数据
		numbers := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			numbers[i] = i // 假设数组是递增的
		}

		// 压缩
		compressed := VarintCompress(numbers)
		// 计算压缩前的字节数
		originalSize := len(numbers) * int(unsafe.Sizeof(numbers[0])) // 每个 int 占用的字节数
		fmt.Printf("压缩前的字节数: %d\n", originalSize)
		fmt.Printf("压缩前的字节数为：%d 压缩后的字节数: %d\n", originalSize, len(compressed))

		// 解压
		decompressed, err := VarintDecompress(compressed)
		if err != nil {
			fmt.Println("解压失败:", err)
			return
		}
		fmt.Printf("解压后的数据: %v\n", decompressed[:10]) // 打印前10个数字
	})
}
