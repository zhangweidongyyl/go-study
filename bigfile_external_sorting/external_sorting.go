package bigfile_external_sorting

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ExternalSorting(inputFile, outputFile, tempDir string, lineSize int) {
	// 打开输入文件
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// chunk每次读取 lineSize 里总的数字
	var chunk []int
	chunkIndex := 0
	lineCount := 0
	var tempFiles []string

	// 分块读取并排序
	for scanner.Scan() {
		line := scanner.Text()

		numStrings := strings.Split(line, " ")
		for _, numString := range numStrings {
			value, err := strconv.Atoi(numString)
			if err != nil {
				panic(fmt.Sprintf("Failed to parse integer in input file: %v", err))
			}
			chunk = append(chunk, value)
		}
		lineCount++
		if lineCount >= lineSize {
			// 写入
			tempFile := sortChunk(chunk, chunkIndex, tempDir)
			tempFiles = append(tempFiles, tempFile)
			chunk = []int{}
			chunkIndex++
			lineCount = 0
		}

	}

	// 检查是否有扫描错误
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading input file: %v", err))
	}

	// 处理最后一个块
	if len(chunk) > 0 {
		tempFile := sortChunk(chunk, chunkIndex, tempDir)
		tempFiles = append(tempFiles, tempFile)
	}

	// 多路归并
	mergeFiles(outputFile, tempFiles)
}

type HeapItem struct {
	val        int
	chunkIndex int
}

// 定义一个最小堆用于多路归并
type MinHeap []HeapItem

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(HeapItem))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 分块排序
func sortChunk(data []int, chunkIndex int, tempDir string) string {
	sort.Ints(data)
	tempFile := fmt.Sprintf("%s/chunk_%d.txt", tempDir, chunkIndex)
	file, err := os.Create(tempFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, value := range data {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", value))
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
	return tempFile
}

// 多路归并
func mergeFiles(outputFile string, tempFiles []string) {
	// 打开所有临时文件
	files := make([]*os.File, len(tempFiles))
	scanners := make([]*bufio.Scanner, len(tempFiles))
	for i, tempFile := range tempFiles {
		file, err := os.Open(tempFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		files[i] = file
		scanners[i] = bufio.NewScanner(file)
	}

	// 创建输出文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	writer := bufio.NewWriter(outFile)

	// 初始化最小堆
	h := &MinHeap{}
	heap.Init(h)
	// 将每个小文件的第一行写入 小顶堆
	for i, scanner := range scanners {
		if scanner.Scan() {
			value, _ := strconv.Atoi(scanner.Text())
			heap.Push(h, HeapItem{
				val:        value,
				chunkIndex: i,
			})
		}
	}

	// 多路归并
	for h.Len() > 0 {
		minHeapItem := heap.Pop(h).(HeapItem)
		value := minHeapItem.val
		chunkIndex := minHeapItem.chunkIndex
		_, err := writer.WriteString(fmt.Sprintf("%d\n", value))
		if err != nil {
			panic(err)
		}

		if scanners[chunkIndex].Scan() {
			nextValue, _ := strconv.Atoi(scanners[chunkIndex].Text())
			heap.Push(h, HeapItem{
				val:        nextValue,
				chunkIndex: chunkIndex,
			})
		}
	}
	writer.Flush()
}
