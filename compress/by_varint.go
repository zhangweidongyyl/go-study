package compress

import (
	"bytes"
	"encoding/binary"
)

// VarintCompress 压缩函数：将整数数组编码为 Varint 字节流
func VarintCompress(numbers []int) []byte {
	var buf bytes.Buffer

	for _, num := range numbers {
		// 使用 binary.PutVarint 将整数编码为 Varint
		varintBuf := make([]byte, binary.MaxVarintLen64)
		n := binary.PutVarint(varintBuf, int64(num))
		buf.Write(varintBuf[:n])
	}

	return buf.Bytes()
}

// VarintDecompress 解压函数：将 Varint 字节流解码为整数数组
func VarintDecompress(compressed []byte) ([]int, error) {
	var numbers []int
	reader := bytes.NewReader(compressed)

	for {
		// 使用 binary.ReadVarint 从字节流中读取 Varint
		num, err := binary.ReadVarint(reader)
		if err != nil {
			break // 读取完毕或出错
		}
		numbers = append(numbers, int(num))
	}

	return numbers, nil
}
