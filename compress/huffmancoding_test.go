package compress

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHuffmanCompress(t *testing.T) {
	convey.Convey("huffman code compress", t, func() {

		// 示例数据
		data := "AABBBCCCCDDDDEEEEFFFFF"
		// 构建哈夫曼树
		root := buildHuffmanTree(data)

		// 生成哈夫曼编码表
		codes := make(map[rune]string)
		buildHuffmanCodes(root, "", codes)

		// 打印编码表
		fmt.Println("哈夫曼编码表:")
		for value, code := range codes {
			fmt.Printf("%c: %s\n", value, code)
		}
		b := HuffmanCompress(data, codes)
		decodedString := HuffmanDecompress(b, root, len(data))
		fmt.Printf("decodedString is %s \r\n", string(decodedString))
	})
}
