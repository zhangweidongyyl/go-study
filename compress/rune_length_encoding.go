package compress

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

/**
rle  rune length encoding 游程编码，简单的无损压缩算法
适用于 连续重复数据较多的场景。核心思想是将连续的重复数据用出现频率来表示，具体的格式为
【数据值，出现频率】用于减少存储空间
example：
input:[5 5 5 8 8 9 9 3 3 1]
output:[[5,3],[8,2],[9,2],[3,2],[1,1]]
对于数据量的场景，可以直接使用游程编码进行实现
对于大文件，可以使用channel 流式读取
*/

func CreateFile(fileName string, datas []uint32) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Panicln(err)
	}

	for _, data := range datas {
		err := binary.Write(file, binary.LittleEndian, data)
		if err != nil {
			log.Panicln(err)
		}
	}
}
func StreamReadDecodeNoMemory(outputFileName string) {

	datas, err := os.Open(outputFileName)
	if err != nil {
		log.Panicln(err)
	}
	buf := make([]byte, 8)

	for {
		// read 8 bytes every times
		_, err := io.ReadFull(datas, buf)
		if err != nil {
			if err == io.EOF {
				// read file finish break for loop
				break
			}
			log.Panicln(err)
		}
		//log.Printf("read %d bytes from file %s \r\n", length, outputFileName)
		val := binary.LittleEndian.Uint32(buf[:4])

		count := binary.LittleEndian.Uint32(buf[4:8])
		for i := uint32(0); i < count; i++ {
			fmt.Printf("%d ", val)
		}

	}
}

// StreamReadEncodeNoMemory 每次从文件读一个数字进行
func StreamReadEncodeNoMemory(fileName string, outputFileName string) {
	output, err := os.Create(outputFileName)
	if err != nil {
		log.Panicln(err)
	}

	datas, err := os.Open(fileName)
	if err != nil {
		log.Panicln(err)
	}
	buf := make([]byte, 4)
	isFirst := true

	var currentValue uint32
	var count int
	for {
		length, err := io.ReadFull(datas, buf)
		if err != nil {
			if err == io.EOF {
				// read file finish break for loop
				break
			}
			log.Panicln(err)
		}
		log.Printf("read %d bytes from file %s \r\n", length, fileName)
		val := binary.LittleEndian.Uint32(buf)
		if isFirst {
			// the first count
			count = 1
			isFirst = false
			currentValue = val
		} else if val == currentValue {
			// current val is the same with last
			count++
		} else {
			// new
			errVal := binary.Write(output, binary.LittleEndian, currentValue)
			if errVal != nil {
				log.Panicln(errVal)
			}
			errCount := binary.Write(output, binary.LittleEndian, uint32(count))
			if errCount != nil {
				log.Panicln(errCount)
			}
			// 重新记录新的
			currentValue = val
			count = 1
		}
	}
	if !isFirst {
		errVal := binary.Write(output, binary.LittleEndian, currentValue)
		if errVal != nil {
			log.Panicln(errVal)
		}
		errCount := binary.Write(output, binary.LittleEndian, uint32(count))
		if errCount != nil {
			log.Panicln(errCount)
		}
	}
}

// StreamReadEncode 每次从文件读一个数字进行
// 还是有问题，因为最终还是会记录到内存中
func StreamReadEncode(fileName string) []RuneLengthDomain {
	datas, err := os.Open(fileName)
	if err != nil {
		log.Panicln(err)
	}
	buf := make([]byte, 4)
	isFirst := true
	res := make([]RuneLengthDomain, 0)
	var lastValue uint32
	var lastDomain *RuneLengthDomain
	for {
		length, err := io.ReadFull(datas, buf)
		if err != nil {
			if err == io.EOF {
				// read file finish break for loop
				break
			}
			log.Panicln(err)
		}
		log.Printf("read %d bytes from file %s \r\n", length, fileName)
		val := binary.LittleEndian.Uint32(buf)
		if len(res) > 0 {
			lastDomain = &res[len(res)-1]
			lastValue = uint32(lastDomain.Val)
		}
		if isFirst {
			// the first count
			res = append(res, RuneLengthDomain{
				Val:   int(val),
				Count: 1,
			})
			isFirst = false
		} else if val == lastValue {
			// current val is the same with last
			if lastDomain != nil {
				count := lastDomain.Count + 1
				lastDomain.Count = count
			}

		} else {
			// new
			res = append(res, RuneLengthDomain{
				Val:   int(val),
				Count: 1,
			})
		}
	}
	return res
}

type RuneLengthDomain struct {
	Val   int
	Count int
}

func RuneLengthEncode(datas []int) []RuneLengthDomain {
	res := make([]RuneLengthDomain, 0)
	if len(datas) == 0 {
		return res
	}
	res = append(res, RuneLengthDomain{
		Val:   datas[0],
		Count: 1,
	})
	for i := 1; i < len(datas); i++ {
		currentValue := datas[i]
		lastDomain := &res[len(res)-1]
		lastVal := lastDomain.Val
		lastCount := lastDomain.Count
		if currentValue == lastVal {
			lastCount++
			lastDomain.Count = lastCount
		} else {
			res = append(res, RuneLengthDomain{
				Val:   currentValue,
				Count: 1,
			})
		}
	}
	return res

}
func EncodeToBinary(domains []RuneLengthDomain) []byte {
	buf := new(bytes.Buffer)
	for _, domain := range domains {
		err1 := binary.Write(buf, binary.LittleEndian, uint32(domain.Val))
		if err1 != nil {
			log.Panicln(err1)
		}
		err2 := binary.Write(buf, binary.LittleEndian, uint32(domain.Count))
		if err2 != nil {
			log.Panicln(err2)
		}
	}
	return buf.Bytes()
}
func DecodeBinaryToDomains(byteDatas []byte) []RuneLengthDomain {
	res := make([]RuneLengthDomain, 0)
	if len(byteDatas) == 0 {
		return res
	}
	buf := bytes.NewReader(byteDatas)
	for {
		var val, count uint32
		err1 := binary.Read(buf, binary.LittleEndian, &val)
		err2 := binary.Read(buf, binary.LittleEndian, &count)
		if err1 != nil || err2 != nil {
			break
		}
		res = append(res, RuneLengthDomain{
			Val:   int(val),
			Count: int(count),
		})
	}
	return res

}
func RuneLengthDecode(encodeResult []RuneLengthDomain) []int {
	res := make([]int, 0)
	if len(encodeResult) == 0 {
		return res
	}
	for _, domain := range encodeResult {
		currentCount := domain.Count
		for currentCount > 0 {
			res = append(res, domain.Val)
			currentCount--
		}
	}
	return res
}
